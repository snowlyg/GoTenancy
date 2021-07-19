package request

type UserLabelPageInfo struct {
	Page      int `json:"page" form:"page" binding:"required"`
	PageSize  int `json:"pageSize" form:"pageSize" binding:"required"`
	LabelType int `json:"labelType" form:"labelType" `
}
