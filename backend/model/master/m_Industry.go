package master

import "job-maching/global"

type Industry struct {
	global.MODEL
	Industry string `json:"industry" gorm:"comment:業界"`
}

func (Industry) TableName() string {
	return "m_industries"
}
