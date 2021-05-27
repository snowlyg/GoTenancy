package request

type CreateTenancyCategory struct {
	Pid      string `json:"pid" validate:"required"`       // 父id
	CateName string `json:"cateName"  validate:"required"` // 分类名称
	Path     string `json:"path"  `                        // 路径
	Sort     int32  `json:"sort"`                          // 排序
	Level    uint   `json:"level"`                         // 等级
	Pic      string `json:"pic"`                           // 图标
	IsShow   bool   `json:"isShow"`                        // 是否显示
}

type UpdateTenancyCategory struct {
	Id       uint   `json:"id" form:"id" validate:"required,gt=0"`
	Pid      string `json:"pid" validate:"required"`       // 父id
	CateName string `json:"cateName"  validate:"required"` // 分类名称
	Path     string `json:"path"`                          // 路径
	Sort     int32  `json:"sort"`                          // 排序
	Level    uint   `json:"level"`                         // 等级
	Pic      string `json:"pic"`                           // 图标
	IsShow   bool   `json:"isShow"`                        // 是否显示
}