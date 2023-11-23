package helpers

import (
	"encoding/json"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashData(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
func VerifyHashedData(hashedString string, dataString string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(dataString))
	if err != nil {
		return err
	}
	return nil
}

func GetUserDataFromTokan(c *context.Context) map[string]interface{} {
	userClaims := c.Input.GetData("user")
	userID := userClaims.(jwt.MapClaims)["ID"]
	userEmail := userClaims.(jwt.MapClaims)["Email"]
	response := map[string]interface{}{"Email": userEmail, "User_id": userID}
	return response
}

func RequestBody(ctx *context.Context, Struct interface{}) error {
	bodyData := ctx.Input.RequestBody
	// fmt.Println(bodyData)
	err := json.Unmarshal(bodyData, Struct)
	if err != nil {
		return err
	}
	return nil
}

func ApiSuccess(c *context.Context, data interface{}, messageCode int) {
	type ApiSuccessResponse struct {
		Message string
		Success int
		Data    interface{}
	}
	message := Messagess(messageCode)
	Response := ApiSuccessResponse{
		Message: message,
		Success: 1,
		Data:    data,
	}
	c.Output.JSON(Response, true, false)
}

func ApiFailure(c *context.Context, data interface{}, messageCode int) {
	type ApiSuccessResponse struct {
		Message string
		Success int
		Data    interface{}
	}
	message := Messagess(messageCode)
	Response := ApiSuccessResponse{
		Message: message,
		Success: 0,
		Data:    data,
	}
	c.Output.JSON(Response, true, false)
}

// otp verification from here
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: beego.AppConfig.String("TWILIO_ACCOUNT_SID"),
	Password: beego.AppConfig.String("TWILIO_AUTHTOKEN"),
})

func TwilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo("+91" + phoneNumber)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(beego.AppConfig.String("TWILIO_SERVICES_ID"), params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

func TwilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(beego.AppConfig.String("TWILIO_SERVICES_ID"), params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return nil
}

func UploadFile(c beego.Controller, filedName string, fileheader *multipart.FileHeader, uploadPath string) (string, error) {
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		// Create the directory and any necessary parent directories
		err := os.MkdirAll("./"+uploadPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	filePath := uploadPath + strconv.FormatInt(time.Now().UnixNano(), 10) + fileheader.Filename
	err := c.SaveToFile(filedName, filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
