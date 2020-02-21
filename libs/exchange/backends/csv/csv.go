package csv

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// New initialize CSV backend, config is option, the last one will be used if there are more than one configs
func New(value interface{}, config ...Config) *CSV {
	csv := &CSV{}

	if f, ok := value.(string); ok {
		csv.filename = f
	} else {
		if r, ok := value.(io.ReadCloser); ok {
			csv.reader = r
		}

		if w, ok := value.(io.WriteCloser); ok {
			csv.writer = w
		}
	}

	for _, cfg := range config {
		csv.config = cfg
	}
	return csv
}

// Config CSV exchange config
type Config struct {
	TrimSpace bool
}

// CSV CSV struct
type CSV struct {
	config  Config
	records [][]string

	filename string
	reader   io.ReadCloser
	writer   io.WriteCloser
}

func (c *CSV) getReader() (io.ReadCloser, error) {
	if c.reader != nil {
		return c.reader, nil
	} else if c.filename != "" {
		if isXLSXFile(c.filename) {
			return generateCSVFromXLSXFile(c.filename)
		}

		readerCloser, err := os.Open(c.filename)
		return readerCloser, err
	}

	return nil, errors.New("Nothing available to import")
}

func (c *CSV) getWriter() (io.WriteCloser, error) {
	if c.writer != nil {
		return c.writer, nil
	} else if c.filename != "" {
		dir := filepath.Dir(c.filename)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
		}
		writerCloser, err := os.OpenFile(c.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

		return writerCloser, err
	}

	return nil, errors.New("Nowhere to export")
}
