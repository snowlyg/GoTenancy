# QOR Notification

QOR Notification provides a way to send notifications to [QOR Admin](https://GoTenancy/libs/admin) administrators. Notifications can be anything your system needs, like order update, delivery notices, whatever.

## Usage

```go
Notification := notification.New(&notification.Config{})

// Add to Admin
Admin.NewResource(Notification)

// Register Database Channel
Notification.RegisterChannel(database.New(&database.Config{DB: db.DB}))

// Send Notification
Notification.Send(message *Message, context *qor.Context)

// Get Notification
Notification.GetNotification(user interface{}, messageID string, context *qor.Context)

// Get Notifications
Notification.GetNotifications(user interface{}, context *qor.Context)

// Get Unresolved Notifications Count
Notification.GetUnresolvedNotificationsCount(user interface{}, context *qor.Context)
```

The Notifications List in [QOR Admin](../chapter2/setup.md) looks a bit like this when populated:

![notification](notification-demo.png)

## Register Actions for Notification

This example shows how to add a "Dismiss" button to notification. The button will appears in the notification which `ResolvedAt` is nil. Please read the [Action](http://doc.getqor.com/admin/actions.html) documentation for more details.

```go
Notification.Action(&notification.Action{
  Name:         "Dismiss",
  MessageTypes: []string{"info", "order_processed", "order_returned"},
  Visible: func(data *notification.QorNotification, context *admin.Context) bool {
    return data.ResolvedAt == nil
  },
  Handle: func(argument *notification.ActionArgument) error {
    return argument.Context.GetDB().Model(argument.Message).Update("resolved_at", time.Now()).Error
  },
})
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
