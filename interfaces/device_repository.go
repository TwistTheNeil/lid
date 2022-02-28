package interfaces

import "lid/models"

type DeviceRepository interface {
	Create(device_name string, device_uuid string, size int64) error
	FindByName(device_name string) (models.Device, error)
	FindByUUID(device_uuid string) (models.Device, error)
	FindAll() ([]models.Device, error)
	DeleteByName(device_name string) error
	DeleteByUUID(device_uuid string) error
}
