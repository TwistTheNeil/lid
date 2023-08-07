package models

import "time"

type Device struct {
	Name      string    `gorm:"primaryKey" json:"name"`
	Size      int64     `gorm:"not null" json:"size"`
	UUID      string    `gorm:"primaryKey" json:"uuid"`
	CreatedAt time.Time `json:"-"`
	Files     []*File   `gorm:"many2many:file_devices;" json:"files,omitempty"`
}
