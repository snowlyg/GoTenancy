package publish2

import (
	"fmt"
	"path"
	"reflect"
	"regexp"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/validations"
	"github.com/jinzhu/gorm"
)

func init() {
	admin.RegisterViewPath("libs/publish2/views")
}

func (SharedVersion) ConfigureQorResource(res resource.Resourcer) {
	enablePublishMode(res)
}

func (Version) ConfigureQorResource(res resource.Resourcer) {
	enablePublishMode(res)
}

func (Schedule) ConfigureQorResource(res resource.Resourcer) {
	enablePublishMode(res)
}

func (Visible) ConfigureQorResource(res resource.Resourcer) {
	enablePublishMode(res)
}

func enablePublishMode(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		if res.GetTheme("publish2") == nil {
			res.UseTheme("publish2")

			if IsPublishReadyableModel(res.Value) {
				res.OverrideIndexAttrs(func() {
					res.IndexAttrs(res.IndexAttrs(), "-PublishReady")
				})
				res.OverrideEditAttrs(func() {
					res.EditAttrs(res.EditAttrs(), "PublishReady")
				})
				res.OverrideNewAttrs(func() {
					res.NewAttrs(res.NewAttrs(), "PublishReady")
				})
			}

			if IsSchedulableModel(res.Value) {
				res.Meta(&admin.Meta{
					Name:  "ScheduledEventID",
					Label: "Scheduled Event",
					Type:  "hidden",
				})

				res.OverrideIndexAttrs(func() {
					res.IndexAttrs(res.IndexAttrs(), "-ScheduledEventID")
				})
				res.OverrideEditAttrs(func() {
					res.EditAttrs(res.EditAttrs(), "ScheduledStartAt", "ScheduledEndAt", "ScheduledEventID")
				})
				res.OverrideNewAttrs(func() {
					res.NewAttrs(res.NewAttrs(), "ScheduledStartAt", "ScheduledEndAt", "ScheduledEventID")
				})

				if res.GetAdmin().GetResource(utils.ModelType(&Publish{}).Name()) == nil {
					res.GetAdmin().AddResource(&Publish{})
				}
			}

			if IsVersionableModel(res.Value) {
				res.Meta(&admin.Meta{
					Name: "VersionPriority",
					Type: "hidden",
				})

				res.Meta(&admin.Meta{
					Name: "VersionName",
					Type: "hidden",
				})

				res.Action(&admin.Action{
					Name:        "Create New Version",
					Method:      "GET",
					URLOpenType: "slideout",
					URL: func(record interface{}, context *admin.Context) string {
						if versionable, ok := record.(VersionableInterface); ok {
							url, _ := utils.PatchURL(context.URLFor(record), "new_version", "true")
							if versionName := versionable.GetVersionName(); versionName != "" {
								url, _ = utils.PatchURL(url, "version_name", versionName)
							}
							return url
						}
						return ""
					},
					Modes: []string{"menu_item"},
				})

				res.AddValidator(&resource.Validator{
					Name: "version-name-validator",
					Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
						if meta := metaValues.Get("VersionName"); meta != nil {
							versionName := utils.ToString(metaValues.Get("VersionName").Value)
							if strings.ContainsAny(versionName, `!@#$%^&*?/\<>=;`) {
								return validations.NewError(record, "VersionName", "Invalid Version Name")
							}
							return nil
						}
						return nil
					},
				})

				router := res.GetAdmin().GetRouter()
				ctr := controller{Resource: res}
				router.Get(path.Join(res.RoutePrefix(), res.ToParam(), res.ParamIDName(), "versions"), ctr.Versions, &admin.RouteConfig{Resource: res})

				res.OverrideIndexAttrs(func() {
					res.IndexAttrs(res.IndexAttrs(), "-VersionPriority")
				})
				res.OverrideEditAttrs(func() {
					res.EditAttrs(res.EditAttrs(), "-VersionPriority", "VersionName")
				})
				res.OverrideNewAttrs(func() {
					res.NewAttrs(res.NewAttrs(), "-VersionPriority", "VersionName")
				})
			}

			if IsPublishReadyableModel(res.Value) || IsSchedulableModel(res.Value) || IsVersionableModel(res.Value) {
				res.Meta(&admin.Meta{
					Name:  "PublishLiveNow",
					Label: "Live Now",
					Type:  "publish_live_now",
					Valuer: func(interface{}, *qor.Context) interface{} {
						return ""
					},
				})

				res.OverrideIndexAttrs(func() {
					res.IndexAttrs(res.IndexAttrs(), "PublishLiveNow")
				})
				res.OverrideEditAttrs(func() {
					res.EditAttrs(res.EditAttrs(), "-PublishLiveNow")
				})
				res.OverrideNewAttrs(func() {
					res.NewAttrs(res.NewAttrs(), "-PublishLiveNow")
				})
			}

			if IsShareableVersionModel(res.Value) {
				res.Meta(&admin.Meta{
					Name: "VersionName",
					Type: "hidden",
					Valuer: func(record interface{}, context *qor.Context) interface{} {
						if shareableVersion, ok := record.(ShareableVersionInterface); ok {
							return shareableVersion.GetSharedVersionName()
						}
						return ""
					},
					Setter: func(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
					},
				})

				res.Meta(&admin.Meta{
					Name: "ShareableVersion",
					Type: "string",
					Valuer: func(record interface{}, context *qor.Context) interface{} {
						if shareableVersion, ok := record.(ShareableVersionInterface); ok {
							return shareableVersion.GetSharedVersionName() != ""
						}
						return false
					},
					Setter: func(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
						if utils.ToString(metaValue.Value) == "true" {
							if shareableVersion, ok := record.(ShareableVersionInterface); ok {
								versionName := context.Request.Form.Get("QorResource.VersionName")
								shareableVersion.SetSharedVersionName(versionName)
							}
						}
					},
				})
			}

			res.GetAdmin().RegisterFuncMap("get_schedule_event", func(record interface{}, context *admin.Context) interface{} {
				if scheduledInterface, ok := record.(ScheduledInterface); ok {
					var scheduledEvent ScheduledEvent
					if scheduledInterface.GetScheduledEventID() != nil {
						context.GetDB().First(&scheduledEvent, "id = ?", scheduledInterface.GetScheduledEventID())
						return scheduledEvent
					}
				}
				return nil
			})

			res.GetAdmin().RegisterFuncMap("get_default_version_name", func() interface{} {
				return DefaultVersionName
			})

			res.GetAdmin().RegisterFuncMap("get_scheduled_events", func(context *admin.Context) interface{} {
				if res := context.Admin.GetResource("ScheduledEvent"); res != nil {
					scheduleEvents := res.NewSlice()
					context.GetDB().Find(scheduleEvents)
					return scheduleEvents
				}
				return []ScheduledEvent{}
			})

			res.GetAdmin().RegisterFuncMap("get_scheduled_event_resource", func(context *admin.Context) interface{} {
				return context.Admin.GetResource("ScheduledEvent")
			})

			getVersionsCount := func(record interface{}, context *admin.Context) interface{} {
				var (
					count        int
					db           = context.GetDB().Set(VersionNameMode, "").Set(VersionMode, VersionMultipleMode)
					scope        = db.NewScope(record)
					primaryField = scope.PrimaryField()
				)
				db.Set(admin.DisableCompositePrimaryKeyMode, "on").Model(context.Resource.NewStruct()).Where(fmt.Sprintf("%v = ?", scope.Quote(primaryField.DBName)), primaryField.Field.Interface()).Count(&count)
				return count
			}

			res.GetAdmin().RegisterFuncMap("get_new_version_name", func(record interface{}, context *admin.Context) interface{} {
				return fmt.Sprintf("v%v", getVersionsCount(record, context))
			})

			res.GetAdmin().RegisterFuncMap("get_publish_schedule_time", func(context *admin.Context) interface{} {
				return getPublishScheduleTime(context.Context)
			})

			res.GetAdmin().RegisterFuncMap("get_requesting_publish_draft_content", func(context *admin.Context) interface{} {
				return requestingPublishDraftContent(context.Context)
			})

			res.GetAdmin().RegisterFuncMap("get_versions_count", func(record interface{}, context *admin.Context) interface{} {
				return getVersionsCount(record, context)
			})

			res.GetAdmin().RegisterFuncMap("get_live_status", func(record interface{}, context *admin.Context) interface{} {
				var (
					db           = context.GetDB().Set(admin.DisableCompositePrimaryKeyMode, "on").Set(ScheduleMode, "on").Set(VersionMode, "on").Set(VisibleMode, "on").Set(VersionNameMode, "").Set(ScheduledTime, "").Set(ScheduledStart, "").Set(ScheduledEnd, "")
					scope        = db.NewScope(record)
					primaryField = scope.PrimaryField()
					newrecord    = reflect.New(utils.ModelType(record)).Interface()
				)

				if !db.First(newrecord, fmt.Sprintf("%v = ?", scope.Quote(primaryField.DBName)), primaryField.Field.Interface()).RecordNotFound() {
					if oldVersion, ok := record.(VersionableInterface); ok {
						if versionable, ok := newrecord.(VersionableInterface); ok {
							if oldVersion.GetVersionName() == versionable.GetVersionName() {
								return "live"
							}
							return "available"
						}
					}
					return "live"
				}

				return ""
			})
		}
	}
}

