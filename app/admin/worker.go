package admin

import (
	"fmt"
	"path/filepath"
	"time"

	"GoTenancy/config/db"
	"GoTenancy/config/i18n"
	"GoTenancy/models/products"
	"github.com/qor/admin"
	"github.com/qor/exchange"
	"github.com/qor/exchange/backends/csv"
	"github.com/qor/i18n/exchange_actions"
	"github.com/qor/media/oss"
	"github.com/qor/qor"
	"github.com/qor/worker"
)

// SetupWorker setup worker
func SetupWorker(Admin *admin.Admin) {
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
		Resource: Admin.NewResource(&sendNewsletterArgument{}),
	})

	type importProductArgument struct {
		File oss.OSS
	}

	_ = Worker.RegisterJob(&worker.Job{
		Name:  "Import Products",
		Group: "Products Management",
		Handler: func(arg interface{}, qorJob worker.QorJobInterface) error {
			argument := arg.(*importProductArgument)

			context := &qor.Context{DB: db.DB}

			var errorCount uint

			if err := ProductExchange.Import(
				csv.New(filepath.Join("public", argument.File.URL())),
				context,
				func(progress exchange.Progress) error {
					var cells = []worker.TableCell{
						{Value: fmt.Sprint(progress.Current)},
					}

					var hasError bool
					for _, cell := range progress.Cells {
						var tableCell = worker.TableCell{
							Value: fmt.Sprint(cell.Value),
						}

						if cell.Error != nil {
							hasError = true
							errorCount++
							tableCell.Error = cell.Error.Error()
						}

						cells = append(cells, tableCell)
					}

					if hasError {
						if errorCount == 1 {
							var headerCells = []worker.TableCell{
								{Value: "Line No."},
							}
							for _, cell := range progress.Cells {
								headerCells = append(headerCells, worker.TableCell{
									Value: cell.Header,
								})
							}
							_ = qorJob.AddResultsRow(headerCells...)
						}

						_ = qorJob.AddResultsRow(cells...)
					}

					_ = qorJob.SetProgress(uint(float32(progress.Current) / float32(progress.Total) * 100))
					_ = qorJob.AddLog(fmt.Sprintf("%d/%d Importing product %v", progress.Current, progress.Total, progress.Value.(*products.Product).Code))
					return nil
				},
			); err != nil {
				_ = qorJob.AddLog(err.Error())
			}

			return nil
		},
		Resource: Admin.NewResource(&importProductArgument{}),
	})

	_ = Worker.RegisterJob(&worker.Job{
		Name:  "Export Products",
		Group: "Products Management",
		Handler: func(arg interface{}, qorJob worker.QorJobInterface) error {
			_ = qorJob.AddLog("Exporting products...")

			context := &qor.Context{DB: db.DB}
			fileName := fmt.Sprintf("/downloads/products.%v.csv", time.Now().UnixNano())
			if err := ProductExchange.Export(
				csv.New(filepath.Join("public", fileName)),
				context,
				func(progress exchange.Progress) error {
					_ = qorJob.AddLog(fmt.Sprintf("%v/%v Exporting product %v", progress.Current, progress.Total, progress.Value.(*products.Product).Code))
					return nil
				},
			); err != nil {
				qorJob.AddLog(err.Error())
			}

			_ = qorJob.SetProgressText(fmt.Sprintf("<a href='%v'>Download exported products</a>", fileName))
			return nil
		},
	})

	exchange_actions.RegisterExchangeJobs(i18n.I18n, Worker)
	Admin.AddResource(Worker, &admin.Config{Menu: []string{"Site Management"}, Priority: 3})
}
