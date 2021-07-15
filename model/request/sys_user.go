package request

// User register structure
type Register struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	AuthorityId string `json:"authorityId" binding:"required"`
}

// User login structure
type Login struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"dev-required"`
	CaptchaId string `json:"captchaId" binding:"dev-required"`
}

// Modify password structure
type ChangePassword struct {
	Password        string `json:"password" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

// Modify password structure
type ChangeProfile struct {
	NickName string `json:"nickName" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	Id          uint   `json:"id" form:"id" binding:"required,gt=0"`
	AuthorityId string `json:"authorityId" binding:"required"`
}

// Modify  user's auth structure
type SetAdminInfo struct {
	Id        uint   `json:"id" form:"id"  binding:"required,gt=0"`
	Email     string `json:"email"  binding:"required"`
	Phone     string `json:"phone"  binding:"required"`
	Name      string `json:"nickName"  binding:"required"`
	HeaderImg string `json:"headerImg"  binding:"required"`
}

type UserPageInfo struct {
	Page         int    `json:"page" form:"page" binding:"required"`
	PageSize     int    `json:"pageSize" form:"pageSize" binding:"required"`
	GroupId      string `json:"groupId" form:"groupId"`
	LabelId      string `json:"labelId" form:"labelId"`
	Sex          string `json:"sex" form:"sex"`
	Country      string `json:"country" form:"country"`
	NickName     string `json:"nickName" form:"nickName"`
	UserTime     string `json:"userTime" form:"userTime"`
	UserTimeType string `json:"userTimeType" form:"userTimeType"`
	UserType     string `json:"userType" form:"userType"`
	PayCount     string `json:"payCount" form:"payCount"`
}
