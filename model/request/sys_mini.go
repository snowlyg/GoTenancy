package request

type CreateSysMini struct {
	Name      string `json:"name" form:"name" validate:"required"`
	AppID     string `json:"appId" form:"appId" validate:"required"`
	AppSecret string `json:"appSecret" form:"appSecret" validate:"required"`
	Remark    string `json:"remark" form:"remark"`
}

type UpdateSysMini struct {
	Id        uint   `json:"id" form:"id" validate:"required,gt=0"`
	Name      string `json:"name" form:"name" validate:"required"`
	AppID     string `json:"appId" form:"appId" validate:"required"`
	AppSecret string `json:"appSecret" form:"appSecret" validate:"required"`
	Remark    string `json:"remark" form:"remark"`
}
