package sorting

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/roles"
	"github.com/jinzhu/gorm"
)

type SortableCollection struct {
	PrimaryField string
	PrimaryKeys  []string
}

func (sortableCollection *SortableCollection) Scan(value interface{}) error {
	switch values := value.(type) {
	case []string:
		sortableCollection.PrimaryKeys = values
	case string:
		return json.Unmarshal([]byte(values), sortableCollection)
	case []byte:
		return json.Unmarshal(values, sortableCollection)
	default:
		return errors.New("unsupported driver -> Scan pair for MediaLibrary")
	}

	return nil
}

func (sortableCollection SortableCollection) Value() (driver.Value, error) {
	results, err := json.Marshal(sortableCollection)
	return string(results), err
}

func (sortableCollection SortableCollection) Sort(results interface{}) error {
	values := reflect.Indirect(reflect.ValueOf(results))

	if values.Kind() != reflect.Slice {
		return errors.New("invalid type")
	}

	scope := gorm.Scope{Value: values.Interface()}
	if primaryField := scope.PrimaryField(); primaryField != nil {
		var (
			primaryFieldName = primaryField.Name
			indirectValues   = reflect.Indirect(values)
			sliceType        = indirectValues.Type()
			slice            = reflect.MakeSlice(sliceType, 0, 0)
			orderedMap       = map[int]bool{}
		)

		for _, primaryKey := range sortableCollection.PrimaryKeys {
			for i := 0; i < indirectValues.Len(); i++ {
				value := indirectValues.Index(i)
				field := reflect.Indirect(value).FieldByName(primaryFieldName)
				if fmt.Sprint(field.Interface()) == primaryKey {
					if _, ok := orderedMap[i]; !ok {
						slice = reflect.Append(slice, value)
						orderedMap[i] = true
					}
					break
				}
			}
		}

		for i := 0; i < indirectValues.Len(); i++ {
			if _, ok := orderedMap[i]; !ok {
				slice = reflect.Append(slice, indirectValues.Index(i))
			}
		}

		if values.Kind() != reflect.Ptr {
			for i := 0; i < slice.Len(); i++ {
				values.Index(i).Set(slice.Index(i))
			}
		}
	}

	return nil
}

func (sortableCollection *SortableCollection) ConfigureQorMeta(metaor resource.Metaor) {
	if meta, ok := metaor.(*admin.Meta); ok {
		var (
			name         = strings.TrimSuffix(meta.GetName(), "Sorter")
			res          = meta.GetBaseResource().(*admin.Resource)
			sortableMeta = res.GetMeta(name)
		)

		res.UseTheme("sortable_collection")

		if sortableMeta != nil {
			sortableMeta.AddProcessor(&admin.MetaProcessor{
				Name: "sortable-collection-meta-processor",
				Handler: func(sortableMeta *admin.Meta) {
					if sortableMeta.Type == "select_many" {
						if selectManyConfig, ok := sortableMeta.Config.(*admin.SelectManyConfig); ok {
							if selectManyConfig.SelectMode == "" {
								selectManyConfig.SelectMode = "select"
							}
							if selectManyConfig.SelectionTemplate == "" {
								selectManyConfig.SelectionTemplate = "metas/form/sortable_select_many.tmpl"
							}
						}

						setter := sortableMeta.GetSetter()
						sortableMeta.SetSetter(func(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
							primaryKeys := utils.ToArray(metaValue.Value)
							reflectValue := reflect.Indirect(reflect.ValueOf(record))
							reflectValue.FieldByName(meta.GetName()).Addr().Interface().(*SortableCollection).Scan(primaryKeys)
							setter(record, metaValue, context)
						})

						valuer := sortableMeta.GetValuer()
						sortableMeta.SetValuer(func(record interface{}, context *qor.Context) interface{} {
							results := valuer(record, context)
							reflectValue := reflect.Indirect(reflect.ValueOf(record))
							reflectValue.FieldByName(meta.GetName()).Interface().(SortableCollection).Sort(results)
							return results
						})

						meta.SetSetter(func(interface{}, *resource.MetaValue, *qor.Context) {})
						meta.SetPermission(roles.Deny(roles.CRUD, roles.Anyone))
					}

					if sortableMeta.Type == "collection_edit" {
						if collectionEditConfig, ok := sortableMeta.Config.(*admin.CollectionEditConfig); ok {
							if collectionEditConfig.Template == "" {
								collectionEditConfig.Template = "metas/form/sortable_collection_edit.tmpl"
							}
						}

						valuer := sortableMeta.GetValuer()
						sortableMeta.SetValuer(func(record interface{}, context *qor.Context) interface{} {
							results := valuer(record, context)
							reflectValue := reflect.Indirect(reflect.ValueOf(record))
							reflectValue.FieldByName(meta.GetName()).Interface().(SortableCollection).Sort(results)
							return results
						})

						res.AddProcessor(&resource.Processor{
							Name: "sortable_collection_processor",
							Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
								var primaryValues []string
								reflectValue := reflect.Indirect(reflect.ValueOf(record))
								fieldValue := reflect.Indirect(reflectValue.FieldByName(name))
								if fieldValue.Kind() == reflect.Slice {
									for i := 0; i < fieldValue.Len(); i++ {
										scope := gorm.Scope{Value: fieldValue.Index(i).Interface()}
										primaryValues = append(primaryValues, fmt.Sprint(scope.PrimaryKeyValue()))
									}
									reflectValue.FieldByName(meta.GetName()).Addr().Interface().(*SortableCollection).Scan(primaryValues)
								}

								return nil
							},
						})

						meta.SetSetter(func(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
							primaryKeys := utils.ToArray(metaValue.Value)
							reflectValue := reflect.Indirect(reflect.ValueOf(record))
							reflectValue.FieldByName(meta.GetName()).Addr().Interface().(*SortableCollection).Scan(primaryKeys)
						})

						meta.SetPermission(roles.Deny(roles.CRUD, roles.Anyone))
					}
				},
			})
		}
	}
}
