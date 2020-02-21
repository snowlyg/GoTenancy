package exchange

import (
	"reflect"

	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/roles"
)

// Meta defines importable/exportable fields
type Meta struct {
	base *Resource
	resource.Meta
	Name       string
	Header     string
	Valuer     func(interface{}, *qor.Context) interface{}
	Setter     func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context)
	Permission *roles.Permission
}

// GetMetas get defined sub metas
func (meta *Meta) GetMetas() []resource.Metaor {
	return []resource.Metaor{}
}

// GetResource get its resource
func (meta *Meta) GetResource() resource.Resourcer {
	return nil
}

func (meta *Meta) updateMeta() {
	meta.Meta = resource.Meta{
		Name:         meta.Name,
		FieldName:    meta.FieldName,
		Setter:       meta.Setter,
		Valuer:       meta.Valuer,
		Permission:   meta.Permission,
		BaseResource: meta.base,
	}

	meta.PreInitialize()
	if meta.FieldStruct != nil {
		if injector, ok := reflect.New(meta.FieldStruct.Struct.Type).Interface().(resource.ConfigureMetaBeforeInitializeInterface); ok {
			injector.ConfigureQorMetaBeforeInitialize(meta)
		}
	}

	meta.Initialize()

	if meta.FieldStruct != nil {
		if injector, ok := reflect.New(meta.FieldStruct.Struct.Type).Interface().(resource.ConfigureMetaInterface); ok {
			injector.ConfigureQorMeta(meta)
		}
	}

	meta.SetFormattedValuer(func(record interface{}, context *qor.Context) interface{} {
		if valuer := meta.GetValuer(); valuer != nil {
			result := valuer(record, context)

			if reflectValue := reflect.ValueOf(result); reflectValue.IsValid() {
				if reflectValue.Kind() == reflect.Ptr {
					if reflectValue.IsNil() || !reflectValue.Elem().IsValid() {
						return ""
					}

					result = reflectValue.Elem().Interface()
				}

				return result
			}
		}
		return ""
	})
}
