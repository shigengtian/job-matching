package master

import "job-maching/global"

type Occupation struct {
	global.MODEL
	Occupation string `json:"occupation" gorm:"comment:職種"`
}

func (Occupation) TableName() string {
	return "m_industries"
}
