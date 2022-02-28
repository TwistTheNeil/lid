package models

type Device struct {
	Name  string  `gorm:"primaryKey"`
	Size  int64   `gorm:"not null"`
	UUID  string  `gorm:"primaryKey"`
	Nodes []*Node `gorm:"many2many:node_devices;"`
}
