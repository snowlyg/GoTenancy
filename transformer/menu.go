package transformer

type TableMenu struct {
	Id          int64  `json:"authorityId"`
	Title       string `json:"authorityName"`
	OrderNumber int    `json:"orderNumber"`
	Href        string `json:"menuUrl"`
	Icon        string `json:"menuIcon"`
	Authority   string `json:"authority"`
	CreatedAt   string `json:"createTime"`
	UpdatedAt   string `json:"updateTime"`
	Checked     int    `json:"checked"`
	IsMenu      int    `json:"isMenu"`
	ParentId    int64  `json:"parentId"`
}
