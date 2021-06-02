package request

// api分页条件查询及排序结构体
type SearchApiParams struct {
	Path        string `json:"path"`
	Description string `json:"description"`
	ApiGroup    string `json:"apiGroup"`
	Method      string `json:"method"`
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
type DeleteApi struct {
	Id     uint   `json:"id" form:"id" binding:"required,gt=0"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}
