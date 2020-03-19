package common

type Table struct {
	Code  int         `json:"code"` //0,1
	Msg   string      `json:"msg"`
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}

type ActionResponse struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type Pagination struct {
	Page  int `url:"page"`
	Limit int `url:"limit"`
}

type Id struct {
	Id int `json:"id"`
}
