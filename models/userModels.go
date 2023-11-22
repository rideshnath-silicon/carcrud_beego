package models

import (
	"CarCrudDemo/helpers"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=mydb sslmode=disable")
	orm.RegisterModel(new(Users), new(Car))
	orm.RunSyncdb("default", false, true)
}

func GetUserByEmail(email string) (Users, error) {
	o := orm.NewOrm()
	var user Users
	_, err := o.QueryTable(new(Users)).Filter("email", email).All(&user)
	if err != nil {
		return user, err
	}
	return user, nil
} 

func LoginUser(email string, pass string) (Users, error) {
	o := orm.NewOrm()
	var user Users
	_, err := o.QueryTable(new(Users)).Filter("email", email).Filter("password", pass).All(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserDetails(id interface{}) (Users, error) {
	o := orm.NewOrm()
	// orm.Debug = true
	var user Users
	_, err := o.QueryTable(new(Users)).Filter("id", id).All(&user, "first_name", "last_name", "email", "password", "phone_number")
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetAllUser() ([]Users, error) {
	o := orm.NewOrm()
	// orm.Debug = true
	var user []Users
	_, err := o.QueryTable(new(Users)).All(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func InsertNewUser(Data NewUserRequest) (interface{}, error) {
	o := orm.NewOrm()
	pass, err := helpers.HashData(Data.Password)
	if err != nil {
		return nil, err
	}
	var user = Users{
		FirstName: Data.FirstName,
		LastName:  Data.LastName,
		Country:   Data.Country,
		Email:     Data.Email,
		Age:       Data.Age,
		Password:  pass,
		Role:      Data.Role,
		CreatedAt: time.Now(),
	}
	data, err := o.Insert(&user)
	if err != nil {
		return data, err
	}
	return user, nil
}

func UpdateUser(Data UpdateUserRequest) (interface{}, error) {
	var user = Users{
		Id:        Data.Id,
		FirstName: Data.FirstName,
		LastName:  Data.LastName,
		Country:   Data.Country,
		Email:     Data.Email,
		Age:       Data.Age,
		Role:      Data.Role,
		UpdatedAt: time.Now(),
	}

	o := orm.NewOrm()
	num, err := o.Update(&user, "id", "first_name", "last_name", "country", "email", "age", "role", "updated_at")
	if err != nil {
		return num, err
	}
	return user, nil
}

func ResetPassword(Password string, id float64) (interface{}, error) {
	pass, err := helpers.HashData(Password)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var user = Users{Id: uint(id), Password: pass}
	num, err := o.Update(&user, "password")
	if err != nil {
		return num, err
	}
	return user, nil
}
