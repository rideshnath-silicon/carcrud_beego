package controllers

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"

	"github.com/astaxie/beego"
)

type CarController struct {
	beego.Controller
}

func (c *CarController) GetAllCars() {
	Data, err := models.GetAllCars()
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
	}
	helpers.ApiSuccess(c.Ctx, Data, 1000)
}

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
