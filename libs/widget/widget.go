package widget

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/assetfs"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/roles"
	"github.com/jinzhu/gorm"
)

var (
	viewPaths              []string
	registeredWidgets      []*Widget
	registeredWidgetsGroup []*WidgetsGroup
)

// Config widget config
type Config struct {
	DB            *gorm.DB
	Admin         *admin.Admin
	PreviewAssets []string
}

// New new widgets container
func New(config *Config) *Widgets {
	widgets := &Widgets{Config: config, funcMaps: template.FuncMap{}, AssetFS: assetfs.AssetFS().NameSpace("widgets")}

	if utils.AppRoot != "" {
		widgets.RegisterViewPath(filepath.Join(utils.AppRoot, "app/views/widgets"))
	}
	widgets.RegisterViewPath("app/views/widgets")
	return widgets
}

// Widgets widgets container
type Widgets struct {
	funcMaps              template.FuncMap
	Config                *Config
	Resource              *admin.Resource
	AssetFS               assetfs.Interface
	WidgetSettingResource *admin.Resource
}

// SetAssetFS set asset fs for render
func (widgets *Widgets) SetAssetFS(assetFS assetfs.Interface) {
	for _, viewPath := range viewPaths {
		assetFS.RegisterPath(viewPath)
	}

	widgets.AssetFS = assetFS
}

// RegisterWidget register a new widget
func (widgets *Widgets) RegisterWidget(w *Widget) {
	registeredWidgets = append(registeredWidgets, w)
}

// RegisterWidgetsGroup register widgets group
func (widgets *Widgets) RegisterWidgetsGroup(group *WidgetsGroup) {
	registeredWidgetsGroup = append(registeredWidgetsGroup, group)
}

// RegisterFuncMap register view funcs, it could be used when render templates
func (widgets *Widgets) RegisterFuncMap(name string, fc interface{}) {
	widgets.funcMaps[name] = fc
}

// ConfigureQorResourceBeforeInitialize a method used to config Widget for qor admin
func (widgets *Widgets) ConfigureQorResourceBeforeInitialize(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		// register view paths
		res.GetAdmin().RegisterViewPath("libs/widget/views")

		// set resources
		widgets.Resource = res

		// set setting resource
		if widgets.WidgetSettingResource == nil {
			widgets.WidgetSettingResource = res.GetAdmin().NewResource(&QorWidgetSetting{}, &admin.Config{Name: res.Name})
		}

		res.Name = widgets.WidgetSettingResource.Name

		for funcName, fc := range funcMap {
			res.GetAdmin().RegisterFuncMap(funcName, fc)
		}

		// configure routes
		controller := widgetController{Widgets: widgets}
		router := res.GetAdmin().GetRouter()
		router.Get(widgets.WidgetSettingResource.ToParam(), controller.Index, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
		router.Get(fmt.Sprintf("%v/new", widgets.WidgetSettingResource.ToParam()), controller.New, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
		router.Get(fmt.Sprintf("%v/!setting", widgets.WidgetSettingResource.ToParam()), controller.Setting, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
		router.Get(fmt.Sprintf("%v/%v", widgets.WidgetSettingResource.ToParam(), widgets.WidgetSettingResource.ParamIDName()), controller.Edit, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
		router.Get(fmt.Sprintf("%v/%v/!preview", widgets.WidgetSettingResource.ToParam(), widgets.WidgetSettingResource.ParamIDName()), controller.Preview, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
		router.Get(fmt.Sprintf("%v/%v/edit", widgets.WidgetSettingResource.ToParam(), widgets.WidgetSettingResource.ParamIDName()), controller.Edit, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
		router.Put(fmt.Sprintf("%v/%v", widgets.WidgetSettingResource.ToParam(), widgets.WidgetSettingResource.ParamIDName()), controller.Update, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
		router.Post(widgets.WidgetSettingResource.ToParam(), controller.Update, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
		router.Get(fmt.Sprintf("%v/inline-edit", res.ToParam()), controller.InlineEdit, &admin.RouteConfig{Resource: widgets.WidgetSettingResource})
	}
}

// Widget widget struct
type Widget struct {
	Name          string
	PreviewIcon   string
	Group         string
	Templates     []string
	Setting       *admin.Resource
	Permission    *roles.Permission
	InlineEditURL func(*Context) string
	Context       func(context *Context, setting interface{}) *Context
}

// WidgetsGroup widgets Group
type WidgetsGroup struct {
	Name    string
	Widgets []string
}

// GetWidget get widget by name
func GetWidget(name string) *Widget {
	for _, w := range registeredWidgets {
		if w.Name == name {
			return w
		}
	}

	for _, g := range registeredWidgetsGroup {
		if g.Name == name {
			for _, widgetName := range g.Widgets {
				return GetWidget(widgetName)
			}
		}
	}

	log.Println("widget " + name + " not exists, make sure to register it")
	return nil
}

// GetWidgets GetWidgets return registered widgets
func GetWidgets() []*Widget {
	return registeredWidgets
}
