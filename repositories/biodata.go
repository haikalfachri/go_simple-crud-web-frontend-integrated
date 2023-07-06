package repositories

import (
	"biodata/database"
	"biodata/models"
	"os"
)

type BiodataRepositoryImp struct {
}

func InitBiodataRepository() BiodataRepository {
	return &BiodataRepositoryImp{}
}

func (br *BiodataRepositoryImp) GetAll() ([]models.Biodata, error) {
	var biodatas []models.Biodata

	if err := database.DB.Find(&biodatas).Error; err != nil {
		return biodatas, err
	}
	return biodatas, nil
}

func (br *BiodataRepositoryImp) GetById(id string) (models.Biodata, error) {
	var biodata models.Biodata

	if err := database.DB.First(&biodata, id).Error; err != nil {
		return models.Biodata{}, err
	}
	return biodata, nil
}

func (br *BiodataRepositoryImp) Create(request models.Request) (models.Biodata, error) {

	var biodata models.Biodata = models.Biodata{
		Name:    request.Name,
		Phone:   request.Phone,
		Gender:  request.Gender,
		Address: request.Address,
		DOB:     request.DOB,
		URL:     request.URL,
	}

	if err := database.DB.Create(&biodata).Error; err != nil {
		return models.Biodata{}, err
	}

	if err := database.DB.Last(&biodata).Error; err != nil {
		return models.Biodata{}, err
	}
	return biodata, nil
}

func (br *BiodataRepositoryImp) Update(request models.Request, id string) (models.Biodata, error) {

	biodata, err := br.GetById(id)

	if err != nil {
		return models.Biodata{}, err
	}

	if biodata.Name != request.Name {
		biodata.Name = request.Name
	}

	if biodata.DOB != request.DOB {
		biodata.DOB = request.DOB
	}

	if biodata.Address != request.Address {
		biodata.Address = request.Address
	}

	if biodata.Gender != request.Gender {
		biodata.Gender = request.Gender
	}

	if biodata.Phone != request.Phone {
		biodata.Phone = request.Phone
	}

	if biodata.URL != request.URL {
		imgUrl := "./public/assets/picture" + biodata.URL[7:]

		if err := database.DB.Unscoped().Delete(&biodata).Error; err != nil {
			return models.Biodata{}, err
		}

		err = os.Remove(imgUrl)

		if err != nil {
			return models.Biodata{}, err
		}
		
		biodata.URL = request.URL
	}

	if err := database.DB.Save(&biodata).Error; err != nil {
		return models.Biodata{}, err
	}
	return biodata, nil
}

func (br *BiodataRepositoryImp) Delete(id string) error {
	biodata, err := br.GetById(id)

	if err != nil {
		return err
	}

	imgUrl := "./public/assets/picture" + biodata.URL[7:]

	if err := database.DB.Unscoped().Delete(&biodata).Error; err != nil {
		return err
	}

	err = os.Remove(imgUrl)

	if err != nil {
		return err
	}

	return nil
}
