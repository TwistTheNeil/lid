package models

type Host struct {
	Name    string   `gorm:"primaryKey" json:"name"`
	Devices []Device `gorm:"foreignKey:UUID" json:"devices"`
}
