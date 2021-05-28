package request

// User register structure
type Register struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	AuthorityId string `json:"authorityId" validate:"required"`
}

// User login structure
type Login struct {
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	Captcha       string `json:"captcha" validate:"dev-required"`
	CaptchaId     string `json:"captchaId" validate:"dev-required"`
	AuthorityType int    `json:"authorityType" validate:"required,gt=0"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	NewPassword   string `json:"newPassword" validate:"required"`
	AuthorityType int    `json:"authorityType"  validate:"required,gt=0"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	Id          float64 `json:"id" form:"id" validate:"required,gt=0"`
	AuthorityId string  `json:"authorityId" validate:"required"`
}

// Modify  user's auth structure
type SetAdminInfo struct {
	Id        float64 `json:"id" form:"id"  validate:"required,gt=0"`
	Email     string  `json:"email"  validate:"required"`
	Phone     string  `json:"phone"  validate:"required"`
	Name      string  `json:"nickName"  validate:"required"`
	HeaderImg string  `json:"headerImg"  validate:"required"`
}
