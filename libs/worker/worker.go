package worker

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/roles"
	"github.com/jinzhu/gorm"
)

const (
	// JobStatusScheduled job status scheduled
	JobStatusScheduled = "scheduled"
	// JobStatusCancelled job status cancelled
	JobStatusCancelled = "cancelled"
	// JobStatusNew job status new
	JobStatusNew = "new"
	// JobStatusRunning job status running
	JobStatusRunning = "running"
	// JobStatusDone job status done
	JobStatusDone = "done"
	// JobStatusException job status exception
	JobStatusException = "exception"
	// JobStatusKilled job status killed
	JobStatusKilled = "killed"
)

// New create Worker with Config
func New(config ...*Config) *Worker {
	var cfg = &Config{}
	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.Job == nil {
		cfg.Job = &QorJob{}
	}

	if cfg.Queue == nil {
		cfg.Queue = NewCronQueue()
	}

	return &Worker{Config: cfg}
}

// Config worker config
type Config struct {
	Queue Queue
	Job   QorJobInterface
	Admin *admin.Admin
}

// Worker worker definition
type Worker struct {
	*Config
	JobResource *admin.Resource
	Jobs        []*Job
	mounted     bool
}

// ConfigureQorResourceBeforeInitialize a method used to config Worker for qor admin
func (worker *Worker) ConfigureQorResourceBeforeInitialize(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		res.GetAdmin().RegisterViewPath("GoTenancy/libs/worker/views")
		res.UseTheme("worker")

		worker.Admin = res.GetAdmin()
		worker.JobResource = worker.Admin.NewResource(worker.Config.Job)
		worker.JobResource.UseTheme("worker")
		worker.JobResource.Meta(&admin.Meta{Name: "Name", Valuer: func(record interface{}, context *qor.Context) interface{} {
			return record.(QorJobInterface).GetJobName()
		}})
		worker.JobResource.IndexAttrs("ID", "Name", "Status", "CreatedAt")
		worker.JobResource.Name = res.Name

		for _, status := range []string{JobStatusScheduled, JobStatusNew, JobStatusRunning, JobStatusDone, JobStatusException} {
			var status = status
			worker.JobResource.Scope(&admin.Scope{Name: status, Handler: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
				return db.Where("status = ?", status)
			}})
		}

		// default scope
		worker.JobResource.Scope(&admin.Scope{
			Handler: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
				if jobName := ctx.Request.URL.Query().Get("job"); jobName != "" {
					return db.Where("kind = ?", jobName)
				}

				if groupName := ctx.Request.URL.Query().Get("group"); groupName != "" {
					var jobNames []string
					for _, job := range worker.Jobs {
						if groupName == job.Group {
							jobNames = append(jobNames, job.Name)
						}
					}
					if len(jobNames) > 0 {
						return db.Where("kind IN (?)", jobNames)
					}
					return db.Where("kind IS NULL")
				}

				{
					var jobNames []string
					for _, job := range worker.Jobs {
						jobNames = append(jobNames, job.Name)
					}
					if len(jobNames) > 0 {
						return db.Where("kind IN (?)", jobNames)
					}
				}

				return db
			},
			Default: true,
		})

		// Auto Migration
		worker.Admin.DB.AutoMigrate(worker.Config.Job)

		// Configure jobs
		for _, job := range worker.Jobs {
			if job.Resource == nil {
				job.Resource = worker.Admin.NewResource(worker.JobResource.Value)
			}
		}
	}
}

// ConfigureQorResource a method used to config Worker for qor admin
func (worker *Worker) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		// Parse job
		cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		qorJobID := cmdLine.String("qor-job", "", "Qor Job ID")
		runAnother := cmdLine.Bool("run-another", false, "Run another qor job")
		cmdLine.Parse(os.Args[1:])
		worker.mounted = true

		if *qorJobID != "" {
			if *runAnother == true {
				if newJob := worker.saveAnotherJob(*qorJobID); newJob != nil {
					newJobID := newJob.GetJobID()
					qorJobID = &newJobID
				} else {
					fmt.Println("failed to clone job " + *qorJobID)
					os.Exit(1)
				}
			}

			if err := worker.RunJob(*qorJobID); err == nil {
				os.Exit(0)
			} else {
				fmt.Println(err)
				// os.Exit(1)
			}
		}

		// register view funcmaps
		worker.Admin.RegisterFuncMap("get_grouped_jobs", func(worker *Worker, context *admin.Context) map[string][]*Job {
			var groupedJobs = map[string][]*Job{}
			var groupName = context.Request.URL.Query().Get("group")
			var jobName = context.Request.URL.Query().Get("job")
			for _, job := range worker.Jobs {
				if !(job.HasPermission(roles.Read, context.Context) && job.HasPermission(roles.Create, context.Context)) {
					continue
				}

				if (groupName == "" || groupName == job.Group) && (jobName == "" || jobName == job.Name) {
					groupedJobs[job.Group] = append(groupedJobs[job.Group], job)
				}
			}
			return groupedJobs
		})

		// configure routes
		router := worker.Admin.GetRouter()
		controller := workerController{Worker: worker}
		jobParamIDName := worker.JobResource.ParamIDName()

		router.Get(res.ToParam(), controller.Index, &admin.RouteConfig{Resource: worker.JobResource})
		router.Get(res.ToParam()+"/new", controller.New, &admin.RouteConfig{Resource: worker.JobResource})
		router.Get(fmt.Sprintf("%v/%v", res.ToParam(), jobParamIDName), controller.Show, &admin.RouteConfig{Resource: worker.JobResource})
		router.Get(fmt.Sprintf("%v/%v/edit", res.ToParam(), jobParamIDName), controller.Show, &admin.RouteConfig{Resource: worker.JobResource})
		router.Post(fmt.Sprintf("%v/%v/run", res.ToParam(), jobParamIDName), controller.RunJob, &admin.RouteConfig{Resource: worker.JobResource})
		router.Post(res.ToParam(), controller.AddJob, &admin.RouteConfig{Resource: worker.JobResource})
		router.Put(fmt.Sprintf("%v/%v", res.ToParam(), jobParamIDName), controller.Update, &admin.RouteConfig{Resource: worker.JobResource})
		router.Delete(fmt.Sprintf("%v/%v", res.ToParam(), jobParamIDName), controller.KillJob, &admin.RouteConfig{Resource: worker.JobResource})
	}
}

