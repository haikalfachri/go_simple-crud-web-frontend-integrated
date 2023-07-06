package controllers

import (
	"biodata/models"
	"biodata/services"
	"biodata/utils"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

type BiodataController struct {
	service services.BiodataService
}

func InitBiodataContoller() BiodataController {
	return BiodataController{
		service: services.InitBiodataService(),
	}
}

func (bc *BiodataController) GetAll(c echo.Context) error {
	biodatas, err := bc.service.GetAll()
	if err != nil {
		return models.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}
	return models.NewResponse(c, http.StatusOK, "success", "success fetch all biodatas", biodatas)
}

func (bc *BiodataController) GetById(c echo.Context) error {
	id := c.Param("id")
	biodata, err := bc.service.GetById(id)
	if err != nil {
		return models.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}
	return models.NewResponse(c, http.StatusOK, "success", "success fetch a biodata", biodata)
}

func (bc *BiodataController) Create(c echo.Context) error {
	name := c.FormValue("name")
	dob := c.FormValue("dob")
	parseBod, err := time.Parse("2006-01-02T15:04:05.000-07:00", dob)
	address := c.FormValue("address")
	phone := c.FormValue("phone")
	gender := c.FormValue("gender")

	image, err := c.FormFile("image")

	if err != nil {
		return models.NewResponse(c, http.StatusBadRequest, "failed", "failed to upload file", "")
	}

	imgName := utils.GenerateUniqueFileName("user.png")

	var biodataInput models.Request = models.Request{
		Name:    name,
		Phone:   phone,
		Address: address,
		Gender:  gender,
		DOB:     parseBod,
		URL:     "picture/" + imgName,
	}

	biodata, err := bc.service.Create(biodataInput)

	if err != nil {
		return models.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	src, err := image.Open()
	if err != nil {
		return err
	}

	localUrl := "./public/assets/picture/"

	dst, err := os.Create(localUrl + imgName)
	if err != nil {
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	defer src.Close()
	defer dst.Close()

	return models.NewResponse(c, http.StatusOK, "success", "biodata created", biodata)
}

func (bc *BiodataController) Delete(c echo.Context) error {
	id := c.Param("id")
	err := bc.service.Delete(id)
	if err != nil {
		return models.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}
	return models.NewResponse(c, http.StatusOK, "success", "success delete a biodata", "")
}

func (bc *BiodataController) Update(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	dob := c.FormValue("dob")
	parseBod, err := time.Parse("2006-01-02T15:04:05.000-07:00", dob)
	address := c.FormValue("address")
	phone := c.FormValue("phone")
	gender := c.FormValue("gender")

	image, err := c.FormFile("image")

	if err != nil {
		return models.NewResponse(c, http.StatusBadRequest, "failed", "failed to upload file", "")
	}

	imgName := utils.GenerateUniqueFileName("user.png")

	var biodataInput models.Request = models.Request{
		Name:    name,
		Phone:   phone,
		Address: address,
		Gender:  gender,
		DOB:     parseBod,
		URL:     "picture/" + imgName,
	}

	biodata, err := bc.service.Update(biodataInput, id)

	if err != nil {
		return models.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	src, err := image.Open()
	if err != nil {
		return err
	}

	localUrl := "./public/assets/picture/"

	dst, err := os.Create(localUrl + imgName)
	if err != nil {
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	defer src.Close()
	defer dst.Close()

	return models.NewResponse(c, http.StatusOK, "success", "biodata updated", biodata)
}

