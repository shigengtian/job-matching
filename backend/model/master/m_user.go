package master

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primarykey;index;comment:用户UUID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Username    string         `json:"userName" gorm:"index;comment:用户登录名"`
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`
	HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	Authority   Authority      `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []Authority    `json:"authorities" gorm:"many2many:m_authorities;"`
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`
	Enable      bool           `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
}

func (User) TableName() string {
	return "m_users"
}
