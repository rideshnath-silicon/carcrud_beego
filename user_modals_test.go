package main_test

import (
	"CarCrudDemo/models"
	"log"
	"testing"
)

func TestUserModels(t *testing.T) {
	t.Run("Get All User", func(t *testing.T) {
		data, err := models.GetAllUser()
		if err != nil {
			t.Errorf("Error ;- %s", err.Error())
			return
		}
		log.Print(data)
	})
	t.Run("Get UserByEmail", func(t *testing.T) {
		data, err := models.GetUserByEmail("rideshnath.siliconithub@gmail.com")
		if err != nil {
			t.Errorf("Error ;- %s", err.Error())
			return
		}
		log.Print(data)
	})
	t.Run("Register new user", func(t *testing.T) {
		var user = models.NewUserRequest{
			FirstName:   "Devendra",
			LastName:    "pohekar",
			Email:       "devendrapohekar.siliconithub@gmail.com",
			PhoneNumber: "9109396802",
			Role:        "developer",
			Country:     "India",
			Age:         24,
			Password:    "123456",
		}
		data, err := models.InsertNewUser(user)
		if err != nil {
			t.Errorf("Error ;- %s", err.Error())
			return
		}
		log.Print(data)
	})
	t.Run("Update user", func(t *testing.T) {
		var user = models.UpdateUserRequest{
			Id:          2,
			FirstName:   "Devendra",
			LastName:    "pohekar",
			Email:       "devendrapohekar.siliconithub@gmail.com",
			PhoneNumber: "9109396802",
			Role:        "developer",
			Country:     "India",
			Age:         23,
		}
		data, err := models.UpdateUser(user)
		if err != nil {
			t.Errorf("Error ;- %s", err.Error())
			return
		}
		log.Print(data)
	})
}
