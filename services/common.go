package services

var (
	// 管理员
	IsAdmin = map[string]interface{}{"is_admin": 1}
	// 非管理员
	NotAdmin = map[string]interface{}{"is_admin": 0}
)
