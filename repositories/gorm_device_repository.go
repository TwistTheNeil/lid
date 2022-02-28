package repositories

import (
	"lid/models"

	"gorm.io/gorm"
)

type GORMDeviceRepositoryService struct {
	db *gorm.DB
}

var deviceRepositoryService GORMDeviceRepositoryService

func NewGORMDeviceRepositoryService(db *gorm.DB) GORMDeviceRepositoryService {
	if (deviceRepositoryService == GORMDeviceRepositoryService{}) {
		deviceRepositoryService = GORMDeviceRepositoryService{
			db: db,
		}
	}

	return deviceRepositoryService
}

func (drs GORMDeviceRepositoryService) Create(device_name string, device_uuid string, size int64) error {
	d := models.Device{Name: device_name, UUID: device_uuid, Size: size}
	err := drs.db.Create(&d).Error
	return err
}

func (drs GORMDeviceRepositoryService) FindByName(device_name string) (models.Device, error) {
	d := models.Device{Name: device_name}
	err := drs.db.First(&d).Error
	return d, err
}

func (drs GORMDeviceRepositoryService) FindByUUID(device_uuid string) (models.Device, error) {
	d := models.Device{UUID: device_uuid}
	err := drs.db.First(&d).Error
	return d, err
}

func (drs GORMDeviceRepositoryService) FindAll() ([]models.Device, error) {
	var d []models.Device
	err := drs.db.Find(&d).Error
	return d, err
}

func (drs GORMDeviceRepositoryService) DeleteByName(device_name string) error {
	d := models.Device{Name: device_name}
	err := drs.db.First(&d).Error
	if err != nil {
		return err
	}
	err = drs.db.First(&d).Error
	return err
}

func (drs GORMDeviceRepositoryService) DeleteByUUID(device_uuid string) error {
	d := models.Device{UUID: device_uuid}
	err := drs.db.First(&d).Error
	if err != nil {
		return err
	}
	err = drs.db.First(&d).Error
	return err
}
