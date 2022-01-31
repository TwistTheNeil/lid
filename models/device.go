package models

type Device struct {
	Name  string  `gorm:"primaryKey"`
	UUID  string  `gorm:"primaryKey"`
	Nodes []*Node `gorm:"many2many:node_devices;"`
}
