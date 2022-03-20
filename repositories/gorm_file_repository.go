package repositories

import (
	"lid/models"

	"gorm.io/gorm"
)

type GORMNodeRepositoryService struct {
	db *gorm.DB
}

var fileRepositoryService GORMNodeRepositoryService

func NewGORMNodeRepositoryService(db *gorm.DB) GORMNodeRepositoryService {
	if (fileRepositoryService == GORMNodeRepositoryService{}) {
		fileRepositoryService = GORMNodeRepositoryService{
			db: db,
		}
	}

	return fileRepositoryService
}

func (frs GORMNodeRepositoryService) Create(name string, hash string, size int64, storageDevice models.Device) error {
	n := models.File{Name: name, MD5: hash, Size: size}
	err := frs.db.Create(&n).Error
	err = frs.db.Model(&n).Association("Devices").Append(&storageDevice)
	return err
}

func (frs GORMNodeRepositoryService) FindByName(name string) (models.File, error) {
	n := models.File{Name: name}
	err := frs.db.First(&n).Error
	return n, err
}

func (frs GORMNodeRepositoryService) FindByMD5(hash string) (models.File, error) {
	n := models.File{MD5: hash}
	err := frs.db.First(&n).Error
	return n, err
}

func (frs GORMNodeRepositoryService) FindAll() ([]models.File, error) {
	var n []models.File
	err := frs.db.Find(&n).Error
	return n, err
}

func (frs GORMNodeRepositoryService) FindAllPreload() ([]models.File, error) {
	var n []models.File
	err := frs.db.Preload("Devices").Find(&n).Error
	return n, err
}

func (frs GORMNodeRepositoryService) DeleteByName(name string) error {
	n := models.File{Name: name}
	err := frs.db.First(&n).Error
	if err != nil {
		return err
	}
	err = frs.db.Delete(&n).Error
	return err
}

func (frs GORMNodeRepositoryService) DeleteByMD5(hash string) error {
	n := models.File{MD5: hash}
	err := frs.db.First(&n).Error
	if err != nil {
		return err
	}
	err = frs.db.Delete(&n).Error
	return err
}
