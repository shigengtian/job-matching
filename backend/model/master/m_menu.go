package master

import "job-maching/global"

type Menu struct {
	global.MODEL
	MenuLevel     uint        `json:"-"`
	ParentId      string      `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path          string      `json:"path" gorm:"comment:路由path"`        // 路由path
	Name          string      `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden        bool        `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component     string      `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort          int         `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	SysAuthoritys []Authority `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []Menu      `json:"children" gorm:"-"`
}

func (Menu) TableName() string {
	return "m_menus"
}