// SetQueue set worker's queue
func (worker *Worker) SetQueue(queue Queue) {
	worker.Queue = queue
}

// RegisterJob register a job into Worker
func (worker *Worker) RegisterJob(job *Job) error {
	if worker.mounted {
		debug.PrintStack()
		fmt.Printf("Job should be registered before Worker mounted into admin, but %v is registered after that", job.Name)
	}

	job.Worker = worker
	worker.Jobs = append(worker.Jobs, job)
	return nil
}

// GetRegisteredJob register a job into Worker
func (worker *Worker) GetRegisteredJob(name string) *Job {
	for _, job := range worker.Jobs {
		if job.Name == name {
			return job
		}
	}
	return nil
}

// GetJob get job with id
func (worker *Worker) GetJob(jobID string) (QorJobInterface, error) {
	qorJob := worker.JobResource.NewStruct().(QorJobInterface)

	context := worker.Admin.NewContext(nil, nil)
	context.ResourceID = jobID
	context.Resource = worker.JobResource

	if err := worker.JobResource.FindOneHandler(qorJob, nil, context.Context); err == nil {
		for _, job := range worker.Jobs {
			if job.Name == qorJob.GetJobName() {
				qorJob.SetJob(job)
				return qorJob, nil
			}
		}
		return nil, fmt.Errorf("failed to load job: %v, unregistered job type: %v", jobID, qorJob.GetJobName())
	}
	return nil, fmt.Errorf("failed to find job: %v", jobID)
}

// AddJob add job to worker
func (worker *Worker) AddJob(qorJob QorJobInterface) error {
	return worker.Queue.Add(qorJob)
}

// RunJob run job with job id
func (worker *Worker) RunJob(jobID string) error {
	qorJob, err := worker.GetJob(jobID)

	if qorJob != nil && err == nil {
		defer func() {
			if r := recover(); r != nil {
				qorJob.AddLog(string(debug.Stack()))
				qorJob.SetProgressText(fmt.Sprint(r))
				qorJob.SetStatus(JobStatusException)
			}
		}()

		if qorJob.GetStatus() != JobStatusNew && qorJob.GetStatus() != JobStatusScheduled {
			return errors.New("invalid job status, current status: " + qorJob.GetStatus())
		}

		if err = qorJob.SetStatus(JobStatusRunning); err == nil {
			if err = qorJob.GetJob().GetQueue().Run(qorJob); err == nil {
				if len(qorJob.GetResultsTable().TableCells) > 0 {
					return qorJob.SetStatus(JobStatusException)
				}
				return qorJob.SetStatus(JobStatusDone)
			}

			qorJob.SetProgressText(err.Error())
			qorJob.SetStatus(JobStatusException)
		}
	}

	return err
}

func (worker *Worker) saveAnotherJob(jobID string) QorJobInterface {
	jobResource := worker.JobResource
	newJob := jobResource.NewStruct().(QorJobInterface)

	job, err := worker.GetJob(jobID)
	if err == nil {
		newJob.SetJob(job.GetJob())
		newJob.SetSerializableArgumentValue(job.GetArgument())
		context := worker.Admin.NewContext(nil, nil)
		if err := jobResource.CallSave(newJob, context.Context); err == nil {
			return newJob
		}
	}
	return nil
}

// KillJob kill job with job id
func (worker *Worker) KillJob(jobID string) error {
	if qorJob, err := worker.GetJob(jobID); err == nil {
		if qorJob.GetStatus() == JobStatusRunning {
			if err = qorJob.GetJob().GetQueue().Kill(qorJob); err == nil {
				qorJob.SetStatus(JobStatusKilled)
				return nil
			}
			return err
		} else if qorJob.GetStatus() == JobStatusScheduled || qorJob.GetStatus() == JobStatusNew {
			qorJob.SetStatus(JobStatusKilled)
			return worker.RemoveJob(jobID)
		} else {
			return errors.New("invalid job status")
		}
	} else {
		return err
	}
}

// RemoveJob remove job with job id
func (worker *Worker) RemoveJob(jobID string) error {
	qorJob, err := worker.GetJob(jobID)
	if err == nil {
		return qorJob.GetJob().GetQueue().Remove(qorJob)
	}
	return err
}
