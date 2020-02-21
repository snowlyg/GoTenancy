package admin

import (
	"errors"

	"GoTenancy/libs/qor"
)

// metaConfig meta config
type metaConfig struct {
}

// GetTemplate get customized template for meta
func (metaConfig) GetTemplate(context *Context, metaType string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

var defaultMetaConfigorMaps = map[string]func(*Meta){
	"date": func(meta *Meta) {
		if _, ok := meta.Config.(*DatetimeConfig); !ok || meta.Config == nil {
			meta.Config = &DatetimeConfig{}
			meta.Config.ConfigureQorMeta(meta)
		}
	},

	"datetime": func(meta *Meta) {
		if _, ok := meta.Config.(*DatetimeConfig); !ok || meta.Config == nil {
			meta.Config = &DatetimeConfig{ShowTime: true}
			meta.Config.ConfigureQorMeta(meta)
		}
	},

	"string": func(meta *Meta) {
		if meta.FormattedValuer == nil {
			meta.SetFormattedValuer(func(value interface{}, context *qor.Context) interface{} {
				switch str := meta.GetValuer()(value, context).(type) {
				case *string:
					if str != nil {
						return *str
					}
					return ""
				case string:
					return str
				default:
					return str
				}
			})
		}
	},

	"text": func(meta *Meta) {
		if meta.FormattedValuer == nil {
			meta.SetFormattedValuer(func(value interface{}, context *qor.Context) interface{} {
				switch str := meta.GetValuer()(value, context).(type) {
				case *string:
					if str != nil {
						return *str
					}
					return ""
				case string:
					return str
				default:
					return str
				}
			})
		}
	},

	"select_one": func(meta *Meta) {
		if metaConfig, ok := meta.Config.(*SelectOneConfig); !ok || metaConfig == nil {
			meta.Config = &SelectOneConfig{Collection: meta.Collection}
			meta.Config.ConfigureQorMeta(meta)
		} else if meta.Collection != nil {
			metaConfig.Collection = meta.Collection
			meta.Config.ConfigureQorMeta(meta)
		}
	},

	"select_many": func(meta *Meta) {
		if metaConfig, ok := meta.Config.(*SelectManyConfig); !ok || metaConfig == nil {
			meta.Config = &SelectManyConfig{Collection: meta.Collection}
			meta.Config.ConfigureQorMeta(meta)
		} else if meta.Collection != nil {
			metaConfig.Collection = meta.Collection
			meta.Config.ConfigureQorMeta(meta)
		}
	},

	"single_edit": func(meta *Meta) {
		if _, ok := meta.Config.(*SingleEditConfig); !ok || meta.Config == nil {
			meta.Config = &SingleEditConfig{}
			meta.Config.ConfigureQorMeta(meta)
		}
	},

	"collection_edit": func(meta *Meta) {
		if _, ok := meta.Config.(*CollectionEditConfig); !ok || meta.Config == nil {
			meta.Config = &CollectionEditConfig{}
			meta.Config.ConfigureQorMeta(meta)
		}
	},

	"rich_editor": func(meta *Meta) {
		if _, ok := meta.Config.(*RichEditorConfig); !ok || meta.Config == nil {
			meta.Config = &RichEditorConfig{}
			meta.Config.ConfigureQorMeta(meta)
		}
	},
}
