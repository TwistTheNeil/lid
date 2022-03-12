package repositories

import (
	"lid/models"

	"gorm.io/gorm"
)

type GORMNodeRepositoryService struct {
	db *gorm.DB
}

var nodeRepositoryService GORMNodeRepositoryService

func NewGORMNodeRepositoryService(db *gorm.DB) GORMNodeRepositoryService {
	if (nodeRepositoryService == GORMNodeRepositoryService{}) {
		nodeRepositoryService = GORMNodeRepositoryService{
			db: db,
		}
	}

	return nodeRepositoryService
}

func (nrs GORMNodeRepositoryService) Create(name string, hash string, size int64, storageDevice models.Device) error {
	n := models.Node{Name: name, MD5: hash, Size: size}
	err := nrs.db.Create(&n).Error
	err = nrs.db.Model(&n).Association("Devices").Append(&storageDevice)
	return err
}

func (nrs GORMNodeRepositoryService) FindByName(name string) (models.Node, error) {
	n := models.Node{Name: name}
	err := nrs.db.First(&n).Error
	return n, err
}

func (nrs GORMNodeRepositoryService) FindByMD5(hash string) (models.Node, error) {
	n := models.Node{MD5: hash}
	err := nrs.db.First(&n).Error
	return n, err
}

func (nrs GORMNodeRepositoryService) FindAll() ([]models.Node, error) {
	var n []models.Node
	err := nrs.db.Find(&n).Error
	return n, err
}

func (nrs GORMNodeRepositoryService) FindAllPreload() ([]models.Node, error) {
	var n []models.Node
	err := nrs.db.Preload("Devices").Find(&n).Error
	return n, err
}

func (nrs GORMNodeRepositoryService) DeleteByName(name string) error {
	n := models.Node{Name: name}
	err := nrs.db.First(&n).Error
	if err != nil {
		return err
	}
	err = nrs.db.Delete(&n).Error
	return err
}

func (nrs GORMNodeRepositoryService) DeleteByMD5(hash string) error {
	n := models.Node{MD5: hash}
	err := nrs.db.First(&n).Error
	if err != nil {
		return err
	}
	err = nrs.db.Delete(&n).Error
	return err
}
