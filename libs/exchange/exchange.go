package exchange

import (
	"fmt"
	"reflect"

	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/roles"
	"GoTenancy/libs/validations"
	"github.com/jinzhu/gorm"
)

// Resource defined an exchange resource, which includes importing/exporting fields definitions
type Resource struct {
	*resource.Resource
	Config *Config
	Metas  []*Meta
}

// Config is exchange resource config
type Config struct {
	// PrimaryField that used as primary field when searching resource from database
	PrimaryField string
	// Permission defined permission
	Permission *roles.Permission
	// WithoutHeader no header in the data file
	WithoutHeader      bool
	DisableTransaction bool
}

// NewResource new exchange Resource
func NewResource(value interface{}, config ...Config) *Resource {
	res := Resource{Resource: resource.New(value)}
	if len(config) > 0 {
		res.Config = &config[0]
	} else {
		res.Config = &Config{}
	}
	res.Permission = res.Config.Permission

	if res.Config.PrimaryField != "" {
		if err := res.SetPrimaryFields(res.Config.PrimaryField); err != nil {
			fmt.Println(err)
		}
	}
	return &res
}

// Meta define exporting/importing meta for exchange Resource
func (res *Resource) Meta(meta *Meta) *Meta {
	if meta.Header == "" {
		meta.Header = meta.Name
	}

	meta.base = res
	meta.updateMeta()
	res.Metas = append(res.Metas, meta)
	return meta
}

// GetMeta get defined Meta from exchange Resource
func (res *Resource) GetMeta(name string) *Meta {
	for _, meta := range res.Metas {
		if meta.Header == name {
			return meta
		}
	}
	return nil
}

// GetMetas get all defined Metas from exchange Resource
func (res *Resource) GetMetas([]string) []resource.Metaor {
	metas := []resource.Metaor{}
	for _, meta := range res.Metas {
		metas = append(metas, meta)
	}
	return metas
}

type errorsInterface interface {
	GetErrors() []error
}

// Import used to import data into a exchange Resource
//     product.Import(csv.New("products.csv"), context)
func (res *Resource) Import(container Container, context *qor.Context, callbacks ...func(Progress) error) error {
	rows, err := container.NewReader(res, context)
	if err == nil {
		var hasError bool
		var current uint
		var total = rows.Total()

		if db := context.GetDB(); db != nil && !res.Config.DisableTransaction {
			tx := db.Begin()
			context.SetDB(tx)
			defer func() {
				if hasError {
					tx.Rollback()
				} else {
					tx.Commit()
				}
			}()
		}

		for rows.Next() {
			current++
			progress := Progress{Total: total, Current: current}

			var metaValues *resource.MetaValues
			var handleError func(err error)

			if metaValues, err = rows.ReadRow(); err == nil {
				for _, metaValue := range metaValues.Values {
					progress.Cells = append(progress.Cells, Cell{
						Header: metaValue.Name,
						Value:  metaValue.Value,
					})
				}

				handleError = func(err error) {
					hasError = true
					progress.Errors.AddError(err)

					if errors, ok := err.(errorsInterface); ok {
						for _, err := range errors.GetErrors() {
							handleError(err)
						}
					} else if errValidations, ok := err.(*validations.Error); ok {
						for idx, cell := range progress.Cells {
							if cell.Header == errValidations.Column {
								cell.Error = errValidations
								progress.Cells[idx] = cell
								break
							}
						}
					} else if len(progress.Cells) > 0 {
						var err error = err
						cell := &progress.Cells[0]
						if cell.Error != nil {
							var errors qor.Errors
							errors.AddError(cell.Error)
							errors.AddError(err)
							err = errors
						}
						cell.Error = err
					}
				}

				result := res.NewStruct()
				progress.Value = result

				if err = res.FindOneHandler(result, metaValues, context); err == nil || err == gorm.ErrRecordNotFound {
					if err = resource.DecodeToResource(res, result, metaValues, context).Start(); err == nil {
						hasErr := context.HasError()
						if !hasErr {
							if err = res.CallSave(result, context); err != nil {
								handleError(err)
							}
						} else {
							handleError(context.Errors)
							clearContextErrorsForCurrentRow(context)
						}
					} else {
						handleError(err)
					}
				} else {
					handleError(err)
				}
			}

			for _, callback := range callbacks {
				if err := callback(progress); err != nil {
					return err
				}
			}
		}
	}
	return err
}
func clearContextErrorsForCurrentRow(context *qor.Context) {
	context.Errors = qor.Errors{}
}

// Export used export data from a exchange Resource
//     product.Export(csv.New("products.csv"), context)
func (res *Resource) Export(container Container, context *qor.Context, callbacks ...func(Progress) error) error {
	var (
		total   uint
		results = res.NewSlice()
		err     = context.GetDB().Find(results).Count(&total).Error
	)

	if err == nil {
		reflectValue := reflect.Indirect(reflect.ValueOf(results))

		writer, err := container.NewWriter(res, context)

		if err == nil {
			writer.WriteHeader()

			for i := 0; i < reflectValue.Len(); i++ {
				var result = reflectValue.Index(i).Interface()
				var metaValues *resource.MetaValues
				if metaValues, err = writer.WriteRow(result); err != nil {
					return err
				}

				var progress = Progress{
					Current: uint(i + 1),
					Total:   total,
					Value:   result,
				}

				for _, metaValue := range metaValues.Values {
					progress.Cells = append(progress.Cells, Cell{
						Header: metaValue.Name,
						Value:  metaValue.Value,
					})
				}

				for _, callback := range callbacks {
					if err := callback(progress); err != nil {
						return err
					}
				}
			}
			err = writer.Flush()
		}

		return err
	}

	return err
}
