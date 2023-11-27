package controllers

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"
	"os"

	"github.com/astaxie/beego"
)

type HomeSettingController struct {
	beego.Controller
}

func (c *HomeSettingController) GetHomeSetting() {
	var bodyData models.GetHomeSettingRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	data, err := models.GetHomeSetting(bodyData.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, data, 1000)
}

func (c *HomeSettingController) InsertNewHomeSetting() {
	var bodyData models.InserNewHomeSettingRequest
	if err := c.ParseForm(&bodyData); err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	if bodyData.Type == "Banner url" || bodyData.Type == "Logo url" {
		_, fileheader, err := c.GetFile("file")
		if err != nil {
			helpers.ApiFailure(c.Ctx, "File Getting Error", 1001)
			return
		}
		filedName := "file"
		uploadDir := "./uploads/Homesetings/images/"
		filepaths, err := helpers.UploadFile(c.Controller, filedName, fileheader, uploadDir)
		if err != nil {
			helpers.ApiFailure(c.Ctx, err.Error(), 1001)
			return
		}
		bodyData.Value = filepaths
	}
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
	if err := c.ParseForm(&bodyData); err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	data, err := models.GetHomeSetting(bodyData.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}

	if bodyData.Type == "Banner url" || bodyData.Type == "Logo url" {
		_, fileheader, err := c.GetFile("file")
		if err != nil {
			helpers.ApiFailure(c.Ctx, "File Getting Error", 1001)
			return
		}
		filedName := "file"
		uploadDir := "./uploads/Homesetings/images/"
		filepaths, err := helpers.UploadFile(c.Controller, filedName, fileheader, uploadDir)
		if err != nil {
			helpers.ApiFailure(c.Ctx, err.Error(), 1001)
			return
		}
		bodyData.Value = filepaths
		os.Remove(data.Value)
	}
	bodyData.Key = helpers.GenereateKeyForHomeSection(bodyData.Section, bodyData.Type)
	output, err := models.UpdateHomeSeting(bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, output, 1003)
}

func (c *HomeSettingController) GetUserWiseHome() {
	var bodyData models.GetHomeSettingRequest
	if err := c.ParseForm(&bodyData); err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	res, err := models.UserWiseHomeseting(bodyData.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, res, 1000)
}
