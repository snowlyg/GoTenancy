package response

type SysBrandCategory struct {
	TenancyResponse
	Pid      string             `json:"pid"`      // 父id
	CateName string             `json:"cateName"` // 分类名称
	Path     string             `json:"path"`     // 路径
	Sort     int32              `json:"sort"`     // 排序
	Level    uint               `json:"level"`    // 等级
	IsShow   bool               `json:"isShow"`   // 是否显示
	Children []SysBrandCategory `json:"children" gorm:"-"`
}
