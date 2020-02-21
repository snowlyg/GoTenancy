package seo

import (
	"GoTenancy/libs/l10n"
	"GoTenancy/libs/seo"
)

type MySEOSetting struct {
	seo.QorSEOSetting
	l10n.Locale
}

type SEOGlobalSetting struct {
	SiteName string
}

var SEOCollection *seo.Collection
