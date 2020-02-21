package publish2

type Visible struct {
	PublishReady bool
}

func (visible Visible) GetPublishReady() bool {
	return visible.PublishReady
}

func (visible *Visible) SetPublishReady(b bool) {
	visible.PublishReady = b
}

type PublishReadyInterface interface {
	GetPublishReady() bool
	SetPublishReady(bool)
}
