package admin

import (
	"errors"

	"GoTenancy/libs/qor/resource"
)

// CollectionEditConfig meta configuration used for collection edit
type CollectionEditConfig struct {
	Template string
	Max      uint
	metaConfig
}

// GetTemplate get template for collection edit
func (collectionEditConfig CollectionEditConfig) GetTemplate(context *Context, metaType string) ([]byte, error) {
	if metaType == "form" && collectionEditConfig.Template != "" {
		return context.Asset(collectionEditConfig.Template)
	}
	return nil, errors.New("not implemented")
}

// ConfigureQorMeta configure collection edit meta
func (collectionEditConfig *CollectionEditConfig) ConfigureQorMeta(metaor resource.Metaor) {
	if meta, ok := metaor.(*Meta); ok {
		if meta.Permission != nil || meta.Resource.Permission != nil {
			meta.Permission = meta.Permission.Concat(meta.Resource.Permission)
		}
	}
}
