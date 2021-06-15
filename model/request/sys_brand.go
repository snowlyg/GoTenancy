package request

type BrandPageInfo struct {
	BrandCategoryId int32 `json:"brandCategoryId" form:"brandCategoryId"`
	Page            int   `json:"page" form:"page" binding:"required"`
	PageSize        int   `json:"pageSize" form:"pageSize" binding:"required"`
}
