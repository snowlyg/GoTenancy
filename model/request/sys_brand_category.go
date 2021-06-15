package request

type CreateSysBrandCategory struct {
	Pid      string `json:"pid" binding:"required"`       // 父id
	CateName string `json:"cateName"  binding:"required"` // 分类名称
	Path     string `json:"path"  `                       // 路径
	Sort     int32  `json:"sort"`                         // 排序
	Level    uint   `json:"level"`                        // 等级
	IsShow   int    `json:"isShow"`                       // 是否显示
}

type UpdateSysBrandCategory struct {
	Id       uint   `json:"id" form:"id" binding:"required,gt=0"`
	Pid      string `json:"pid" binding:"required"`       // 父id
	CateName string `json:"cateName"  binding:"required"` // 分类名称
	Path     string `json:"path"`                         // 路径
	Sort     int32  `json:"sort"`                         // 排序
	Level    uint   `json:"level"`                        // 等级
	IsShow   int    `json:"isShow"`                       // 是否显示
}
