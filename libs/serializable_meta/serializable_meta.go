package serializable_meta

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
)

// SerializableMetaInterface is a interface defined methods need for a serializable model
type SerializableMetaInterface interface {
	GetSerializableArgumentResource() *admin.Resource
	GetSerializableArgument(SerializableMetaInterface) interface{}
	GetSerializableArgumentKind() string
	SetSerializableArgumentKind(name string)
	SetSerializableArgumentValue(interface{})
}

// SerializableMeta default struct that implemented SerializableMetaInterface
type SerializableMeta struct {
	Kind  string
	Value serializableArgument `sql:"size:65532"`
}

type serializableArgument struct {
	SerializedValue string
	OriginalValue   interface{}
}

func (sa *serializableArgument) Scan(data interface{}) (err error) {
	switch values := data.(type) {
	case []byte:
		sa.SerializedValue = string(values)
	case string:
		sa.SerializedValue = values
	default:
		err = errors.New("unsupported driver -> Scan pair for MediaLibrary")
	}
	return
}

func (sa serializableArgument) Value() (driver.Value, error) {
	if sa.OriginalValue != nil {
		result, err := json.Marshal(sa.OriginalValue)
		return string(result), err
	}
	return sa.SerializedValue, nil
}

// GetSerializableArgumentKind get serializable argument kind
func (serialize SerializableMeta) GetSerializableArgumentKind() string {
	return serialize.Kind
}

// SetSerializableArgumentKind set serializable argument kind
func (serialize *SerializableMeta) SetSerializableArgumentKind(name string) {
	serialize.Kind = name
}

// GetSerializableArgument get serializable argument
func (serialize *SerializableMeta) GetSerializableArgument(serializableMetaInterface SerializableMetaInterface) interface{} {
	if serialize.Value.OriginalValue != nil {
		return serialize.Value.OriginalValue
	}

	if res := serializableMetaInterface.GetSerializableArgumentResource(); res != nil {
		value := res.NewStruct()
		json.Unmarshal([]byte(serialize.Value.SerializedValue), value)
		return value
	}
	return nil
}

// SetSerializableArgumentValue set serializable argument value
func (serialize *SerializableMeta) SetSerializableArgumentValue(value interface{}) {
	serialize.Value.OriginalValue = value
}

