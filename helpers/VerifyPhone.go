package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

// TwilioCredentials stores Twilio account SID and Auth Token
type TwilioCredentials struct {
	AccountSID string
	AuthToken  string
}

// OutgoingCallerID represents the data needed to add a new caller ID
type OutgoingCallerID struct {
	PhoneNumber string `json:"phone_number"`
}

// VerificationInput represents the data needed to initiate verification
type VerificationInput struct {
	To      string `json:"to"`
	Channel string `json:"channel"`
}

// VerificationCheckInput represents the data needed to check the verification code
type VerificationCheckInput struct {
	Sid  string `json:"-"`
	Code string `json:"code"`
}

var credentials = TwilioCredentials{
	AccountSID: beego.AppConfig.String("TWILIO_ACCOUNT_SID"),
	AuthToken:  beego.AppConfig.String("TWILIO_AUTHTOKEN"),
}

func main() {
	// Your Twilio Account SID and Auth Token
	credentials := TwilioCredentials{
		AccountSID: beego.AppConfig.String("TWILIO_ACCOUNT_SID"),
		AuthToken:  beego.AppConfig.String("TWILIO_AUTHTOKEN"),
	}

	// Your Twilio Phone Number (Twilio phone number to send verification code)
	// twilioNumber := "your_twilio_phone_number"

	// Your desired caller ID
	callerID := "+1234567890" // Replace with your desired caller ID

	// Step 1: Add Caller ID to Twilio
	addCallerID(credentials, callerID)

	// Step 2: Get Verification Code
	verificationSID := initiateVerification(credentials, callerID)

	fmt.Println("Verification SID:", verificationSID)

	// Step 3: Verify Caller ID
	// verificationCode := getVerificationCode()
	verifyCallerID(credentials, verificationSID, "dsssd")
}

func addCallerID(credentials TwilioCredentials, callerID string) {
	url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/OutgoingCallerIds.json", credentials.AccountSID)

	outgoingCallerID := OutgoingCallerID{
		PhoneNumber: callerID,
	}

	payload, err := json.Marshal(outgoingCallerID)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.SetBasicAuth(credentials.AccountSID, credentials.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Add Caller ID Response:", string(body))
}

func initiateVerification(credentials TwilioCredentials, to string) string {
	url := fmt.Sprintf("https://verify.twilio.com/v2/Services/%s/Verifications.json",beego.AppConfig.String("TWILIO_ACCOUNT_SID"))

	verificationInput := VerificationInput{
		To:      to,
		Channel: "sms",
	}

	payload, err := json.Marshal(verificationInput)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ""
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.SetBasicAuth(credentials.AccountSID, credentials.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Initiate Verification Response:", string(body))

	// Extract Verification SID from the response
	var response map[string]interface{}
	json.Unmarshal(body, &response)
	verificationSID, _ := response["sid"].(string)
	return verificationSID
}

func verifyCallerID(credentials TwilioCredentials, verificationSID, code string) {
	url := fmt.Sprintf("https://verify.twilio.com/v2/Services/your_verify_service_sid/VerificationChecks.json")

	verificationCheckInput := VerificationCheckInput{
		Sid:  verificationSID,
		Code: code,
	}

	payload, err := json.Marshal(verificationCheckInput)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.SetBasicAuth(credentials.AccountSID, credentials.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Verification Check Response:", string(body))
}

func VerifyPhone(callerID string) string {
	url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/OutgoingCallerIds.json", credentials.AccountSID)

	outgoingCallerID := OutgoingCallerID{
		PhoneNumber: "+91 " + callerID,
	}

	payload, err := json.Marshal(outgoingCallerID)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ""
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.SetBasicAuth(credentials.AccountSID, credentials.AuthToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Add Caller ID Response:", string(body))
	verificationSID := initiateVerification(credentials, callerID)
	return verificationSID
}
