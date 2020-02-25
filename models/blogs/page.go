package blogs

import (
	"github.com/qorpage_builder"
	"github.com/qorpublish2"
)

type Page struct {
	page_builder.Page

	publish2.Version
	publish2.Schedule
	publish2.Visible
}
