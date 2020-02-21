package media

import (
	"encoding/json"
	"errors"
	"reflect"

	"GoTenancy/libs/serializable_meta"
	"github.com/jinzhu/gorm"
)

var (
	// set MediaLibraryURL to change the default url /system/{{class}}/{{primary_key}}/{{column}}.{{extension}}
	MediaLibraryURL = ""
)

func cropField(field *gorm.Field, scope *gorm.Scope) (cropped bool) {
	if field.Field.CanAddr() {
		// TODO Handle scanner
		if media, ok := field.Field.Addr().Interface().(Media); ok && !media.Cropped() {
			option := parseTagOption(field.Tag.Get("media_library"))
			if MediaLibraryURL != "" {
				option.Set("url", MediaLibraryURL)
			}
			if media.GetFileHeader() != nil || media.NeedCrop() {
				var mediaFile FileInterface
				var err error
				if fileHeader := media.GetFileHeader(); fileHeader != nil {
					mediaFile, err = media.GetFileHeader().Open()
				} else {
					mediaFile, err = media.Retrieve(media.URL("original"))
				}

				if err != nil {
					scope.Err(err)
					return false
				}

				media.Cropped(true)

				if url := media.GetURL(option, scope, field, media); url == "" {
					scope.Err(errors.New("invalid URL"))
				} else {
					result, _ := json.Marshal(map[string]string{"Url": url})
					media.Scan(string(result))
				}

				if mediaFile != nil {
					defer mediaFile.Close()
					var handled = false
					for _, handler := range mediaHandlers {
						if handler.CouldHandle(media) {
							mediaFile.Seek(0, 0)
							if scope.Err(handler.Handle(media, mediaFile, option)) == nil {
								handled = true
							}
						}
					}

					// Save File
					if !handled {
						scope.Err(media.Store(media.URL(), option, mediaFile))
					}
				}
				return true
			}
		}
	}
	return false
}

func saveAndCropImage(isCreate bool) func(scope *gorm.Scope) {
	return func(scope *gorm.Scope) {
		if !scope.HasError() {
			var updateColumns = map[string]interface{}{}

			// Handle SerializableMeta
			if value, ok := scope.Value.(serializable_meta.SerializableMetaInterface); ok {
				var (
					isCropped        bool
					handleNestedCrop func(record interface{})
				)

				handleNestedCrop = func(record interface{}) {
					newScope := scope.New(record)
					for _, field := range newScope.Fields() {
						if cropField(field, scope) {
							isCropped = true
							continue
						}

						if reflect.Indirect(field.Field).Kind() == reflect.Struct {
							handleNestedCrop(field.Field.Addr().Interface())
						}

						if reflect.Indirect(field.Field).Kind() == reflect.Slice {
							for i := 0; i < reflect.Indirect(field.Field).Len(); i++ {
								handleNestedCrop(reflect.Indirect(field.Field).Index(i).Addr().Interface())
							}
						}
					}
				}

				record := value.GetSerializableArgument(value)
				handleNestedCrop(record)
				if isCreate && isCropped {
					updateColumns["value"], _ = json.Marshal(record)
				}
			}

			// Handle Normal Field
			for _, field := range scope.Fields() {
				if cropField(field, scope) && isCreate {
					updateColumns[field.DBName] = field.Field.Interface()
				}
			}

			if !scope.HasError() && len(updateColumns) != 0 {
				scope.Err(scope.NewDB().Model(scope.Value).UpdateColumns(updateColumns).Error)
			}
		}
	}
}

// RegisterCallbacks register callback into GORM DB
func RegisterCallbacks(db *gorm.DB) {
	if db.Callback().Create().Get("media:save_and_crop") == nil {
		db.Callback().Create().After("gorm:after_create").Register("media:save_and_crop", saveAndCropImage(true))
	}
	if db.Callback().Update().Get("media:save_and_crop") == nil {
		db.Callback().Update().Before("gorm:before_update").Register("media:save_and_crop", saveAndCropImage(false))
	}
}
