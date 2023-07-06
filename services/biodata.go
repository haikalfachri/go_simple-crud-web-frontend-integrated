package services

import (
	"biodata/repositories"
	"biodata/models"
)

type BiodataService struct {
	repository repositories.BiodataRepository
}

func InitBiodataService() BiodataService {
	return BiodataService{
		repository: &repositories.BiodataRepositoryImp{},
	}
}

func (bs *BiodataService) GetAll() ([]models.Biodata, error){
	return bs.repository.GetAll()
}

func (bs *BiodataService) GetById(id string) (models.Biodata, error){
	return bs.repository.GetById(id)
}

func (bs *BiodataService) Create(biodataInput models.Request) (models.Biodata, error){
	return bs.repository.Create(biodataInput)
}

func (bs *BiodataService) Update(biodataInput models.Request, id string) (models.Biodata, error){
	return bs.repository.Update(biodataInput, id)
}

func (bs *BiodataService) Delete(id string) (error){
	return bs.repository.Delete(id)
}
