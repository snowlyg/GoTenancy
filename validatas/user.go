package validatas

type LoginInput struct {
	Username string `json:"username" validate:"required,gte=6,lte=50" comment:"用户名"`
	Password string `json:"password" validate:"required,gte=6,lte=12"  comment:"密码"`
}
