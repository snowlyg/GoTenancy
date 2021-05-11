package response

type LoginResponse struct {
	Token string `json:"AccessToken"`
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
	TenancyName   string `json:"tenancyName"`
}

type SysGeneralUser struct {
	TenancyResponse
	Username      string `json:"userName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	NickName      string `json:"nickName"`
	HeaderImg     string `json:"headerImg"`
	AuthorityName string `json:"authorityName"`
	AuthorityType int    `json:"authorityType"`
	TenancyName   string `json:"tenancyName"`
}
