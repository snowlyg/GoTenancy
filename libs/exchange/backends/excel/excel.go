package excel

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// New new excel backend
func New(value interface{}, config ...*Config) *Excel {
	excel := &Excel{}

	if f, ok := value.(string); ok {
		excel.filename = f
	} else {
		if r, ok := value.(io.ReadCloser); ok {
			excel.reader = r
		}

		if w, ok := value.(io.WriteCloser); ok {
			excel.writer = w
		}
	}

	for _, cfg := range config {
		excel.config = cfg
	}

	if excel.config == nil {
		excel.config = &Config{}
	}

	return excel
}

// Config excel config
type Config struct {
	TrimSpace bool
	SheetName string
}

// Excel excel struct
type Excel struct {
	filename string
	reader   io.ReadCloser
	writer   io.WriteCloser
	config   *Config
}

func (excel *Excel) getReader() (io.ReadCloser, error) {
	if excel.reader != nil {
		return excel.reader, nil
	} else if excel.filename != "" {
		readerCloser, err := os.Open(excel.filename)
		return readerCloser, err
	}

	return nil, errors.New("Nothing available to import")
}

func (excel *Excel) getWriter() (*excelize.File, error) {
	if excel.filename != "" {
		dir := filepath.Dir(excel.filename)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
		}

		f := excelize.NewFile()

		return f, nil
	}

	if excel.writer != nil {
		return excelize.NewFile(), nil
	}

	return nil, errors.New("Nowhere to export")
}
