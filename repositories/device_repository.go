package repsitories

type DeviceRepository interface {
	Create(device_name string, device_uuid string)
	FindByName(name string)
	FindByUUID(uuid string)
	FindAll()
	DeleteByName(device_name string)
	DeleteByUUID(device_uuid string)
}
