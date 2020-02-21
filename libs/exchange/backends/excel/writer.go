package excel

import (
	"GoTenancy/libs/exchange"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/roles"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// NewWriter new csv writer
func (excel *Excel) NewWriter(res *exchange.Resource, context *qor.Context) (exchange.Writer, error) {
	writer := &Writer{Excel: excel, Resource: res, context: context, sheetName: excel.config.SheetName}

	var metas []*exchange.Meta
	for _, meta := range res.Metas {
		if meta.HasPermission(roles.Read, context) {
			metas = append(metas, meta)
		}
	}

	writer.metas = metas

	excelWriter, err := excel.getWriter()

	if err == nil {
		if writer.sheetName == "" {
			writer.sheetName = "Sheet1"
		}

		if !excelWriter.GetSheetVisible(writer.sheetName) {
			idx := excelWriter.NewSheet(writer.sheetName)
			excelWriter.SetActiveSheet(idx)
		}

		writer.Writer = excelWriter
	}

	return writer, err
}

// Writer CSV writer struct
type Writer struct {
	*Excel
	currentRow int
	sheetName  string
	context    *qor.Context
	metas      []*exchange.Meta
	Resource   *exchange.Resource
	Writer     *excelize.File
}

func toAxis(x, y int) string {
	col, err := excelize.ColumnNumberToName(x + 1)
	// must not have any err
	if err != nil {
		panic(err)
	}

	// must not have any err
	ax, err := excelize.JoinCellName(col, y)
	if err != nil {
		panic(err)
	}

	return ax
}

// WriteHeader write header
func (writer *Writer) WriteHeader() error {
	if !writer.Resource.Config.WithoutHeader {
		writer.currentRow++
		for key, meta := range writer.metas {
			writer.Writer.SetCellValue(writer.sheetName, toAxis(key, writer.currentRow), meta.Header)
		}
	}
	return nil
}

// WriteRow write row
func (writer *Writer) WriteRow(record interface{}) (*resource.MetaValues, error) {
	var metaValues resource.MetaValues
	writer.currentRow++

	for key, meta := range writer.metas {
		value := meta.GetFormattedValuer()(record, writer.context)
		metaValue := resource.MetaValue{
			Name:  meta.GetName(),
			Value: value,
		}

		metaValues.Values = append(metaValues.Values, &metaValue)
		writer.Writer.SetCellValue(writer.sheetName, toAxis(key, writer.currentRow), value)
	}

	return &metaValues, nil
}

// Flush flush all changes
func (writer *Writer) Flush() error {
	if writer.Excel.writer != nil {
		defer writer.Excel.writer.Close()
		return writer.Writer.Write(writer.Excel.writer)
	}

	return writer.Writer.SaveAs(writer.Excel.filename)
}
