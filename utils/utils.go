package utils

import (
	"fmt"
	"go/build"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"GoTenancy/config/auth"
	"GoTenancy/config/db"
	"GoTenancy/models/users"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"github.com/microcosm-cc/bluemonday"
	"github.com/qor/l10n"
	"github.com/qor/qor/utils"
	"github.com/qor/session"
	"github.com/qor/session/manager"
)

// GetCurrentUser 从请求中获取当前用户
func GetCurrentUser(req *http.Request) *users.User {
	if currentUser, ok := auth.Auth.GetCurrentUser(req).(*users.User); ok {
		return currentUser
	}
	return nil
}

// GetCurrentLocale 从请求中获取本地设置
func GetCurrentLocale(req *http.Request) string {
	locale := l10n.Global
	if cookie, err := req.Cookie("locale"); err == nil {
		locale = cookie.Value
	}
	return locale
}

// GetDB 从请求中获取 DB
func GetDB(req *http.Request) *gorm.DB {
	if db := utils.GetDBFromRequest(req); db != nil {
		return db
	}
	return db.DB
}

// AddFlashMessage 辅助方法
func AddFlashMessage(w http.ResponseWriter, req *http.Request, message string, mtype string) error {
	return manager.SessionManager.Flash(w, req, session.Message{Message: template.HTML(message), Type: mtype})
}

// HTMLSanitizer HTML 消毒器
var HTMLSanitizer = bluemonday.UGCPolicy()

// FormatPrice 价格格式化
func FormatPrice(price interface{}) string {
	switch price.(type) {
	case float32, float64:
		return fmt.Sprintf("%0.2f", price)
	case int, uint, int32, int64, uint32, uint64:
		return fmt.Sprintf("%d.00", price)
	}
	return ""
}

// DetectViewsDir 解决 go mod 模式无法注册 qor-admin 等包的 views
func DetectViewsDir(pkgorg, pkgname string) string {
	var foundp string
	var found bool

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	path := filepath.Join(gopath, "/pkg/mod/")
	ppath := filepath.Join(path, pkgorg)
	if _, err := os.Stat(ppath); err == nil {
		if err = filepath.Walk(ppath, func(p string, f os.FileInfo, err error) error { // nolint: errcheck, gosec, unparam
			if found {
				return nil
			}
			if strings.HasPrefix(filepath.Base(p), pkgname) {
				vp := filepath.Join(p, "views")
				if _, err := os.Stat(vp); err == nil {
					foundp = vp
					found = true
				} else {
					return err
				}
			}
			return nil
		}); err != nil {
			color.Red(fmt.Sprintf("os.Stat error %v\n", err))
		}

	} else {
		color.Red(fmt.Sprintf("os.Stat error %v\n", err))
	}

	return foundp
}
