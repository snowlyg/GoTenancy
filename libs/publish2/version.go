package publish2

var DefaultVersionName = "Default"

type Version struct {
	VersionName     string `gorm:"primary_key;size:128"`
	VersionPriority string `gorm:"index"`
}

func (version *Version) SetVersionName(name string) {
	version.VersionName = name
}

func (version Version) GetVersionName() string {
	return version.VersionName
}

type VersionableInterface interface {
	SetVersionName(string)
	GetVersionName() string
}

type SharedVersion struct {
	VersionName string `gorm:"primary_key;size:128"`
}

func (version *SharedVersion) SetSharedVersionName(name string) {
	version.VersionName = name
}

func (version SharedVersion) GetSharedVersionName() string {
	return version.VersionName
}

type ShareableVersionInterface interface {
	SetSharedVersionName(string)
	GetSharedVersionName() string
}
