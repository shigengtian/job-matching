package master

import "job-maching/global"

type ProgramLanguage struct {
	global.MODEL
	ProgramLanguage string `json:"programLanguage" gorm:"comment:開発言語;not null;" `
}

func (ProgramLanguage) TableName() string {
	return "m_program_languages"
}
