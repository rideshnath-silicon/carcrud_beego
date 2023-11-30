package controllers

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) GetAllUser() {
	user, err := models.GetAllUser()
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}

	var output []models.UserDetailsRequest
	for i := 0; i < len(user); i++ {
		userDetails := models.UserDetailsRequest{Id: user[i].Id, FirstName: user[i].FirstName, LastName: user[i].LastName, Email: user[i].Email, Country: user[i].Country, Age: user[i].Age}
		output = append(output, userDetails)
	}
	p := pagination.NewPaginator(c.Ctx.Request, 3, len(output))

	// Get the current page's items
	startIndex := p.Offset()

	// Calculate the end index
	endIndex := startIndex + 3
	if endIndex > len(output) {
		endIndex = len(output)
	}
	pageItems := output[startIndex:endIndex]
	helpers.ApiSuccess(c.Ctx, pageItems, 1000)
}

func (c *UserController) RegisterNewUser() {
	var bodyData models.NewUserRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	data, _ := models.GetUserByEmail(bodyData.Email)

	if data.Email == bodyData.Email {
		helpers.ApiFailure(c.Ctx, "Email already used by another account please try with new email", 10001)
		return
	}
	output, err := models.InsertNewUser(bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, output, 1002)
}

func (c *UserController) UpdateUser() {
	var bodyData models.UpdateUserRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}

	data, err := models.GetUserDetails(bodyData.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	if bodyData.Email != data.Email {
		res, _ := models.GetUserByEmail(bodyData.Email)
		if res.Email == bodyData.Email {
			helpers.ApiFailure(c.Ctx, "Email already used by another account please try with new email", 10001)
			return
		}
	}
	output, err := models.UpdateUser(bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, output, 1003)
}

func (c *UserController) ResetPassword() {
	claims := helpers.GetUserDataFromTokan(c.Ctx)
	id := claims["User_id"].(float64)
	output, err := models.GetUserDetails(id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	var bodyData models.ResetUserPassword
	err = helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	err = helpers.VerifyHashedData(output.Password, bodyData.CurrentPass)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	if bodyData.ConfirmPass != bodyData.NewPass {
		helpers.ApiFailure(c.Ctx, "Please match new password and confirm password", 1001)
		return
	}
	uppass, err := models.ResetPassword(bodyData.NewPass, id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, uppass, 1003)
}

func (c *UserController) SendOtp() {
	var bodyData models.SendOtpData
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	output, err := models.GetUserByEmail(bodyData.Username)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	_, err = helpers.TwilioSendOTP(output.PhoneNumber)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	otp, err := helpers.SendMailOTp(output.Email, output.FirstName)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	Response, err := models.UpadteOtpForEmail(output.Id, otp)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, Response, 1000)
	go func() {
		newOtp := helpers.GenerateOtp()
		models.UpdateColumnOTP(output.Id, newOtp)
	}()
}

func (c *UserController) VerifyOtpResetpassword() {
	var bodyData models.ResetUserPasswordOtp
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	output, err := models.GetUserByEmail(bodyData.Email)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	err = helpers.TwilioVerifyOTP(output.PhoneNumber, bodyData.Otp)
	if err != nil {
		data, err := models.VerifyEmailOTP(bodyData.Email, bodyData.Otp)
		if err != nil {
			helpers.ApiFailure(c.Ctx, err.Error(), 1001)
			return
		}
		if data.Otp != bodyData.Otp {
			helpers.ApiFailure(c.Ctx, "Please Eenter Valid otp", 5001)
		}
		err = models.UpdateVerified(data.Id)
		if err != nil {
			helpers.ApiFailure(c.Ctx, err.Error(), 1001)
			return
		}
		uppass, err := models.ResetPassword(bodyData.NewPass, float64(output.Id))
		if err != nil {
			helpers.ApiFailure(c.Ctx, err.Error(), 1001)
			return
		}
		helpers.ApiSuccess(c.Ctx, uppass, 1003)
	}
	uppass, err := models.ResetPassword(bodyData.NewPass, float64(output.Id))
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, uppass, 1003)
}

func (c *UserController) VerifyUserEmail() {
	var bodyData models.SendOtpData
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	output, err := models.GetUserByEmail(bodyData.Username)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	otp, err := helpers.SendMailOTp(output.Email, output.FirstName)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	res, err := models.UpadteOtpForEmail(output.Id, otp)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, res, 1000)
	go func() {
		newOtp := helpers.GenerateOtp()
		models.UpdateColumnOTP(output.Id, newOtp)
	}()
}

func (c *UserController) VerifyEmailOTP() {
	var bodyData models.VerifyEmailOTPRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	data, err := models.VerifyEmailOTP(bodyData.Username, bodyData.Otp)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	if data.Otp != bodyData.Otp {
		helpers.ApiFailure(c.Ctx, "Please Eenter Valid otp", 5001)
	}
	err = models.UpdateVerified(data.Id)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, "Your Account is Successfully Verified", 5000)
}

func (c *UserController) GetCountryWiseCountUser() {
	res, err := models.GetCountryWiseCountUser()
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	helpers.ApiSuccess(c.Ctx, res, 1000)
}

func (c *UserController) GetVerifiedUsers() {
	user, err := models.GetVerifiedUsers()
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}

	var output []models.UserDetailsRequest
	for i := 0; i < len(user); i++ {
		userDetails := models.UserDetailsRequest{Id: user[i].Id, FirstName: user[i].FirstName, LastName: user[i].LastName, Email: user[i].Email, Country: user[i].Country, Age: user[i].Age}
		output = append(output, userDetails)
	}
	helpers.ApiSuccess(c.Ctx, output, 1000)
}

func (c *UserController) SearchUser() {
	var bodyData models.SearchRequest
	err := helpers.RequestBody(c.Ctx, &bodyData)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	user, err := models.SearchUser(bodyData.Search)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	var output []models.UserDetailsRequest
	for i := 0; i < len(user); i++ {
		userDetails := models.UserDetailsRequest{Id: user[i].Id, LastName: user[i].LastName, Email: user[i].Email, FirstName: user[i].FirstName, Country: user[i].Country, Age: user[i].Age}
		output = append(output, userDetails)
	}
	helpers.ApiSuccess(c.Ctx, output, 1000)
}
