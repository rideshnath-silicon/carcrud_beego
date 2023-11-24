package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

func GetAllCars() (interface{}, error) {
	o := orm.NewOrm()
	var cars []Car
	num, err := o.QueryTable(new(Car)).All(&cars)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return nil, errors.New("error :- Data Not Found")
	}
	return cars, nil
}

func GetSingleCar(id uint) (Car, error) {
	o := orm.NewOrm()
	var car Car
	num, err := o.QueryTable(new(Car)).Filter("id", id).All(&car)
	if err != nil {
		return car, err
	}
	if num == 0 {
		return car, errors.New("error :- please enter valid car id")
	}
	return car, nil
}

func GetCarUsingSearch(search string) ([]Car, error) {
	o := orm.NewOrm()
	var car []Car
	num, err := o.QueryTable(new(Car)).SetCond(orm.NewCondition().Or("car_name__icontains", search).Or("carmodel__icontains", search).Or("car_type__icontains", search)).All(&car)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return nil, errors.New("error :- No car found")
	}
	return car, nil
}

func InsertNewCar(data GetNewCarRequest) (interface{}, error) {
	o := orm.NewOrm()
	var car = Car{
		CarName: data.CarName, 
		CarImage: data.CarImage, 
		ModifiedBy: data.ModifiedBy, 
		Model: data.Model, 
		Type: data.Type, 
		CreatedDate: time.Now(),
	}
	_, err := o.Insert(&car)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func UpdateCar(data UpdateCarRequest) (interface{}, error) {
	o := orm.NewOrm()
	var car = Car{
		Id:         data.Id,
		CarName:    data.CarName,
		ModifiedBy: data.ModifiedBy,
		Model:      data.Model,
		Type:       data.Type,
		CarImage:   data.CarImage,
		UpdateAt:   time.Now(),
	}
	fmt.Println(data)

	num, err := o.Update(&car, "id", "car_name", "modified_by", "model", "car_type", "car_image", "update_at")
	if err != nil {
		return num, err
	}
	return car, nil
}

func DeleteCar(id uint) (interface{}, error) {
	o := orm.NewOrm()
	var car = Car{Id: id}
	_, err := o.Delete(&car)
	if err != nil {
		return nil, err
	}
	return car, nil
}

