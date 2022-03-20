package interfaces

import "lid/models"

type FileRepository interface {
	Create(name string, hash string, size int64, storageDevice models.Device) error
	FindByName(name string) (models.File, error)
	FindByMD5(hash string) (models.File, error)
	FindAll() ([]models.File, error)
	FindAllPreload() ([]models.File, error)
	DeleteByName(name string) error
	DeleteByMD5(hash string) error
}
