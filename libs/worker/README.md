# Worker

Worker runs a single [*Job*](<https://en.wikipedia.org/wiki/Job_(computing)>) in the background, it can do so immediately or at a scheduled time.

Once registered with [QOR Admin](http://GoTenancy/libs/admin), [Worker](https://GoTenancy/libs/worker) will provide a `Workers` section in the navigation tree, containing pages for listing and managing the following aspects of Workers:

  - All *Jobs*.
  - Running: *Jobs* that are currently running.
  - Scheduled: *Jobs* which have been scheduled to run at a time in the future.
  - Done: finished *Jobs*.
  - Errors: any errors output from any Workers that have been run.

The admin interface for a schedulable *Job* will have an additional `Schedule Time` input, with which administrators can set the scheduled date and time.

[![GoDoc](https://godoc.org/GoTenancy/libs/worker?status.svg)](https://godoc.org/GoTenancy/libs/worker)

## Usage

```go
import "GoTenancy/libs/worker"

func main() {
  // Define Worker
  Worker := worker.New()

  // Arguments used to run a job
  type sendNewsletterArgument struct {
    Subject      string
    Content      string `sql:"size:65532"`
    SendPassword string

    // If job's argument has `worker.Schedule` embedded, it will get run at a scheduled time
    worker.Schedule
  }

  // Register Job
  Worker.RegisterJob(&worker.Job{
    Name: "Send Newsletter", // Registerd Job Name
    Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
      // `AddLog` add job log
      qorJob.AddLog("Started sending newsletters...")
      qorJob.AddLog(fmt.Sprintf("Argument: %+v", argument.(*sendNewsletterArgument)))

      for i := 1; i <= 100; i++ {
        time.Sleep(100 * time.Millisecond)
        qorJob.AddLog(fmt.Sprintf("Sending newsletter %v...", i))
        // `SetProgress` set job progress percent, from 0 - 100
        qorJob.SetProgress(uint(i))
      }

      qorJob.AddLog("Finished send newsletters")
      return nil
    },
    // Arguments used to run a job
    Resource: Admin.NewResource(&sendNewsletterArgument{}),
  })

  // Add Worker to qor admin, so you could manage jobs in the admin interface
  Admin.AddResource(Worker)
}
```

## Things to note

- If a *Job* is scheduled within 2 minutes of the current time, then it will be run immediately.
- It is possible, via the admin interface, to abort a currently running *Job*: view the *Job*'s data via `Workers > Running` or `Workers > All Jobs` and press the `Abort running Job` button.
- It is possible, via the admin interface, to abort a scheduled *Job*: view the *Job*'s data via `Workers > Scheduled` or `Workers > All Jobs` and press the `Cancel scheduled Job` button.
- It is possible, via the admin interface, to update a scheduled *Job*, including setting a new date and time: view the *Job*'s data via `Workers > Scheduled` or `Workers > All Jobs`, update the `Schedule Time` field's value, and press the `Update scheduled Job` button. Please be aware that scheduling a *Job* to a date/time in the past will see the Job get run immediately.

[Worker Demo:  http://demo.getqor.com/admin/workers](http://demo.getqor.com/admin/workers)

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
