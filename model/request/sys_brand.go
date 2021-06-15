package request

type CreateSysBrand struct {
	BrandName       string `json:"brandName" binding:"required"`            // 品牌名称
	Sort            int32  `json:"sort"`                                    // 排序
	Pic             string `json:"pic" binding:"required"`                  // 图标
	IsShow          int    `json:"isShow"`                                  // 是否显示
	BrandCategoryID int32  `json:"brandCategoryId" binding:"required,gt=0"` // 分类id
}

type UpdateSysBrand struct {
	Id              uint   `json:"id" form:"id" binding:"required,gt=0"`
	BrandName       string `json:"brandName" binding:"required"`            // 品牌名称
	Sort            int32  `json:"sort"`                                    // 排序
	Pic             string `json:"pic" binding:"required"`                  // 图标
	IsShow          int    `json:"isShow"`                                  // 是否显示
	BrandCategoryID int32  `json:"brandCategoryId" binding:"required,gt=0"` // 分类id
}

type SetSysBrand struct {
	Id              uint  `json:"id" form:"id" binding:"required,gt=0"`
	BrandCategoryID int32 `json:"brandCategoryId" binding:"required,gt=0"` // 分类id
}
