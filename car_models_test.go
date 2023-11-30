package main_test

import (
	"CarCrudDemo/models"
	"log"
	"testing"
)

func TestCarmoduls(t *testing.T) {
	t.Run("Get all cars", func(t *testing.T) {
		data, err := models.GetAllCars()
		if err != nil {
			t.Errorf("Error ;- %s", err.Error())
			return
		}
		log.Print(data)
	})
	t.Run("Register New Car", func(t *testing.T) {
		var car = models.GetNewCarRequest{
			CarName:    "baleno",
			CarImage:   "test_case",
			ModifiedBy: "suzuki",
			Model:      "new Beleno",
			Type:       "sedan",
		}
		data, err := models.InsertNewCar(car)
		if err != nil {
			t.Errorf("Error ;- %s", err.Error())
			return
		}
		log.Print(data)
	})

	t.Run("Update Car", func(t *testing.T) {
		var car = models.UpdateCarRequest{
			Id:         2,
			CarName:    "baleno",
			CarImage:   "test_case",
			ModifiedBy: "suzuki",
			Model:      "Beleno",
			Type:       "sedan",
		}
		data, err := models.UpdateCar(car)
		if err != nil {
			t.Errorf("Error ;- %s", err.Error())
			return
		}
		log.Print(data)
	})
	t.Run("search car", func(t *testing.T) {
		data, err := models.GetCarUsingSearch("th")
		if err != nil {
			t.Errorf("Error ;- %s", err.Error())
			return
		}
		log.Print(data)
	})
}
