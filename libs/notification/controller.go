package notification

import (
	"net/http"
	"strconv"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/responder"
)

type controller struct {
	Notification *Notification
	action       *Action
}

func (c *controller) List(context *admin.Context) {
	context.Set("Notification", c.Notification)
	var currentPage int
	if p, err := strconv.Atoi(context.Request.URL.Query().Get("page")); err == nil {
		currentPage = p
	}

	context.Execute("notifications/notifications", map[string]interface{}{
		"Messages":         c.Notification.GetNotifications(context.CurrentUser, context.Context),
		"LoadMoreNextPage": currentPage + 1,
	})
}

func (c *controller) Action(context *admin.Context) {
	action := c.action
	message := c.Notification.GetNotification(context.CurrentUser, context.ResourceID, context.Context)
	context.Set("Notification", c.Notification)

	if context.Request.Method == "GET" {
		context.Execute("action", action)
	} else {
		var actionArgument = &ActionArgument{
			Message: message,
			Context: context,
		}

		if action.Resource != nil {
			result := action.Resource.NewStruct()
			action.Resource.Decode(context.Context, result)
			actionArgument.Argument = result
		}

		if err := action.Handler(actionArgument); err == nil {
			flash := action.FlashMessage(actionArgument, true /* succeed */, false /* undo */)
			responder.With("html", func() {
				context.Flash(flash, "success")
				http.Redirect(context.Writer, context.Request, context.Request.Referer(), http.StatusFound)
			}).With("json", func() {
				notification := c.Notification.GetNotification(context.CurrentUser, context.ResourceID, context.Context)
				context.JSON("OK", map[string]string{"status": "ok", "message": flash, "notification": string(context.Render("notification", notification))})
			}).Respond(context.Request)
		} else {
			notification := c.Notification.GetNotification(context.CurrentUser, context.ResourceID, context.Context)
			context.JSON("OK", map[string]string{"status": "error", "error": action.FlashMessage(actionArgument, false /* succeed */, false /* undo */), "notification": string(context.Render("notification", notification))})
		}
	}
}

func (c *controller) UndoAction(context *admin.Context) {
	action := c.action
	message := c.Notification.GetNotification(context.CurrentUser, context.ResourceID, context.Context)
	context.Set("Notification", c.Notification)

	var actionArgument = &ActionArgument{
		Message: message,
		Context: context,
	}

	if action.Resource != nil {
		result := action.Resource.NewStruct()
		action.Resource.Decode(context.Context, result)
		actionArgument.Argument = result
	}

	if err := action.Undo(actionArgument); err == nil {
		flash := action.FlashMessage(actionArgument, true /* succeed */, true /* undo */)
		responder.With("html", func() {
			context.Flash(flash, "success")
			http.Redirect(context.Writer, context.Request, context.Request.Referer(), http.StatusFound)
		}).With("json", func() {
			notification := c.Notification.GetNotification(context.CurrentUser, context.ResourceID, context.Context)
			context.JSON("OK", map[string]string{"status": "ok", "message": flash, "notification": string(context.Render("notification", notification))})
		}).Respond(context.Request)
	} else {
		notification := c.Notification.GetNotification(context.CurrentUser, context.ResourceID, context.Context)
		context.JSON("OK", map[string]string{"status": "error", "error": action.FlashMessage(actionArgument, false /* succeed */, true /* undo */), "notification": string(context.Render("notification", notification))})
	}
}
