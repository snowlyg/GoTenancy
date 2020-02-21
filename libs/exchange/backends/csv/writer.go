package csv

import (
	"encoding/csv"
	"fmt"

	"GoTenancy/libs/exchange"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/roles"
)

// NewWriter new csv writer
func (c *CSV) NewWriter(res *exchange.Resource, context *qor.Context) (exchange.Writer, error) {
	writer := &Writer{CSV: c, Resource: res, context: context}

	var metas []*exchange.Meta
	for _, meta := range res.Metas {
		if meta.HasPermission(roles.Read, context) {
			metas = append(metas, meta)
		}
	}
	writer.metas = metas

	csvWriter, err := c.getWriter()

	if err == nil {
		writer.Writer = csv.NewWriter(csvWriter)
	}

	return writer, err
}

// Writer CSV writer struct
type Writer struct {
	*CSV
	context  *qor.Context
	Resource *exchange.Resource
	Writer   *csv.Writer
	metas    []*exchange.Meta
}

// WriteHeader write header
func (writer *Writer) WriteHeader() error {
	if !writer.Resource.Config.WithoutHeader {
		var results []string
		for _, meta := range writer.metas {
			results = append(results, meta.Header)
		}
		writer.Writer.Write(results)
	}
	return nil
}

// WriteRow write row
func (writer *Writer) WriteRow(record interface{}) (*resource.MetaValues, error) {
	var metaValues resource.MetaValues
	var results []string

	for _, meta := range writer.metas {
		value := meta.GetFormattedValuer()(record, writer.context)
		metaValue := resource.MetaValue{
			Name:  meta.GetName(),
			Value: value,
		}

		metaValues.Values = append(metaValues.Values, &metaValue)
		results = append(results, fmt.Sprint(value))
	}

	return &metaValues, writer.Writer.Write(results)
}

// Flush flush all changes
func (writer *Writer) Flush() error {
	writer.Writer.Flush()
	return nil
}
