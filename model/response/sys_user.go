package response

type LoginResponse struct {
	User  interface{} `json:"user"`
	Token string      `json:"AccessToken"`
}

type SysAdminUser struct {
	TenancyResponse
	Username      string `json:"userName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	NickName      string `json:"nickName"`
	HeaderImg     string `json:"headerImg"`
	AuthorityName string `json:"authorityName"`
	AuthorityType int    `json:"authorityType"`
	AuthorityId   string `json:"authorityId"`
	DefaultRouter string `json:"defaultRouter"`
}

type SysTenancyUser struct {
	TenancyResponse
	Username      string `json:"userName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	NickName      string `json:"nickName"`
	HeaderImg     string `json:"headerImg"`
	AuthorityName string `json:"authorityName"`
	AuthorityType int    `json:"authorityType"`
	AuthorityId   string `json:"authorityId"`
	TenancyId     uint   `json:"tenancyId"`
	TenancyName   string `json:"tenancyName"`
	DefaultRouter string `json:"defaultRouter"`
}
