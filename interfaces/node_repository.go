package interfaces

import "lid/models"

type NodeRepository interface {
	Create(name string, hash string, size int64) error
	FindByName(name string) (models.Node, error)
	FindByMD5(hash string) (models.Node, error)
	FindAll() ([]models.Node, error)
	DeleteByName(name string) error
	DeleteByMD5(hash string) error
}
