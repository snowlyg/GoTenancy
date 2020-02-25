package registerviews

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

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

			if strings.HasPrefix(filepath.Base(p), pkgname+"@") {
				vp := filepath.Join(p, "views")
				cvp := filepath.Join(p, "clean/views") // auth_themes
				if _, err := os.Stat(vp); err == nil {
					foundp = vp
					found = true
				} else if _, err := os.Stat(cvp); err == nil {
					foundp = cvp
					found = true
				} else {
					return err
				}
			}
			return nil
		}); err != nil {
			color.Red(fmt.Sprintf("os.Stat2 error %v\n", err))
		}

	} else {
		color.Red(fmt.Sprintf("os.Stat1 error %v\n", err))
	}

	return foundp
}
