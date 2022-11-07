package entities

import (
	"time"

	"gitlab.com/meta-node/mail/core/pkgs/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID          uint                  `gorm:"type:int(11);autoIncrement;primaryKey" json:"id"`
	UserName    string                `gorm:"type:varchar(255);not null" json:"user_name"`
	Email       string                `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password    string                `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNumber string                `gorm:"type:varchar(255);unique" json:"phone_number"`
	Avatar      string                `gorm:"type:text" json:"avatar"`
	Contacts    datatypes.StringArray `gorm:"type:text" json:"contacts"`
	MailStorage []Email               `gorm:"foreignKey:WriterID" json:"mail_storage"`

	CreatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	UpdatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"datetime(3)" json:"-"`
}
