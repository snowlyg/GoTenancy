package model

type SysMenu struct {
	SysBaseMenu
	MenuId      uint      `json:"menu_id" gorm:"comment:菜单ID"`
	AuthorityId string    `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu `json:"children" gorm:"-"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}
