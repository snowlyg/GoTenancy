package widget

import (
	"fmt"
	"html/template"
	"reflect"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/utils"
	"github.com/jinzhu/gorm"
)

// Context widget context
type Context struct {
	Widgets          *Widgets
	DB               *gorm.DB
	AvailableWidgets []string
	Options          map[string]interface{}
	Body             string
	InlineEdit       bool
	SourceType       string
	SourceID         string
	FuncMaps         template.FuncMap
	WidgetSetting    QorWidgetSettingInterface
}

// Get get option with name
func (context Context) Get(name string) (interface{}, bool) {
	if value, ok := context.Options[name]; ok {
		return value, true
	}

	return nil, false
}

// Set set option by name
func (context *Context) Set(name string, value interface{}) {
	if context.Options == nil {
		context.Options = map[string]interface{}{}
	}
	context.Options[name] = value
}

// GetDB set option by name
func (context *Context) GetDB() *gorm.DB {
	if context.DB != nil {
		return context.DB
	}
	return context.Widgets.Config.DB
}

// Clone clone a context
func (context *Context) Clone() *Context {
	return &Context{
		Widgets:          context.Widgets,
		DB:               context.DB,
		AvailableWidgets: context.AvailableWidgets,
		Options:          context.Options,
		InlineEdit:       context.InlineEdit,
		FuncMaps:         context.FuncMaps,
		WidgetSetting:    context.WidgetSetting,
	}
}

// Render render widget based on context
func (context *Context) Render(widgetName string, widgetGroupName string) template.HTML {
	var (
		visibleScopes         []string
		widgets               = context.Widgets
		widgetSettingResource = widgets.WidgetSettingResource
		clone                 = context.Clone()
	)

	for _, scope := range registeredScopes {
		if scope.Visible(context) {
			visibleScopes = append(visibleScopes, scope.ToParam())
		}
	}

	if setting := context.findWidgetSetting(widgetName, append(visibleScopes, "default"), widgetGroupName); setting != nil {
		clone.WidgetSetting = setting
		adminContext := admin.Context{Admin: context.Widgets.Config.Admin, Context: &qor.Context{DB: context.DB}}

		var (
			widgetObj     = GetWidget(setting.GetSerializableArgumentKind())
			widgetSetting = widgetObj.Context(clone, setting.GetSerializableArgument(setting))
		)

		if clone.InlineEdit {
			prefix := widgets.Resource.GetAdmin().GetRouter().Prefix
			inlineEditURL := adminContext.URLFor(setting, widgetSettingResource)
			if widgetObj.InlineEditURL != nil {
				inlineEditURL = widgetObj.InlineEditURL(context)
			}

			return template.HTML(fmt.Sprintf(
				"<script data-prefix=\"%v\" src=\"%v/assets/javascripts/widget_check.js?theme=widget\"></script><div class=\"qor-widget qor-widget-%v\" data-widget-inline-edit-url=\"%v\" data-url=\"%v\">\n%v\n</div>",
				prefix,
				prefix,
				utils.ToParamString(widgetObj.Name),
				fmt.Sprintf("%v/%v/inline-edit", prefix, widgets.Resource.ToParam()),
				inlineEditURL,
				widgetObj.Render(widgetSetting, setting.GetTemplate()),
			))
		}

		return widgetObj.Render(widgetSetting, setting.GetTemplate())
	}

	return template.HTML("")
}

func (context *Context) findWidgetSetting(widgetName string, scopes []string, widgetGroupName string) QorWidgetSettingInterface {
	var (
		db                    = context.GetDB()
		widgetSettingResource = context.Widgets.WidgetSettingResource
		setting               QorWidgetSettingInterface
		settings              = widgetSettingResource.NewSlice()
	)

	if context.SourceID != "" {
		db.Order("source_id DESC").Where("name = ? AND scope IN (?) AND ((shared = ? AND source_type = ?) OR (source_type = ? AND source_id = ?))", widgetName, scopes, true, "", context.SourceType, context.SourceID).Find(settings)
	} else {
		db.Where("name = ? AND scope IN (?) AND source_type = ?", widgetName, scopes, "").Find(settings)
	}

	settingsValue := reflect.Indirect(reflect.ValueOf(settings))
	if settingsValue.Len() > 0 {
	OUTTER:
		for _, scope := range scopes {
			for i := 0; i < settingsValue.Len(); i++ {
				s := settingsValue.Index(i).Interface().(QorWidgetSettingInterface)
				if s.GetScope() == scope {
					setting = s
					break OUTTER
				}
			}
		}
	}

	if context.SourceType == "" {
		if setting == nil {
			if widgetGroupName == "" {
				utils.ExitWithMsg("Widget: Can't Create Widget Without Widget Type")
				return nil
			}
			setting = widgetSettingResource.NewStruct().(QorWidgetSettingInterface)
			setting.SetWidgetName(widgetName)
			setting.SetGroupName(widgetGroupName)
			setting.SetSerializableArgumentKind(widgetGroupName)
			db.Create(setting)
		} else if setting.GetGroupName() != widgetGroupName {
			setting.SetGroupName(widgetGroupName)
			db.Save(setting)
		}
	}

	return setting
}
