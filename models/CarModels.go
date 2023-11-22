package models

import "github.com/astaxie/beego/orm"

func GetAllCars() (interface{}, error) {
	o := orm.NewOrm()
	var cars Car
	_, err := o.QueryTable(new(Car)).All(&cars)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func GetSingleCar(id int) (interface{}, error) {
	o := orm.NewOrm()
	var car Car
	_, err := o.QueryTable(new(Car)).Filter("id", id).All(&car)
	if err != nil {
		return nil, err
	}
	return car, nil
}
