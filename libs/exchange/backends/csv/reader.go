package csv

import (
	"encoding/csv"
	"strconv"
	"strings"

	"GoTenancy/libs/exchange"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
)

// NewReader new csv reader
func (c *CSV) NewReader(res *exchange.Resource, context *qor.Context) (exchange.Rows, error) {
	var rows = Rows{CSV: c, Resource: res}

	readCloser, err := c.getReader()
	if err == nil {
		defer readCloser.Close()
		reader := csv.NewReader(readCloser)
		reader.TrimLeadingSpace = true

		rows.records, err = reader.ReadAll()

		if c.config.TrimSpace {
			for _, rows := range rows.records {
				for i, record := range rows {
					rows[i] = strings.TrimSpace(record)
				}
			}
		}

		rows.total = len(rows.records) - 1
		if res.Config.WithoutHeader {
			rows.total++
		}
	}

	return &rows, err
}

// Rows CSV rows struct
type Rows struct {
	*CSV
	Resource *exchange.Resource
	current  int
	total    int
}

// Header CSV header column
func (rows Rows) Header() (results []string) {
	if rows.total > 0 {
		if rows.Resource.Config.WithoutHeader {
			for i := 0; i <= len(rows.records[0]); i++ {
				results = append(results, strconv.Itoa(i))
			}
		} else {
			return rows.records[0]
		}
	}
	return
}

// Total CSV total rows
func (rows *Rows) Total() uint {
	return uint(rows.total)
}

// Next read next rows from CSV
func (rows *Rows) Next() bool {
	if rows.total >= rows.current+1 {
		rows.current++
		return true
	}
	return false
}

// ReadRow read row from CSV
func (rows Rows) ReadRow() (*resource.MetaValues, error) {
	var metaValues resource.MetaValues
	columns := rows.Header()

	for index, column := range columns {
		metaValue := resource.MetaValue{
			Name:  column,
			Value: rows.records[rows.current][index],
		}
		if meta := rows.Resource.GetMeta(column); meta != nil {
			metaValue.Meta = meta
			metaValue.Name = meta.Name
		}
		metaValues.Values = append(metaValues.Values, &metaValue)
	}

	return &metaValues, nil
}
