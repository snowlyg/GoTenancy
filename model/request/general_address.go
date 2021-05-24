package request

import (
	"github.com/snowlyg/go-tenancy/model"
)

type CreateAddress struct {
	Name         string    `json:"name" validate:"required"`
	Phone        string    `json:"phone" validate:"required"`
	Sex          model.Sex `json:"sex" validate:"required"`
	Country      string    `json:"country" validate:"required"`
	Province     string    `json:"province" validate:"required"`
	City         string    `json:"city"`
	District     string    `json:"district" validate:"required"`
	IsDefault    bool      `json:"isDefault"`
	Detail       string    `json:"detail" validate:"required"`
	Postcode     string    `json:"postcode"`
	Age          int       `json:"age"`
	HospitalName string    `json:"hospitalName"`
	LocName      string    `json:"locName"`
	BedNum       string    `json:"bedNum"`
	HospitalNO   string    `json:"hospitalNo"`
	Disease      string    `json:"disease"`
}

type UpdateAddress struct {
	Id           uint      `json:"id" form:"id" validate:"required,gt=0"`
	Name         string    `json:"name" validate:"required"`
	Phone        string    `json:"phone" validate:"required"`
	Sex          model.Sex `json:"sex" validate:"required"`
	Country      string    `json:"country" validate:"required"`
	Province     string    `json:"province" validate:"required"`
	City         string    `json:"city"`
	District     string    `json:"district" validate:"required"`
	IsDefault    bool      `json:"isDefault" validate:"required"`
	Detail       string    `json:"detail" validate:"required"`
	Postcode     string    `json:"postcode"`
	Age          int       `json:"age"`
	HospitalName string    `json:"hospitalName"`
	LocName      string    `json:"locName"`
	BedNum       string    `json:"bedNum"`
	HospitalNO   string    `json:"hospitalNo"`
	Disease      string    `json:"disease"`
}
