package request

type CreateAddress struct {
	Name         string `json:"name" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Sex          int    `json:"sex" binding:"required"`
	Country      string `json:"country" binding:"required"`
	Province     string `json:"province" binding:"required"`
	City         string `json:"city" binding:"required"`
	District     string `json:"district" binding:"required"`
	Detail       string `json:"detail" binding:"required"`
	IsDefault    int    `json:"isDefault"`
	Postcode     string `json:"postcode"`
	Age          int    `json:"age"`
	HospitalName string `json:"hospitalName"`
	LocName      string `json:"locName"`
	BedNum       string `json:"bedNum"`
	HospitalNO   string `json:"hospitalNo"`
	Disease      string `json:"disease"`
}

type UpdateAddress struct {
	Id           uint   `json:"id" form:"id" binding:"required,gt=0"`
	Name         string `json:"name" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Sex          int    `json:"sex" binding:"required"`
	Country      string `json:"country" binding:"required"`
	Province     string `json:"province" binding:"required"`
	City         string `json:"city"`
	District     string `json:"district" binding:"required"`
	IsDefault    int    `json:"isDefault"`
	Detail       string `json:"detail" binding:"required"`
	Postcode     string `json:"postcode"`
	Age          int    `json:"age"`
	HospitalName string `json:"hospitalName"`
	LocName      string `json:"locName"`
	BedNum       string `json:"bedNum"`
	HospitalNO   string `json:"hospitalNo"`
	Disease      string `json:"disease"`
}
