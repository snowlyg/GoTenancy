package notification

import "GoTenancy/libs/qor"

func (notification *Notification) RegisterChannel(channel ChannelInterface) {
	notification.Channels = append(notification.Channels, channel)
}

type ChannelInterface interface {
	Send(message *Message, context *qor.Context) error
	GetNotifications(user interface{}, results *NotificationsResult, notification *Notification, context *qor.Context) error
	GetUnresolvedNotificationsCount(user interface{}, notification *Notification, context *qor.Context) uint
	GetNotification(user interface{}, notificationID string, notification *Notification, context *qor.Context) (*QorNotification, error)
}
