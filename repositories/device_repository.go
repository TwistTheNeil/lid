package repsitories

type DeviceRepository interface {
	Create(string, string)
	FindByName(string)
	FindByUUID(string)
	FindAll()
	DeleteByName(string)
	DeleteByUUID(string)
}
