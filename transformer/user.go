package transformer

type UserTable struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Telphone  string `json:"telphone"`
	RoleNames string `json:"role_names"`
	CreatedAt string `json:"created_at"`
}

type UserUpdate struct {
	Name     string `json:"name" validate:"gte=6,lte=50"  comment:"姓名" `
	Username string `json:"username" validate:"required,gte=6,lte=12" comment:"用户名" `
	Email    string `json:"email" validate:"email" comment:"邮箱"`
	Telphone string `json:"telphone" `
	RoleIds  string `json:"role_ids" validate:"required"`
}
