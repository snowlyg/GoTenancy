package publish2

import (
	"fmt"
	"html/template"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/roles"
)

type controller struct {
	Resource *admin.Resource
}

type visiblePublishResourceInterface interface {
	VisiblePublishResource(*qor.Context) bool
}

func (ctr controller) Dashboard(context *admin.Context) {
	type resourceResult struct {
		Resource            *admin.Resource
		ComingOnlineResults interface{}
		GoingOfflineResults interface{}
	}

	var results = []resourceResult{}

	for _, res := range context.Admin.GetResources() {
		if IsSchedulableModel(res.Value) {
			if visibleInterface, ok := res.Value.(visiblePublishResourceInterface); ok {
				if !visibleInterface.VisiblePublishResource(context.Context) {
					continue
				}
			} else if res.Config.Invisible {
				continue
			}

			db := context.GetDB()
			result := resourceResult{Resource: res}

			comingOnlineData := res.NewSlice()
			if db.Set(VisibleMode, "on").Set(ScheduleMode, ComingOnlineMode).Set(VersionMode, VersionMultipleMode).Find(comingOnlineData).RowsAffected > 0 {
				result.ComingOnlineResults = comingOnlineData
			}

			goingOfflineData := res.NewSlice()
			if db.Set(VisibleMode, "on").Set(ScheduleMode, GoingOfflineMode).Set(VersionMode, VersionMultipleMode).Find(goingOfflineData).RowsAffected > 0 {
				result.GoingOfflineResults = goingOfflineData
			}

			if result.ComingOnlineResults != nil || result.GoingOfflineResults != nil {
				results = append(results, result)
			}
		}
	}

	context.Action = "index"
	context.Execute("publish2/dashboard", results)
}

func (ctr controller) Versions(context *admin.Context) {
	records := context.Resource.NewSlice()
	record := context.Resource.NewStruct()
	primaryQuerySQL, primaryParams := ctr.Resource.ToPrimaryQueryParams(context.ResourceID, context.Context)
	tx := context.GetDB().Set(admin.DisableCompositePrimaryKeyMode, "on").Set(VersionMode, VersionMultipleMode).Set(ScheduleMode, ModeOff).Set(VisibleMode, ModeOff)
	tx.Where(primaryQuerySQL, primaryParams...).First(record)

	scope := tx.NewScope(record)
	tx.Find(records, fmt.Sprintf("%v = ?", scope.PrimaryKey()), scope.PrimaryKeyValue())

	result := context.Funcs(template.FuncMap{
		"version_metas": func() (metas []*admin.Meta) {
			for _, name := range []string{"VersionName", "ScheduledStartAt", "ScheduledEndAt", "PublishReady", "PublishLiveNow"} {
				if meta := ctr.Resource.GetMeta(name); meta != nil && meta.HasPermission(roles.Read, context.Context) {
					metas = append(metas, meta)
				}
			}
			return
		},
	}).Render("publish2/versions", records)
	context.Writer.Write([]byte(result))
}
