package models

import "time"

type Device struct {
	Name      string `gorm:"primaryKey"`
	Size      int64  `gorm:"not null"`
	UUID      string `gorm:"primaryKey"`
	CreatedAt time.Time
	Nodes     []*Node `gorm:"many2many:node_devices;"`
}
