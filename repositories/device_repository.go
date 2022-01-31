package repositories

import (
	"lid/models"

	"gorm.io/gorm"
)

type DeviceRepository interface {
	Create(device_name string, device_uuid string) error
	FindByName(device_name string) (models.Device, error)
	FindByUUID(uuid string) (models.Device, error)
	FindAll() ([]models.Device, error)
	DeleteByName(device_name string) error
	DeleteByUUID(device_uuid string) error
}

type DeviceRepositoryService struct {
	db *gorm.DB
}

func NewDeviceRepositoryService(db *gorm.DB) DeviceRepositoryService {
	return DeviceRepositoryService{
		db: db,
	}
}

func (drs DeviceRepositoryService) Create(device_name string, device_uuid string) error {
	d := models.Device{Name: device_name, UUID: device_uuid}
	err := drs.db.Create(&d).Error
	return err
}

func (drs DeviceRepositoryService) FindByName(device_name string) (models.Device, error) {
	d := models.Device{Name: device_name}
	err := drs.db.First(&d).Error
	return d, err
}

func (drs DeviceRepositoryService) FindByUUID(device_uuid string) (models.Device, error) {
	d := models.Device{UUID: device_uuid}
	err := drs.db.First(&d).Error
	return d, err
}

func (drs DeviceRepositoryService) FindAll() ([]models.Device, error) {
	var d []models.Device
	err := drs.db.Find(&d).Error
	return d, err
}

func (drs DeviceRepositoryService) DeleteByName(device_name string) error {
	d := models.Device{Name: device_name}
	err := drs.db.First(&d).Error
	if err != nil {
		return err
	}
	err = drs.db.First(&d).Error
	return err
}

func (drs DeviceRepositoryService) DeleteByUUID(device_uuid string) error {
	d := models.Device{UUID: device_uuid}
	err := drs.db.First(&d).Error
	if err != nil {
		return err
	}
	err = drs.db.First(&d).Error
	return err
}
