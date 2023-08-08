package interfaces

import "lid/models"

type HostRepository interface {
	Create(host_name string) error
	FindByName(host_name string) (models.Host, error)
	FindAll() ([]models.Host, error)
	FindAllPreload() ([]models.Host, error)
	DeleteByName(host_name string) error
}
