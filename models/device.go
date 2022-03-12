package models

import "time"

type Device struct {
	Name      string    `gorm:"primaryKey" json:"name"`
	Size      int64     `gorm:"not null" json:"size"`
	UUID      string    `gorm:"primaryKey" json:"uuid"`
	CreatedAt time.Time `json:"-"`
	Nodes     []*Node   `gorm:"many2many:node_devices;" json:"files"`
}
