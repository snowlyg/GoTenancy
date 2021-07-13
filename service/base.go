package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

var ConfigTypes = []Option{
	{Value: "input", Label: "文本框"},
	{Value: "number", Label: "数字框"},
	{Value: "textarea", Label: "多行文本框"},
	{Value: "radio", Label: "单选框"},
	{Value: "checkbox", Label: "多选框"},
	{Value: "select", Label: "下拉框"},
	{Value: "file", Label: "文件上传"},
	{Value: "image", Label: "图片上传"},
	{Value: "color", Label: "颜色选择框"},
}

func GetConfigTypeName(value string) string {
	for i := 0; i < len(ConfigTypes); i++ {
		if ConfigTypes[i].Value.(string) == value {
			return ConfigTypes[i].Label
		}
	}
	return ""
}

type Form struct {
	Rule    []Rule                   `json:"rule"`
	Action  string                   `json:"action"`
	Method  string                   `json:"method"`
	Title   string                   `json:"title"`
	Config  Config                   `json:"config"`
	Headers []map[string]interface{} `json:"headers,omitempty"`
}

func (form *Form) SetAction(uri string, ctx *gin.Context) {
	form.Action = SetUrl(uri, ctx)
}

func SetUrl(uri string, ctx *gin.Context) string {
	if multi.IsAdmin(ctx) {
		return g.TENANCY_CONFIG.System.AdminPreix + uri
	} else if multi.IsTenancy(ctx) {
		return g.TENANCY_CONFIG.System.ClientPreix + uri
	}
	return ""
}

type Config struct {
}

type Rule struct {
	Title    string                   `json:"title"`
	Type     string                   `json:"type"`
	Field    string                   `json:"field"`
	Info     string                   `json:"info"`
	Value    interface{}              `json:"value"`
	Props    map[string]interface{}   `json:"props"`
	Options  []Option                 `json:"options,omitempty"`
	Control  []Control                `json:"control,omitempty"`
	Validate []map[string]interface{} `json:"validate,omitempty"`
}
type ControlRule struct {
	Title    string                   `json:"title"`
	Type     string                   `json:"type"`
	Field    string                   `json:"field"`
	Info     string                   `json:"info"`
	Value    interface{}              `json:"value"`
	Props    map[string]interface{}   `json:"props"`
	Options  []Option                 `json:"options,omitempty"`
	Validate []map[string]interface{} `json:"validate,omitempty"`
}

type Control struct {
	Value int    `json:"value"`
	Rule  []Rule `json:"rule"`
}

func (r *Rule) TransData(rule string, token []byte) {
	switch r.Type {
	case "input":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
			"type":        "text",
		}
	case "textarea":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
			"type":        "textarea",
		}
		r.Type = "input"
	case "number":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
		}
		r.Type = "inputNumber"
	case "radio":
		r.Props = map[string]interface{}{}
		rules := strings.Split(rule, ";")
		for _, ru := range rules {
			rus := strings.Split(ru, ":")
			if len(rus) == 2 {
				r.Options = append(r.Options, Option{Label: rus[1], Value: rus[0]})
			}
		}
	case "file":
		seitURL, _ := GetSeitURL()
		r.Props = map[string]interface{}{
			"action": seitURL + "v1/admin/media/upload",
			"data":   map[string]interface{}{},
			"headers": map[string]interface{}{
				"Authorization": "Bearer " + string(token),
			},

			"limit":      1,
			"uploadType": "file",
		}
		r.Type = "upload"
	case "image":
		r.Props = map[string]interface{}{
			"footer":    false,
			"height":    "480px",
			"maxLength": 1,
			"modal":     map[string]interface{}{"modal": false},
			"src":       "/admin/setting/uploadPicture?field=" + r.Field + "&type=1",
			"title":     "请选择" + r.Title,
			"type":      r.Type,
			"width":     "896px",
		}
		r.Type = "frame"
	}

}

type Option struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	Children []Option    `json:"children"`
}

type Opt struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

// filterDate
func filterDate(db *gorm.DB, date, perfix string) *gorm.DB {
	field := "created_at"
	if perfix != "" {
		field = fmt.Sprintf("%s.created_at", perfix)
	}
	dates := strings.Split(date, "-")
	if len(dates) == 2 {
		start, _ := time.Parse("2006/01/02", dates[0])
		end, _ := time.Parse("2006/01/02", dates[1])
		return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), start, end)
	}
	if len(dates) == 1 {
		// { text: '今天', val: 'today' },
		// { text: '昨天', val: 'yesterday' },
		// { text: '最近7天', val: 'lately7' },
		// { text: '最近30天', val: 'lately30' },
		// { text: '本月', val: 'month' },
		// { text: '本年', val: 'year' }
		// TODO: 使用内置函数，可能造成索引失效
		switch dates[0] {
		case "today":
			return db.Where(fmt.Sprintf("TO_DAYS(NOW()) - TO_DAYS(%s) < 1", field))
		case "yesterday":
			return db.Where(fmt.Sprintf("TO_DAYS(NOW()) - TO_DAYS(%s) = 1", field))
		case "lately7":
			return db.Where(fmt.Sprintf("DATE_SUB(CURDATE(),INTERVAL 7 DAY) <= DATE(%s)", field))
		case "lately30":
			return db.Where(fmt.Sprintf("DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(%s)", field))
		case "month":
			return db.Where(fmt.Sprintf("DATE_FORMAT( %s, '%%Y%%m' ) = DATE_FORMAT( CURDATE() , '%%Y%%m' )", field))
		case "year":
			return db.Where(fmt.Sprintf("YEAR(%s)=YEAR(NOW())", field))
		}
	}
	return db
}

func GetIsDelField(ctx *gin.Context) string {
	if multi.IsAdmin(ctx) {
		return ""
	} else if multi.IsTenancy(ctx) {
		return "is_system_del"
	}
	// 用户删除
	return "is_del"
}
