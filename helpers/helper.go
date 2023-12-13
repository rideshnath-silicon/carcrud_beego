package helpers

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/smtp"
	"os"
	"strconv"
	"strings"
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

func GenereateKeyForHomeSection(str1, str2 string) string {
	combinedString := str1 + " " + str2
	underscoredString := strings.ReplaceAll(combinedString, " ", "_")
	// Convert to uppercase
	uppercaseCode := strings.ToUpper(underscoredString)
	return uppercaseCode
}

func SendMailOTp(userEmail string, name string) (string, error) {
	from := beego.AppConfig.String("EMAIL")
	password := beego.AppConfig.String("PASSWORD")
	to := []string{
		userEmail,
	}
	OTP := GenerateOtp()
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	subject := "Verify your email"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := `<div style="font-family: Helvetica,Arial,sans-serif;min-width:1000px;overflow:auto;line-height:2">
				<div style="margin:50px auto;width:70%;padding:20px 0">
				<div style="border-bottom:1px solid #eee">
						<a href="" style="font-size:1.4em;color: #00466a;text-decoration:none;font-weight:600">Hello, I am Ridesh</a>
					</div>
					<p style="font-size:1.1em">Hi, ` + name + `</p>
					<p>Thank you for Register in this app . Use the following OTP to verify your email. OTP is valid for 5 minutes</p>
					<h2 style="background: #00466a;margin: 0 auto;width: max-content;padding: 0 10px;color: #fff;border-radius: 4px;">` + OTP + `</h2>
					<p style="font-size:0.9em;">Regards,<br />Er. Ridesh Nath</p>
					<hr style="border:none;border-top:1px solid #eee" />
					<div style="float:right;padding:8px 0;color:#aaa;font-size:0.8em;line-height:1;font-weight:300">
						<p>Ridesh Nath</p>
						<p>Burhanpur M.P</p>
						<p>India</p>
					</div>
				</div>
			</div>`
	message := []byte("Subject: " + subject + "\r\n" + mime + "\r\n" + body)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return "", err
	}
	return OTP, nil
}

func GenerateOtp() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, 4)
	n, err := io.ReadAtLeast(rand.Reader, b, 4)
	if n != 4 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
