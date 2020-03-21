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

// Response 接口响应数据
type Response struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// 错误页面信息
var ErrorStatus = map[int]map[string]interface{}{
	500: {
		"one":   5,
		"two":   0,
		"three": 0,
		"msg":   "不好！有什么不好的事情在发生！！",
	},
	404: {
		"one":   4,
		"two":   0,
		"three": 4,
		"msg":   "不好！这个网页走丢了！！",
	},
	403: {
		"one":   4,
		"two":   0,
		"three": 3,
		"msg":   "不好！你的钥匙好像有点不对！！",
	},
}
