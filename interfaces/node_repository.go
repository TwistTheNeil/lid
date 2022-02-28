package interfaces

import "lid/models"

type NodeRepository interface {
	Create(node_name string, node_hash string, node_size int64) error
	FindByName(node_name string) (models.Node, error)
	FindByMD5(node_md5_hash string) (models.Node, error)
	FindAll() ([]models.Node, error)
	DeleteByName(node_name string) error
	DeleteByMD5(node_md5_hash string) error
}
