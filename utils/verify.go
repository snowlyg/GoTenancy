package utils

import "github.com/snowlyg/go-tenancy/g"

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}, "AuthorityType": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}, "AuthorityType": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"ID": {NotEmpty()}, "AuthorityId": {NotEmpty()}}

	CreateTenancyVerify = Rules{"Name": {NotEmpty()}, "SysRegionCode": {NotEmpty()}}
	UpdateTenancyVerify = Rules{"ID": {NotEmpty()}, "Name": {NotEmpty()}, "SysRegionCode": {NotEmpty()}}
	SetRegionCodeVerify = Rules{"ID": {NotEmpty()}, "SysRegionCode": {NotEmpty()}}

	CreateMiniVerify = Rules{"Name": {NotEmpty()}, "AppID": {NotEmpty()}, "AppSecret": {NotEmpty()}}
	UpdateMiniVerify = Rules{"ID": {NotEmpty()}, "Name": {NotEmpty()}, "AppID": {NotEmpty()}, "AppSecret": {NotEmpty()}}
)

func GetLoginVerify() Rules {
	if g.TENANCY_CONFIG.System.Env == "dev" {
		return Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityType": {NotEmpty()}}
	}
	return Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityType": {NotEmpty()}}
}
