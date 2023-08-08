package repositories

import (
	"lid/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GORMHostRepositoryService struct {
	db *gorm.DB
}

var hostRepositoryService GORMHostRepositoryService

func NewGORMHostRepositoryService(db *gorm.DB) GORMHostRepositoryService {
	if (hostRepositoryService == GORMHostRepositoryService{}) {
		hostRepositoryService = GORMHostRepositoryService{
			db: db,
		}
	}

	return hostRepositoryService
}

func (drs GORMHostRepositoryService) Create(host_name string) error {
	h := models.Host{Name: host_name}
	err := drs.db.Create(&h).Error
	return err
}

func (drs GORMHostRepositoryService) FindByName(host_name string) (models.Host, error) {
	h := models.Host{Name: host_name}
	err := drs.db.First(&h).Error
	return h, err
}

func (drs GORMHostRepositoryService) FindAll() ([]models.Host, error) {
	var h []models.Host
	err := drs.db.Find(&h).Error
	return h, err
}

func (drs GORMHostRepositoryService) FindAllPreload() ([]models.Host, error) {
	var h []models.Host
	err := drs.db.Preload("Devices.Files").Preload(clause.Associations).Find(&h).Error
	return h, err
}

func (drs GORMHostRepositoryService) DeleteByName(host_name string) error {
	h := models.Host{Name: host_name}
	err := drs.db.First(&h).Error
	if err != nil {
		return err
	}
	err = drs.db.Delete(&h).Error
	return err
}
