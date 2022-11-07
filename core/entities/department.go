package entities

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID        uint           `gorm:"type:int(11);autoIncrement;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"department"`
	CreatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	UpdatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"datetime(3);default:null" json:"-"`
}
