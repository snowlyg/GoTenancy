package utils

import "github.com/snowlyg/go-tenancy/g"

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"ID": {NotEmpty()}, "AuthorityId": {NotEmpty()}}

	CreateTenancyVerify = Rules{"Name": {NotEmpty()}, "SysRegionCode": {Ge("0")}}
	UpdateTenancyVerify = Rules{"ID": {NotEmpty()}, "Name": {NotEmpty()}, "SysRegionCode": {Ge("0")}}

	CreateMiniVerify = Rules{"Name": {NotEmpty()}, "AppID": {NotEmpty()}, "AppSecret": {NotEmpty()}}
	UpdateMiniVerify = Rules{"ID": {NotEmpty()}, "Name": {NotEmpty()}, "AppID": {NotEmpty()}, "AppSecret": {NotEmpty()}}
)

func GetLoginVerify() Rules {
	if g.TENANCY_CONFIG.System.Env == "dev" {
		return Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	}
	return Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
}
