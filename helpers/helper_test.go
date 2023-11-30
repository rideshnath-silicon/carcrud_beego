package helpers

import (
	"log"
	"testing"
)

func TestHelpers(t *testing.T) {

	t.Run("Check Hashdata", func(t *testing.T) {
		password := "123456"
		data, err := HashData(password)
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Print("sucees to verify hasdata")

		err = VerifyHashedData(data, password)
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Print("sucees to verify veiry hasdata")
	})

	t.Run("Send Otp", func(t *testing.T) {
		_, err := SendMailOTp("rideshnath.siliconithub@gmail.com", "ridesh")
		if err != nil {
			log.Fatalf("error :- %s", err.Error())
		}
		log.Print("sucees to send mail")
	})
}


