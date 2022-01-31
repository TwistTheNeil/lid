package models

type DeviceRepository interface {
	Create(string, string)
	FindByName(string)
	FindByUUID(string)
	FindAll()
	DeleteByName(string)
	DeleteByUUID(string)
}
type Device struct {
	Name  string  `gorm:"primaryKey"`
	UUID  string  `gorm:"primaryKey"`
	Nodes []*Node `gorm:"many2many:node_devices;"`
}
