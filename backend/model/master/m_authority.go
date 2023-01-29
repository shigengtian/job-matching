package master

import "job-maching/global"

type Authority struct {
	global.MODEL
	AuthorityId     uint         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string       `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId        *uint        `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DataAuthorityId []*Authority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []Authority  `json:"children" gorm:"-"`
	SysBaseMenus    []Menu       `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []User       `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string       `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (Authority) TableName() string {
	return "m_authorities"
}
