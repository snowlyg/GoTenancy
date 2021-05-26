package request

type CreateSysBrand struct {
	BrandName       string `json:"brandName" validate:"required"`            // 品牌名称
	Sort            int32  `json:"sort" validate:"numeric,gte=0"`            // 排序
	Pic             string `json:"pic" validate:"required"`                  // 图标
	IsShow          bool   `json:"isShow"`                                   // 是否显示
	BrandCategoryID int32  `json:"brandCategoryId" validate:"required,gt=0"` // 分类id
}

type UpdateSysBrand struct {
	Id              uint   `json:"id" form:"id" validate:"required,gt=0"`
	BrandName       string `json:"brandName" validate:"required"`            // 品牌名称
	Sort            int32  `json:"sort" validate:"numeric,gte=0"`            // 排序
	Pic             string `json:"pic" validate:"required"`                  // 图标
	IsShow          bool   `json:"isShow"`                                   // 是否显示
	BrandCategoryID int32  `json:"brandCategoryId" validate:"required,gt=0"` // 分类id
}
