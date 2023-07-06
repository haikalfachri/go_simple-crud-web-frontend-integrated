package repositories

import "biodata/models"

type BiodataRepository interface {
	GetAll() ([]models.Biodata, error)
	GetById(id string) (models.Biodata, error)
	Create(biodataInput models.Request) (models.Biodata, error)
	Update(biodataInput models.Request, id string) (models.Biodata, error)
	Delete(id string) error
}