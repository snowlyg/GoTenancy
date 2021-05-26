package request

type CreateSysBrandCategory struct {
	Pid      int32  `json:"pid" validate:"numeric,gte=0"`    // 父id
	CateName string `json:"cateName"  validate:"required"`   // 分类名称
	Path     string `json:"path"  validate:"required"`       // 路径
	Sort     int32  `json:"sort"  validate:"numeric,gte=0"`  // 排序
	Level    uint   `json:"level"  validate:"numeric,gte=0"` // 等级
	IsShow   bool   `json:"isShow"`                          // 是否显示
}

type UpdateSysBrandCategory struct {
	Id       uint   `json:"id" form:"id" validate:"required,gt=0"`
	Pid      int32  `json:"pid" validate:"numeric,gte=0"`    // 父id
	CateName string `json:"cateName"  validate:"required"`   // 分类名称
	Path     string `json:"path"  validate:"required"`       // 路径
	Sort     int32  `json:"sort"  validate:"numeric,gte=0"`  // 排序
	Level    uint   `json:"level"  validate:"numeric,gte=0"` // 等级
	IsShow   bool   `json:"isShow"`                          // 是否显示
}
