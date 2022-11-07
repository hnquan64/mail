package entities

import (
	"time"

	"gitlab.com/meta-node/mail/core/pkgs/datatypes"
	"gorm.io/gorm"
)

type Email struct {
	ID       uint                  `gorm:"type:int(11);autoIncrement;primaryKey" json:"id,omitempty"`
	Title    string                `gorm:"type:text;not null" json:"title,omitempty"`
	Subject  string                `gorm:"type:text;not null" json:"subject,omitempty"`
	WriterID uint                  `gorm:"type:int(11); default:null" json:"writer_id,omitempty"`
	To       datatypes.StringArray `gorm:"type:text;not null" json:"to"`
	From     string                `gorm:"type:varchar(191);not null" json:"from"`
	Content  string                `gorm:"type:text" json:"content,omitempty"`
	Cc       datatypes.StringArray `gorm:"type:text" json:"cc,omitempty"`
	Bcc      datatypes.StringArray `gorm:"type:text" json:"bcc,omitempty"`
	Status   string                `gorm:"type:enum('Draft','Pending', 'Sending', 'Receiving', 'Approved', 'Sent', 'Seen', 'Declined', 'Dropped');default:Pending" json:"status"`
	Files    []File                `gorm:"foreignkey:EmailID" json:"files,omitempty"`

	CreatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	UpdatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"datetime(3)" json:"-"`
}

type File struct {
	ID      uint   `gorm:"type:int(11);primaryKey;autoIncrement" json:"id"`
	EmailID uint   `gorm:"type:int(11)" json:"email,omitempty"`
	Path    string `gorm:"type:text" json:"path,omitempty"`
}

type EmailManager struct {
	HandlerID uint           `gorm:"int(11);primaryKey" json:"hanlder_id,omitempty"`
	EmailID   uint           `gorm:"int(11);primaryKey" json:"email_id,omitempty"`
	Status    string         `gorm:"type:enum('Draft', 'Sending', 'Receiving', 'Approved', 'Sent', 'Seen', 'Declined', 'Dropped');default:Null" json:"status"`
	CreatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	UpdatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"datetime(3)" json:"-"`
}
