package notification

import (
	"path"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/roles"
)

type Notification struct {
	Config   *Config
	Channels []ChannelInterface
	Actions  []*Action
}

func New(config *Config) *Notification {
	notification := &Notification{Config: config}
	return notification
}

func (notification *Notification) Send(message *Message, context *qor.Context) error {
	for _, channel := range notification.Channels {
		if err := channel.Send(message, context); err != nil {
			return err
		}
	}
	return nil
}

type NotificationsResult struct {
	Notification  *Notification
	Notifications []*QorNotification
	Resolved      []*QorNotification
}

func (notification *Notification) GetNotifications(user interface{}, context *qor.Context) *NotificationsResult {
	var results = NotificationsResult{
		Notification: notification,
	}

	for _, channel := range notification.Channels {
		channel.GetNotifications(user, &results, notification, context)
	}

	return &results
}

func (notification *Notification) GetUnresolvedNotificationsCount(user interface{}, context *qor.Context) uint {
	var result uint
	for _, channel := range notification.Channels {
		result += channel.GetUnresolvedNotificationsCount(user, notification, context)
	}
	return result
}

func (notification *Notification) GetNotification(user interface{}, messageID string, context *qor.Context) *QorNotification {
	for _, channel := range notification.Channels {
		if message, err := channel.GetNotification(user, messageID, notification, context); err == nil {
			return message
		}
	}
	return nil
}

func (notification *Notification) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		Admin := res.GetAdmin()
		Admin.RegisterViewPath("GoTenancy/libs/notification/views")

		if len(notification.Channels) == 0 {
			utils.ExitWithMsg("No channel defined for notification")
		}

		Admin.RegisterFuncMap("unresolved_notifications_count", func(context *admin.Context) uint {
			return notification.GetUnresolvedNotificationsCount(context.CurrentUser, context.Context)
		})

		router := Admin.GetRouter()
		notificationController := controller{Notification: notification}

		router.Get("!notifications", notificationController.List, &admin.RouteConfig{
			PermissionMode: roles.Read,
			Resource:       res,
		})

		for _, action := range notification.Actions {
			actionController := controller{Notification: notification, action: action}
			router.Get(path.Join("!notifications", res.ParamIDName(), action.ToParam()), actionController.Action, &admin.RouteConfig{
				PermissionMode: roles.Update,
				Resource:       res,
			})

			router.Put(path.Join("!notifications", res.ParamIDName(), action.ToParam()), actionController.Action, &admin.RouteConfig{
				PermissionMode: roles.Update,
				Resource:       res,
			})

			if action.Undo != nil {
				router.Put(path.Join("!notifications", res.ParamIDName(), action.ToParam(), "undo"), actionController.UndoAction, &admin.RouteConfig{
					PermissionMode: roles.Update,
					Resource:       res,
				})
			}
		}
	}
}
