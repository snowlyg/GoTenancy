package admin

import (
	"fmt"
	"time"

	"github.com/qor/admin"
	"github.com/qor/i18n/exchange_actions"
	"github.com/qor/worker"
	"go-tenancy/config/i18n"
)

// SetupWorker setup worker
func SetupWorker(Tenant *admin.Admin) {
	Worker := worker.New()

	type sendNewsletterArgument struct {
		Subject      string
		Content      string `sql:"size:65532"`
		SendPassword string
		worker.Schedule
	}

	_ = Worker.RegisterJob(&worker.Job{
		Name: "Send Newsletter",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			_ = qorJob.AddLog("Started sending newsletters...")
			_ = qorJob.AddLog(fmt.Sprintf("Argument: %+v", argument.(*sendNewsletterArgument)))
			for i := 1; i <= 100; i++ {
				time.Sleep(100 * time.Millisecond)
				_ = qorJob.AddLog(fmt.Sprintf("Sending newsletter %v...", i))
				_ = qorJob.SetProgress(uint(i))
			}
			_ = qorJob.AddLog("Finished send newsletters")
			return nil
		},
		Resource: Tenant.NewResource(&sendNewsletterArgument{}),
	})

	exchange_actions.RegisterExchangeJobs(i18n.I18n, Worker)
	Tenant.AddResource(Worker, &admin.Config{Menu: []string{"Site Management"}, Priority: 3})
}
