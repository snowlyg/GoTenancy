package request

type SetNowMoney struct {
	NowMoney float64 `json:"nowMoney"`
	Type     int     `json:"type" binding:"required"` // 1:+ ,2:-
}

type UserPageInfo struct {
	Page         int    `json:"page" form:"page" binding:"required"`
	PageSize     int    `json:"pageSize" form:"pageSize" binding:"required"`
	GroupId      string `json:"groupId" form:"groupId"`
	LabelId      string `json:"labelId" form:"labelId"`
	Sex          string `json:"sex" form:"sex"`
	Country      string `json:"country" form:"country"`
	NickName     string `json:"nickName" form:"nickName"`
	UserTime     string `json:"userTime" form:"userTime"`
	UserTimeType string `json:"userTimeType" form:"userTimeType"`
	UserType     string `json:"userType" form:"userType"`
	PayCount     string `json:"payCount" form:"payCount"`
}
