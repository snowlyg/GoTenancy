package exchange

import (
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
)

// Container is an interface, any exporting/importing backends needs to implement this
type Container interface {
	NewReader(*Resource, *qor.Context) (Rows, error)
	NewWriter(*Resource, *qor.Context) (Writer, error)
}

// Rows is an interface, backends need to implement this in order to read data from it
type Rows interface {
	Header() []string
	ReadRow() (*resource.MetaValues, error)
	Next() bool
	Total() uint
}

// Writer is an interface, backends need to implement this in order to write data
type Writer interface {
	WriteHeader() error
	WriteRow(interface{}) (*resource.MetaValues, error)
	Flush() error
}
