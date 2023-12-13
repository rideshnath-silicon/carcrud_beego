package controllers

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/beego/beego/httplib"
	// "github.com/astaxie/beego/utils/pagination"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Loginpage() {
	message := c.GetSession("Error")
	c.DelSession("Error")
	c.Data["message"] = message
	c.TplName = "login.tpl"
	if err := c.Render(); err != nil {
		// Handle the error, e.g., log it or show an error page
		fmt.Println("Template rendering error:", err)
	}
}

func (c *UserController) Loginsave() {
	username := c.GetString("userName")
	password := c.GetString("password")
	// c.Ctx.WriteString(fmt.Sprintf("\nValue: %v,sesssion value: %v\n", username, password))
	// return
	req := httplib.Post("http://localhost:8080/v1/user/login")
	// req.Header("Content-Type", "application/json")
	req.Body(`{"username":"` + username + `","password":"` + password + `"}`)

	str, err := req.String()
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(str), &resultMap)
	if err != nil {
		c.SetSession("Error", err.Error())
		return
	}
	if resultMap["Success"] == 0.0 {
		c.SetSession("Error", resultMap["Data"])
		c.Redirect("/v1/my", 302)
		return
	}
	Data := resultMap["Data"]
	dataMap, ok := Data.(map[string]interface{})
	if !ok {
		fmt.Println("Data is not a map[string]interface{}", resultMap["Success"])
		return
	}
	req = httplib.Get("http://localhost:8080/v1/user/secure/users")
	req.Header("Authorization", dataMap["Tokan"].(string))
	strs, err := req.String()
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	c.DelSession("Success")
	// if resultMap["Success"] == 0.0 {
	// 	c.SetSession("Error", resultMap["Data"])
	// 	c.Redirect("/v1/my", 302)
	// 	return
	// }
	// if resultMap["Success"] == 1.0 {
	// 	c.SetSession("Success", resultMap["Message"])
	// 	return
	// }
	c.Ctx.WriteString(strs)
	// c.Ctx.WriteString(fmt.Sprintf("\nValue: %v,", str))
}

// GetAll ...
// @Title Get All
// @Description get Users
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 200 {object} models.Users
// @Failure 403
// @router /secure/users [get]
func (c *UserController) GetAllUser() {
	user, err := models.GetAllUser()
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	// var output []models.UserDetailsRequest
	// for i := 0; i < len(user); i++ {
	// 	userDetails := models.UserDetailsRequest{Id: user[i].Id, FirstName: user[i].FirstName, LastName: user[i].LastName, Email: user[i].Email, Country: user[i].Country, Age: user[i].Age}
	// 	output = append(output, userDetails)
	// }
	// p := pagination.NewPaginator(c.Ctx.Request, 3, len(output))

	// Get the current page's items
	// startIndex := p.Offset()
	// // Calculate the end index
	// endIndex := startIndex + 3
	// if endIndex > len(output) {
	// 	endIndex = len(output)
	// }
	// c.Ctx.WriteString(fmt.Sprintf("\nValue: %v,", user))
	c.Data["users"] = user
	c.TplName = "index.tpl"
	if err := c.Render(); err != nil {
		// Handle the error, e.g., log it or show an error page
		fmt.Println("Template rendering error:", err)
	}
	// pageItems := output[startIndex:endIndex]
	// helpers.ApiSuccess(c.Ctx, user, 1000)
}

// PostRegisterNewUser ...
// @Title Insert New User
// @Desciption new users
// @Param body body models.NewUserRequest true "Insert New User"
// @Success 201 {object} models.Users
// @Failure 403
// @router /secure/register [post]
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

// UpdateUser ...
// @Title update User
// @Desciption update users
// @Param body body models.UpdateUserRequest true "update New User"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} models.Users
// @Failure 403
// @router /secure/update [put]
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

// ResetPassword ...
// @Title Reset password
// @Desciption Reset password
// @Param body body models.ResetUserPassword true "reset password"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} models.Users
// @Failure 403
// @router /secure/reset_pass [post]
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

// SendOtp ...
// @Title forgot password
// @Desciption forgot password
// @Param body body models.SendOtpData true "forgot password this is send otp on mobile and email"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /secure/forgot_pass [post]
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

// VerifyOtpResetpassword ...
// @Title verify otp
// @Desciption otp verification for forgot password
// @Param body body models.ResetUserPasswordOtp true "otp verification for forgot password"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /secure/reset_pass_otp [post]
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

// VerifyUserEmail ...
// @Title verify email
// @Desciption Verify email
// @Param body body models.SendOtpData true "verify email"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /secure/verify_email [post]
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

// VerifyEmailOTP ...
// @Title verify otp for email
// @Desciption otp verification for eamil
// @Param body body models.VerifyEmailOTPRequest true "otp verification for email"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /secure/verify_email_otp [post]
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

// GetVerifiedUsers ...
// @Title verifid users
// @Desciption Get all verified user
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /secure/verified_user [get]
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

// SearchUser ...
// @Title Search User
// @Desciption SearchUser
// @Param body body models.SearchRequest true "otp verification for email"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 201 {object} string
// @Failure 403
// @router /secure/search [post]
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
