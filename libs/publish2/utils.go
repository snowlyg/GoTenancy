package publish2

import (
	"fmt"
	"net/http"

	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/utils"
	"github.com/jinzhu/gorm"
)

func getPublishScheduleTime(context *qor.Context) string {
	if values, ok := context.Request.URL.Query()["publish_scheduled_time"]; ok {
		if len(values) > 0 && values[0] != "" {
			return values[0]
		}
	} else if cookie, err := context.Request.Cookie("publish2_publish_scheduled_time"); err == nil {
		return cookie.Value
	}
	return ""
}

func requestingPublishDraftContent(context *qor.Context) bool {
	if values, ok := context.Request.URL.Query()["publish_draft_content"]; ok {
		if len(values) > 0 && values[0] != "" {
			return true
		}
	} else if cookie, err := context.Request.Cookie("pubilsh2_publish_draft_content"); err == nil && cookie.Value == "true" {
		return true
	}
	return false
}

func PreviewByDB(tx *gorm.DB, context *qor.Context) *gorm.DB {
	scheduledTime := getPublishScheduleTime(context)
	draftContent := requestingPublishDraftContent(context)

	utils.SetCookie(http.Cookie{Name: "publish2_publish_scheduled_time", Value: scheduledTime}, context)
	utils.SetCookie(http.Cookie{Name: "pubilsh2_publish_draft_content", Value: fmt.Sprint(draftContent)}, context)

	if scheduledTime != "" {
		if t, err := utils.ParseTime(scheduledTime, context); err == nil {
			tx = tx.Set(ScheduledTime, t)
		}
	}

	if draftContent {
		tx = tx.Set(VisibleMode, ModeOff)
	}

	return tx
}
