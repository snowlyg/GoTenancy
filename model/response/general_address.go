package response

import "github.com/snowlyg/go-tenancy/model"

type GeneralAddress struct {
	TenancyResponse
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Sex          model.Sex `json:"sex"`
	Country      string    `json:"country"`
	Province     string    `json:"province"`
	City         string    `json:"city"`
	District     string    `json:"district"`
	IsDefault    int       `json:"isDefault"`
	Detail       string    `json:"detail"`
	Postcode     string    `json:"postcode"`
	Age          int       `json:"age"`
	HospitalName string    `json:"hospitalName"`
	LocName      string    `json:"locName"`
	BedNum       string    `json:"bedNum"`
	HospitalNO   string    `json:"hospitalNo"`
	Disease      string    `json:"disease"`
}
