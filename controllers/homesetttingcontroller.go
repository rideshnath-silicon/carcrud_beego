package controllers

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type HomeSettingController struct {
	beego.Controller
}

func (c *HomeSettingController) InsertNewHomeSetting() {
	var bodyData models.InserNewHomeSettingRequest
	if err := c.ParseForm(&bodyData); err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &bodyData)
	bodyData.Key = helpers.GenereateKeyForHomeSection(bodyData.Section, bodyData.Type)
	output, err := models.InsertNewHomeSetting(bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, output, 1003)
}

func (c *HomeSettingController) UpdateHomeSeting() {
	var bodyData models.UpdateHomeSetingRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	output, err := models.UpdateHomeSeting(bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, output, 1003)
}
