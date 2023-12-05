package controllers

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"
	"encoding/json"
	"errors"
	"os"

	"github.com/astaxie/beego"
)

type CarController struct {
	beego.Controller
}

// GetAllCars ...
// @Title get cars
// @Desciption Get all car
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /cars [get]
func (c *CarController) GetAllCars() {
	Data, err := models.GetAllCars()
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
	}
	helpers.ApiSuccess(c.Ctx, Data, 1000)
}

// GetSingleCar ...
// @Title get car
// @Desciption Get all car
// @Param body body models.GetcarRequest true "get perticuler car"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router / [post]
func (c *CarController) GetSingleCar() {
	var bodyData models.GetcarRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
	}
	Data, err := models.GetSingleCar(bodyData.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
	}
	helpers.ApiSuccess(c.Ctx, Data, 1000)
}

// GetCarUsingSearch ...
// @Title search car
// @Desciption search car
// @Param body body models.SearchRequest true "search car"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /search [post]
func (c *CarController) GetCarUsingSearch() {
	var bodyData models.SearchRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	cars, err := models.GetCarUsingSearch(bodyData.Search)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	var output []models.CarDetailsRequest
	for i := 0; i < len(cars); i++ {
		carDetails := models.CarDetailsRequest{CarName: cars[i].CarName, CarImage: cars[i].CarImage, ModifiedBy: cars[i].ModifiedBy, Model: cars[i].Model, Type: cars[i].Type}
		output = append(output, carDetails)
	}
	helpers.ApiSuccess(c.Ctx, output, 1000)
}

// AddNewCar ...
// @Title new car
// @Desciption insert car
// @swagger:parameters upload
// @Param car_name formData string true "Car name"
// @Param modified_by formData string true "modified by"
// @Param model formData string true "Car Model"
// @Param type formData string true "accepted type 'sedan','SUV','hatchback'"
// @Param file formData file true "File to be uploaded"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /create [post]
func (c *CarController) AddNewCar() {
	var cars models.GetNewCarRequest
	if err := c.ParseForm(&cars); err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &cars)
	_, fileheader, err := c.GetFile("file")
	if err != nil {
		helpers.ApiFailure(c.Ctx, "File Getting Error", 1001)
		return
	}
	var carType string = string(cars.Type)
	cars.Type, err = NewCarType(carType)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	filedName := "file"
	uploadDir := "./uploads/car/images/"
	filepaths, err := helpers.UploadFile(c.Controller, filedName, fileheader, uploadDir)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	cars.CarImage = filepaths
	data, err := models.InsertNewCar(cars)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, data, 1002)
}

func NewCarType(input string) (models.CarType, error) {
	switch input {
	case "sedan", "hatchback", "SUV":
		return models.CarType(input), nil
	default:
		return "", errors.New("invalid car type")
	}
}

// UpdateCar ...
// @Title update car
// @Desciption update car
// @Param car_id formData string true "Car name"
// @Param car_name formData string false "Car name"
// @Param modified_by formData string false "modified by"
// @Param model formData string false "Car Model"
// @Param type formData string false "accepted type 'sedan','SUV','hatchback'"
// @Param file formData file false "File to be uploaded"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /update [PUT]
func (c *CarController) UpdateCar() {
	var cars models.UpdateCarRequest
	if err := c.ParseForm(&cars); err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &cars)
	data, err := models.GetSingleCar(cars.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	_, fileheader, err := c.GetFile("file")
	if err != nil {
		if cars.CarName == "" {
			cars.CarName = data.CarName
		}
		if cars.ModifiedBy == "" {
			cars.ModifiedBy = data.ModifiedBy
		}
		if cars.Model == "" {
			cars.Model = data.Model
		}
		if cars.Type == "" {
			cars.Type = data.Type
		}
		var carType string = string(cars.Type)
		cars.Type, err = NewCarType(carType)
		if err != nil {
			helpers.ApiFailure(c.Ctx, err.Error(), 1001)
			return
		}
		cars.CarImage = data.CarImage
		res, err := models.UpdateCar(cars)
		if err != nil {
			helpers.ApiFailure(c.Ctx, err.Error(), 1001)
			return
		}
		helpers.ApiSuccess(c.Ctx, res, 1003)
		return
	}
	var carType string = string(cars.Type)
	cars.Type, err = NewCarType(carType)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	filedName := "file"
	uploadDir := "./uploads/images/"
	filepaths, err := helpers.UploadFile(c.Controller, filedName, fileheader, uploadDir)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	cars.CarImage = filepaths
	output, err := models.UpdateCar(cars)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	err = os.Remove(data.CarImage)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, output, 1003)
}

// DeleteCar ...
// @Title remove car
// @Desciption delete car
// @Param body body models.GetcarRequest true "delete car"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /delete [delete]
func (c *CarController) DeleteCar() {
	var car models.GetcarRequest
	err := helpers.RequestBody(c.Ctx, &car)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	res, err := models.GetSingleCar(car.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	data, err := models.DeleteCar(car.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	err = os.Remove(res.CarImage)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, data, 1004)
}
