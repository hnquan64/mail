package entities

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"int(11);primaryKey;autoIncrement" json:"id,omitempty"`
	Name      string         `gorm:"varchar(255);unique;not null;default:undefined" json:"name,omitempty"`
	CreatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	UpdatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"datetime(3)" json:"-"`

	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions,omitempty"`
}

type Permission struct {
	Id        uint           `gorm:"int(11);primaryKey;autoIncrement" json:"id,omitempty"`
	Name      string         `gorm:"varchar(255);unique;not null;default:undefined" json:"name,omitempty"`
	Status    string         `gorm:"enum('Active', 'Inactive'); default:Active" json:"status,omitempty"`
	CreatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	UpdatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"datetime(3)" json:"-"`

	Roles []Role `gorm:"many2many:role_permissions" json:"role_permissions"`
}
