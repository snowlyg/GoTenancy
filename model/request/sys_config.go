package request

type GetByConfigCate struct {
	Cate string `json:"category" uri:"category" form:"category" binding:"required"`
}
