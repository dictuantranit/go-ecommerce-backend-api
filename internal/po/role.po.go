package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment:'primary Key is ID'"`
	RoleName string    `gorm:"column:role_name"`
	RoleNote string    `gorm:"column:role_note; type:text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
