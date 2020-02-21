package blogs

import (
	"GoTenancy/libs/page_builder"
	"GoTenancy/libs/publish2"
)

type Page struct {
	page_builder.Page

	publish2.Version
	publish2.Schedule
	publish2.Visible
}
