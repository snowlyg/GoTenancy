package publish

import (
	"fmt"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/roles"
	"GoTenancy/libs/worker"
)

type workerJobLogger struct {
	job worker.QorJobInterface
}

func (job workerJobLogger) Print(results ...interface{}) {
	job.job.AddLog(fmt.Sprint(results...))
}

// QorWorkerArgument used for qor publish job's argument
type QorWorkerArgument struct {
	IDs []string
	worker.Schedule
}

// SetWorker set publish's worker
func (publish *Publish) SetWorker(w *worker.Worker) {
	publish.WorkerScheduler = w
	publish.registerWorkerJob()
}

func (publish *Publish) registerWorkerJob() {
	if w := publish.WorkerScheduler; w != nil {
		if w.Admin == nil {
			fmt.Println("Need to add worker to admin first before set worker")
			return
		}

		qorWorkerArgumentResource := w.Admin.NewResource(&QorWorkerArgument{})
		qorWorkerArgumentResource.Meta(&admin.Meta{Name: "IDs", Type: "publish_job_argument", Valuer: func(record interface{}, context *qor.Context) interface{} {
			var values = map[*admin.Resource][][]string{}

			if workerArgument, ok := record.(*QorWorkerArgument); ok {
				for _, id := range workerArgument.IDs {
					if keys := strings.Split(id, "__"); len(keys) >= 2 {
						name, id := keys[0], keys[1:]
						recordRes := w.Admin.GetResource(name)
						values[recordRes] = append(values[recordRes], id)
					}
				}
			}

			return values
		}})

		w.RegisterJob(&worker.Job{
			Name:       "Publish",
			Group:      "Publish",
			Permission: roles.Deny(roles.Read, roles.Anyone),
			Handler: func(argument interface{}, job worker.QorJobInterface) error {
				if argu, ok := argument.(*QorWorkerArgument); ok {
					records := publish.searchWithPublishIDs(publish.DraftDB(), w.Admin, argu.IDs)
					publish.Logger(&workerJobLogger{job: job}).Publish(records...)
				}
				return nil
			},
			Resource: qorWorkerArgumentResource,
		})

		w.RegisterJob(&worker.Job{
			Name:       "Discard",
			Group:      "Publish",
			Permission: roles.Deny(roles.Read, roles.Anyone),
			Handler: func(argument interface{}, job worker.QorJobInterface) error {
				if argu, ok := argument.(*QorWorkerArgument); ok {
					records := publish.searchWithPublishIDs(publish.DraftDB(), w.Admin, argu.IDs)
					publish.Logger(&workerJobLogger{job: job}).Discard(records...)
				}
				return nil
			},
			Resource: qorWorkerArgumentResource,
		})
	}
}
