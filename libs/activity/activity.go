package activity

import (
	"fmt"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/audited"
	"GoTenancy/libs/media/asset_manager"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/validations"
	"github.com/jinzhu/gorm"
)

// QorActivity default model used to save resource's activities
type QorActivity struct {
	gorm.Model
	Action       string
	Content      string `sql:"size:5000"`
	Note         string `sql:"size:2000"`
	Type         string
	ResourceType string
	ResourceID   string
	CreatorName  string
	audited.AuditedModel
}

// Register register activity feature for an qor resource
func Register(res *admin.Resource) {
	var (
		qorAdmin         = res.GetAdmin()
		activityResource = qorAdmin.GetResource("QorActivity")
	)

	if activityResource == nil {
		// Auto run migration before add resource
		res.GetAdmin().DB.AutoMigrate(&QorActivity{})

		activityResource = qorAdmin.AddResource(&QorActivity{}, &admin.Config{Invisible: true})
		activityResource.Meta(&admin.Meta{Name: "Action", Type: "hidden", Valuer: func(value interface{}, ctx *qor.Context) interface{} {
			act := value.(*QorActivity)
			if act.Action == "" {
				act.Action = "comment on"
			}
			return activityResource.GetAdmin().T(ctx, "activity."+act.Action, act.Action)
		}})
		activityResource.Meta(&admin.Meta{Name: "UpdatedAt", Type: "hidden", Valuer: func(value interface{}, ctx *qor.Context) interface{} {
			return utils.FormatTime(value.(*QorActivity).UpdatedAt, "Jan 2 15:04", ctx)
		}})
		activityResource.Meta(&admin.Meta{Name: "URL", Valuer: func(value interface{}, ctx *qor.Context) interface{} {
			return fmt.Sprintf("/admin/%v/%v/!%v/%v/edit", res.ToParam(), res.GetPrimaryValue(ctx.Request), activityResource.ToParam(), value.(*QorActivity).ID)
		}})

		assetManager := qorAdmin.GetResource("AssetManager")
		if assetManager == nil {
			assetManager = qorAdmin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
		}

		activityResource.Meta(&admin.Meta{Name: "Content", Type: "rich_editor", Resource: assetManager})
		activityResource.Meta(&admin.Meta{Name: "Note", Type: "string", Resource: assetManager})
		activityResource.EditAttrs("Action", "Content", "Note")
		activityResource.ShowAttrs("ID", "Action", "Content", "Note", "URL", "UpdatedAt", "CreatorName")
		activityResource.AddValidator(&resource.Validator{
			Name: "activity-content-validator",
			Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
				if meta := metaValues.Get("Content"); meta != nil {
					if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
						return validations.NewError(record, "Content", "Content can't be blank")
					}
				}
				return nil
			},
		})
	}

	res.GetAdmin().RegisterViewPath("libs/activity/views")
	res.UseTheme("activity")

	qorAdmin.RegisterFuncMap("activity_resource", func() *admin.Resource {
		return qorAdmin.GetResource("QorActivity")
	})

	qorAdmin.RegisterFuncMap("get_activities_count", func(context *admin.Context) int {
		return GetActivitiesCount(context, context.Result)
	})

	router := res.GetAdmin().GetRouter()
	ctrl := controller{ActivityResource: activityResource}
	router.Get(fmt.Sprintf("/%v/%v/!qor_activities", res.ToParam(), res.ParamIDName()), ctrl.GetActivity)
	router.Post(fmt.Sprintf("/%v/%v/!qor_activities", res.ToParam(), res.ParamIDName()), ctrl.CreateActivity)
	router.Post(fmt.Sprintf("/%v/%v/!qor_activities/%v/edit", res.ToParam(), res.ParamIDName(), activityResource.ParamIDName()), ctrl.UpdateActivity)
}
