# Serializable Meta

Serializable Meta allows the developer to specify, for a given model, a custom serialization model along with field mappings. This mechanism thus allows one model to act as another model when it comes to serialization.

[![GoDoc](https://godoc.org/GoTenancy/libs/serializable_meta?status.svg)](https://godoc.org/GoTenancy/libs/serializable_meta)

## Usage

The example herein shows how to manage different kinds of background jobs using https://GoTenancy/libs/serializable_meta.

### Define serializable model

Define `QorJob` model and embed `serializable_meta.SerializableMeta` to apply the feature.

```go
type QorJob struct {
  gorm.Model
  Name string
  serializable_meta.SerializableMeta
}
```

Add function `GetSerializableArgumentResource` to the model, so [Serializable Meta](https://GoTenancy/libs/serializable_meta) can know the type of argument. Then define background jobs.

```go
func (qorJob QorJob) GetSerializableArgumentResource() *admin.Resource {
  return jobsArgumentsMap[qorJob.Kind]
}

var jobsArgumentsMap = map[string]*admin.Resource{
  "newsletter": admin.NewResource(&sendNewsletterArgument{}),
  "import_products": admin.NewResource(&importProductArgument{}),
}

type sendNewsletterArgument struct {
  Subject string
  Content string
}

type importProductArgument struct {}
```

### Use serializable features

At first, Set a job's `Name`, `Kind` and `SetSerializableArgumentValue`. Then save it into database.

```go
var qorJob QorJob
qorJob.Name = "sending newsletter"
qorJob.Kind = "newsletter"
qorJob.SetSerializableArgumentValue(&sendNewsletterArgument{
  Subject: "subject",
  Content: "content",
})

db.Create(&qorJob)
```

This will marshal `sendNewsletterArgument` as a json, and save it into database by this SQL

```sql
INSERT INTO "qor_jobs" (kind, value) VALUES (`newsletter`, `{"Subject":"subject","Content":"content"}`);
```

Now you can fetch the saved `QorJob` from the database. And get the serialized data from the record.

```go
var result QorJob
db.First(&result, "name = ?", "sending newsletter")

var argument = result.GetSerializableArgument(result)
argument.(*sendNewsletterArgument).Subject // "subject"
argument.(*sendNewsletterArgument).Content // "content"
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
