package request

// User register structure
type Register struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	AuthorityId string `json:"authorityId" binding:"required"`
}

// User login structure
type Login struct {
	Username      string `json:"username" binding:"required"`
	Password      string `json:"password" binding:"required"`
	Captcha       string `json:"captcha" binding:"dev-required"`
	CaptchaId     string `json:"captchaId" binding:"dev-required"`
	AuthorityType int    `json:"authorityType" binding:"required,gt=0"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username      string `json:"username" binding:"required"`
	Password      string `json:"password" binding:"required"`
	NewPassword   string `json:"newPassword" binding:"required"`
	AuthorityType int    `json:"authorityType"  binding:"required,gt=0"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	Id          float64 `json:"id" form:"id" binding:"required,gt=0"`
	AuthorityId string  `json:"authorityId" binding:"required"`
}

// Modify  user's auth structure
type SetAdminInfo struct {
	Id        float64 `json:"id" form:"id"  binding:"required,gt=0"`
	Email     string  `json:"email"  binding:"required"`
	Phone     string  `json:"phone"  binding:"required"`
	Name      string  `json:"nickName"  binding:"required"`
	HeaderImg string  `json:"headerImg"  binding:"required"`
}