// ConfigureQorResourceBeforeInitialize configure qor resource for qor admin
func (serialize *SerializableMeta) ConfigureQorResourceBeforeInitialize(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		res.GetAdmin().RegisterViewPath("GoTenancy/libs/serializable_meta/views")

		if _, ok := res.Value.(SerializableMetaInterface); ok {
			res.Meta(&admin.Meta{
				Name: "Kind",
				Type: "hidden",
				Valuer: func(value interface{}, context *qor.Context) interface{} {
					defer func() {
						if r := recover(); r != nil {
							utils.ExitWithMsg("SerializableMeta: Can't Get Kind")
						}
					}()

					return value.(SerializableMetaInterface).GetSerializableArgumentKind()
				},
				Setter: func(value interface{}, metaValue *resource.MetaValue, context *qor.Context) {
					value.(SerializableMetaInterface).SetSerializableArgumentKind(utils.ToString(metaValue.Value))
				},
			})

			res.Meta(&admin.Meta{
				Name: "SerializableMeta",
				Type: "serializable_meta",
				Valuer: func(value interface{}, context *qor.Context) interface{} {
					if serializeArgument, ok := value.(SerializableMetaInterface); ok {
						return struct {
							Value    interface{}
							Resource *admin.Resource
						}{
							Value:    serializeArgument.GetSerializableArgument(serializeArgument),
							Resource: serializeArgument.GetSerializableArgumentResource(),
						}
					}
					return nil
				},
				FormattedValuer: func(value interface{}, context *qor.Context) interface{} {
					if serializeArgument, ok := value.(SerializableMetaInterface); ok {
						return serializeArgument.GetSerializableArgument(serializeArgument)
					}
					return nil
				},
				Setter: func(result interface{}, metaValue *resource.MetaValue, context *qor.Context) {
					if serializeArgument, ok := result.(SerializableMetaInterface); ok {
						if serializeArgumentResource := serializeArgument.GetSerializableArgumentResource(); serializeArgumentResource != nil {
							var clearUpRecord, fillUpRecord func(record interface{}, metaors []resource.Metaor, metaValues []*resource.MetaValue)
							// Keep original value, so if user don't have permission to update some fields, we won't lost the data
							value := serializeArgument.GetSerializableArgument(serializeArgument)

							for _, v := range serializeArgumentResource.Validators {
								context.AddError(v.Handler(value, metaValue.MetaValues, context))
							}

							// Clear all nested slices if has related form data
							clearUpRecord = func(record interface{}, metaors []resource.Metaor, metaValues []*resource.MetaValue) {
								for _, meta := range metaors {
									for _, metaValue := range metaValues {
										if meta.GetName() == metaValue.Name {
											if metaResource, ok := meta.GetResource().(*admin.Resource); ok && metaResource != nil && metaValue.MetaValues != nil {
												nestedFieldValue := reflect.Indirect(reflect.ValueOf(record)).FieldByName(meta.GetFieldName())
												if nestedFieldValue.Kind() == reflect.Struct {
													clearUpRecord(nestedFieldValue.Addr().Interface(), metaResource.GetMetas([]string{}), metaValue.MetaValues.Values)
												} else if nestedFieldValue.Kind() == reflect.Slice {
													nestedFieldValue.Set(reflect.Zero(nestedFieldValue.Type()))
												}
											}
										}
									}
								}
							}
							clearUpRecord(value, serializeArgumentResource.GetMetas([]string{}), metaValue.MetaValues.Values)

							fillUpRecord = func(record interface{}, metaors []resource.Metaor, metaValues []*resource.MetaValue) {
								for _, meta := range metaors {
									for _, metaValue := range metaValues {
										if meta.GetName() == metaValue.Name {
											if setter := meta.GetSetter(); setter != nil {
												setter(record, metaValue, context)
											}

											if metaResource, ok := meta.GetResource().(*admin.Resource); ok && metaResource != nil && metaValue.MetaValues != nil {
												nestedFieldValue := reflect.Indirect(reflect.ValueOf(record)).FieldByName(meta.GetFieldName())
												if meta.GetSetter() == nil || reflect.Indirect(nestedFieldValue).Type() == utils.ModelType(metaResource.NewStruct()) {
													if nestedFieldValue.Kind() == reflect.Struct {
														nestedValue := nestedFieldValue.Addr().Interface()
														for _, v := range metaResource.Validators {
															context.AddError(v.Handler(nestedValue, metaValue.MetaValues, context))
														}

														fillUpRecord(nestedValue, metaResource.GetMetas([]string{}), metaValue.MetaValues.Values)

														for _, p := range metaResource.Processors {
															context.AddError(p.Handler(nestedValue, metaValue.MetaValues, context))
														}
													}

													if nestedFieldValue.Kind() == reflect.Slice {
														nestedValue := reflect.New(nestedFieldValue.Type().Elem())

														for _, v := range metaResource.Validators {
															context.AddError(v.Handler(nestedValue, metaValue.MetaValues, context))
														}

														if destroy := metaValue.MetaValues.Get("_destroy"); destroy == nil || fmt.Sprint(destroy.Value) == "0" {
															fillUpRecord(nestedValue.Interface(), metaResource.GetMetas([]string{}), metaValue.MetaValues.Values)
															if !reflect.DeepEqual(reflect.Zero(nestedFieldValue.Type().Elem()).Interface(), nestedValue.Elem().Interface()) {
																nestedFieldValue.Set(reflect.Append(nestedFieldValue, nestedValue.Elem()))

																for _, p := range metaResource.Processors {
																	context.AddError(p.Handler(nestedValue, metaValue.MetaValues, context))
																}
															}
														}
													}
													continue
												}
											}
										}
									}
								}
							}

							fillUpRecord(value, serializeArgumentResource.GetMetas([]string{}), metaValue.MetaValues.Values)

							for _, p := range serializeArgumentResource.Processors {
								context.AddError(p.Handler(value, metaValue.MetaValues, context))
							}
							serializeArgument.SetSerializableArgumentValue(value)
						}
					}
				},
			})

			res.NewAttrs("Kind", "SerializableMeta")
			res.EditAttrs("ID", "Kind", "SerializableMeta")
		}
	}
}
