package action_bar

import (
	"net/http"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor/utils"
)

type controller struct {
	ActionBar *ActionBar
}

// SwitchMode is handle to store switch status in cookie
func (controller) SwitchMode(context *admin.Context) {
	utils.SetCookie(http.Cookie{Name: "qor-action-bar", Value: context.Request.URL.Query().Get("checked")}, context.Context)

	referrer := context.Request.Referer()
	if referrer == "" {
		referrer = "/"
	}

	http.Redirect(context.Writer, context.Request, referrer, http.StatusFound)
}

// InlineEdit using to make inline edit resource shown as slideout
func (controller) InlineEdit(context *admin.Context) {
	context.Writer.Write([]byte(context.Render("action_bar/inline_edit")))
}
