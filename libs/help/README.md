# QOR Help

QOR Help provides a way to add help documents to [QOR Admin](http://GoTenancy/libs/admin)

## Usage

First, add QOR Help table to the database.

```go
db.DB.AutoMigrate(&help.QorHelpEntry{})
```

Then add QOR Help to [QOR Admin](http://GoTenancy/libs/admin).

```go
Admin.NewResource(&help.QorHelpEntry{})
```

Now start your application. You should see a question mark icon appears at the top right corner of [QOR Admin](http://GoTenancy/libs/admin) interface, click the icon, a slide panel should appear, the Admin user could get knowledge from here directly.

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
