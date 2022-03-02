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

func (nrs GORMNodeRepositoryService) Create(name string, hash string, size int64) error {
	d := models.Node{Name: name, MD5: hash, Size: size}
	err := nrs.db.Create(&d).Error
	return err
}