type Publish struct {
}

func (Publish) ConfigureQorResourceBeforeInitialize(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		res.Meta(&admin.Meta{
			Name: "ScheduledStartAt",
			Valuer: func(interface{}, *qor.Context) interface{} {
				return ""
			},
		})

		res.UseTheme("publish2")

		if res.Config.Name == "" {
			res.Name = "Schedule"
		}

		if len(res.Config.Menu) == 0 {
			res.Config.Menu = []string{"Publishing"}
		}

		Admin := res.GetAdmin()
		if Admin.GetResource("ScheduledEvent") == nil {
			Admin.AddResource(&ScheduledEvent{}, &admin.Config{Name: "Event", Menu: res.Config.Menu, Priority: -1})
			Admin.DB.AutoMigrate(&ScheduledEvent{})
		}

		scheduledEventResource := res.GetAdmin().GetResource("ScheduledEvent")
		scheduledEventResource.AddProcessor(&resource.Processor{
			Name: "scheduled-event-processor",
			Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
				var (
					db             = context.GetDB()
					scope          = db.NewScope(record)
					startAt, endAt interface{}
				)

				if field, ok := scope.FieldByName("ScheduledStartAt"); ok {
					startAt = field.Field.Interface()
				}
				if field, ok := scope.FieldByName("ScheduledEndAt"); ok {
					endAt = field.Field.Interface()
				}

				if startAt != nil || endAt != nil {
					for _, res := range res.GetAdmin().GetResources() {
						if IsSchedulableModel(res.Value) {
							if err := db.Table(db.NewScope(res.Value).TableName()).Where("scheduled_event_id = ?", scope.PrimaryKeyValue()).UpdateColumns(map[string]interface{}{"scheduled_start_at": startAt, "scheduled_end_at": endAt}).Error; err != nil {
								return err
							}
						}
					}
				}

				return nil
			},
		})

		if Admin.GetRouter().GetMiddleware("publish2") == nil {
			Admin.GetRouter().Use(&admin.Middleware{
				Name: "publish2",
				Handler: func(context *admin.Context, middleware *admin.Middleware) {
					tx := context.GetDB()

					if startAt := context.Request.URL.Query().Get("schedule_start_at"); startAt != "" {
						if t, err := utils.ParseTime(startAt, context.Context); err == nil {
							tx = tx.Set(VisibleMode, "on").Set(ScheduleMode, "on").Set(ScheduledStart, t).Set(VersionMode, VersionMultipleMode)
						}
					}

					if endAt := context.Request.URL.Query().Get("schedule_end_at"); endAt != "" {
						if t, err := utils.ParseTime(endAt, context.Context); err == nil {
							tx = tx.Set(VisibleMode, "on").Set(ScheduleMode, "on").Set(ScheduledEnd, t).Set(VersionMode, VersionMultipleMode)
						}
					}

					usingVersionNameAsPrimaryKey := false
					if res := context.Resource; res != nil {
						for idx, primaryField := range res.PrimaryFields {
							if primaryField.Name == "VersionName" {
								_, params := res.ToPrimaryQueryParams(res.GetPrimaryValue(context.Request), context.Context)
								if len(params) > idx {
									usingVersionNameAsPrimaryKey = true
									tx = tx.Scopes(func(db *gorm.DB) *gorm.DB {
										if db.Value == nil {
											return db.Set(VersionNameMode, params[idx])
										} else if utils.ModelType(res.Value).String() == utils.ModelType(db.Value).String() {
											return db.Set(VersionNameMode, params[idx])
										}
										return db
									})
									break
								}
							}
						}
					}

					if !usingVersionNameAsPrimaryKey {
						versionNameRegexp := regexp.MustCompile(`primary_key\[(.+)_version_name\]`)
						for key, values := range context.Request.URL.Query() {
							if versionNameRegexp.MatchString(key) {
								matches := versionNameRegexp.FindStringSubmatch(key)
								if len(values) > 0 {
									tx = tx.Scopes(func(db *gorm.DB) *gorm.DB {
										if db.Value == nil {
											return db.Set(VersionNameMode, values[0])
										} else if matches[1] == utils.ModelType(db.Value).String() {
											return db.Set(VersionNameMode, values[0])
										}
										return db
									})
								}
							}
						}
					}

					if versionName := context.Request.URL.Query().Get("version_name"); versionName != "" {
						tx = tx.Set(VersionNameMode, versionName)
					}

					context.SetDB(tx)
					middleware.Next(context)
				},
			})

			ctr := controller{Resource: res}
			Admin.GetRouter().Get(res.ToParam(), ctr.Dashboard, &admin.RouteConfig{Resource: res})
		}
	}
}
