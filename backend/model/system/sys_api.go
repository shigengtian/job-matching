package system

import "job-maching/global"

type SysApi struct {
	global.GB_MODEL
	Path        string `json:"path" gorm:"comment:api path"`
	Description string `json:"description" gorm:"comment:api description"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api group"`
	Method      string `json:"method" gorm:"default:POST;comment:method"`
}

func (SysApi) TableName() string {
	return "sys_apis"
}
