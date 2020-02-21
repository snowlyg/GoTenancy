package admin

import (
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"github.com/jinzhu/gorm"
)

// Filter filter definiation
type Filter struct {
	Name       string
	Label      string
	Type       string
	Operations []string // eq, cont, gt, gteq, lt, lteq
	Resource   *Resource
	Visible    func(context *Context) bool
	Handler    func(*gorm.DB, *FilterArgument) *gorm.DB
	Config     FilterConfigInterface
}

// SavedFilter saved filter settings
type SavedFilter struct {
	Name string
	URL  string
}

// FilterConfigInterface filter config interface
type FilterConfigInterface interface {
	ConfigureQORAdminFilter(*Filter)
}

// FilterArgument filter argument that used in handler
type FilterArgument struct {
	Value    *resource.MetaValues
	Resource *Resource
	Context  *qor.Context
}

// Filter register filter for qor resource
func (res *Resource) Filter(filter *Filter) {
	filter.Resource = res

	if filter.Label == "" {
		filter.Label = utils.HumanizeString(filter.Name)
	}

	if meta := res.GetMeta(filter.Name); meta != nil {
		if filter.Type == "" {
			filter.Type = meta.Type
		}

		if filter.Config == nil {
			if filterConfig, ok := meta.Config.(FilterConfigInterface); ok {
				filter.Config = filterConfig
			}
		}
	}

	if filter.Config != nil {
		filter.Config.ConfigureQORAdminFilter(filter)
	}

	if filter.Handler == nil {
		// generate default handler
		filter.Handler = func(db *gorm.DB, filterArgument *FilterArgument) *gorm.DB {
			if metaValue := filterArgument.Value.Get("Value"); metaValue != nil {
				if keyword := utils.ToString(metaValue.Value); keyword != "" {
					field := filterField{FieldName: filter.Name}
					if operationMeta := filterArgument.Value.Get("Operation"); operationMeta != nil {
						if operation := utils.ToString(operationMeta.Value); operation != "" {
							field.Operation = operation
						}
					}
					if field.Operation == "" {
						if len(filter.Operations) > 0 {
							field.Operation = filter.Operations[0]
						} else {
							field.Operation = "contains"
						}
					}

					return filterResourceByFields(res, []filterField{field}, keyword, db, filterArgument.Context)
				}
			}
			return db
		}
	}

	if filter.Type != "" {
		res.filters = append(res.filters, filter)
	} else {
		utils.ExitWithMsg("Invalid filter definition %v for resource %v", filter.Name, res.Name)
	}
}

// GetFilters get registered filters
func (res *Resource) GetFilters() []*Filter {
	return res.filters
}

// GetFilter get defined action
func (res *Resource) GetFilter(name string) *Filter {
	for _, action := range res.filters {
		if action.Name == name {
			return action
		}
	}
	return nil
}
