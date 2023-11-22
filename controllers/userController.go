package controllers

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) GetAllUser() {
	user, err := models.GetAllUser()
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var output []models.UserDetailsRequest
	for i := 0; i < len(user); i++ {
		userDetails := models.UserDetailsRequest{Id: user[i].Id, FirstName: user[i].FirstName, LastName: user[i].LastName, Email: user[i].Email, Country: user[i].Country, Age: user[i].Age}
		output = append(output, userDetails)
	}
	helpers.ApiSuccess(c.Ctx, output, 1000)
}

func (c *UserController) RegisterNewUser() {
	var bodyData models.NewUserRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	output, err := models.InsertNewUser(bodyData)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = output
	c.ServeJSON()
}

func (c *UserController) UpdateUser() {
	var bodyData models.UpdateUserRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	output, err := models.UpdateUser(bodyData)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = output
	c.ServeJSON()
}

func (c *UserController) ResetPassword() {
	claims := helpers.GetUserDataFromTokan(c.Ctx)
	id := claims["User_id"].(float64)
	output, err := models.GetUserDetails(id)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var bodyData models.ResetUserPassword
	err = helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	err = helpers.VerifyHashedData(output.Password, bodyData.CurrentPass)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	if bodyData.ConfirmPass != bodyData.NewPass {
		c.Data["json"] = "Please match new password and confirm password"
		c.ServeJSON()
		return
	}
	uppass, err := models.ResetPassword(bodyData.NewPass, id)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = uppass
	c.ServeJSON()
}

func (c *UserController) SendOtp() {
	var bodyData models.SendOtpData
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	output, err := models.GetUserByEmail(bodyData.Email)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	otpstr, err := helpers.TwilioSendOTP(output.PhoneNumber)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	Response := map[string]interface{}{"OTP": otpstr, "message": "Otp sent on registerd mobile number"}
	helpers.ApiSuccess(c.Ctx, Response, 1000)
}

func (c *UserController) VerifyOtpResetpassword() {
	var bodyData models.ResetUserPasswordOtp
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	output, err := models.GetUserByEmail(bodyData.Email)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	err = helpers.TwilioVerifyOTP(output.PhoneNumber, bodyData.Otp)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	uppass, err := models.ResetPassword(bodyData.NewPass, float64(output.Id))
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = uppass
	c.ServeJSON()
}
