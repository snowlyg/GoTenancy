package excel

import (
	"strconv"
	"strings"

	"GoTenancy/libs/exchange"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// NewReader new csv reader
func (excel *Excel) NewReader(res *exchange.Resource, context *qor.Context) (exchange.Rows, error) {
	var rows = Rows{Excel: excel, Resource: res}

	readCloser, err := excel.getReader()
	if err == nil {
		defer readCloser.Close()
		var (
			reader, err = excelize.OpenReader(readCloser)
			sheetName   = excel.config.SheetName
		)

		if sheetName == "" {
			activeSheet := reader.GetActiveSheetIndex()
			if activeSheet == 0 && reader.SheetCount > 0 {
				activeSheet = 1
			}
			sheetName = reader.GetSheetName(activeSheet)
		}

		if err != nil {
			return nil, err
		}

		rows.records, err = reader.GetRows(sheetName)
		if err != nil {
			return nil, err
		}

		for i, r := range rows.records {
			for j, v := range r {
				rows.records[i][j] = strings.TrimSpace(v)
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
	*Excel
	Resource *exchange.Resource
	records  [][]string
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
