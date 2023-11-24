package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

func InsertNewHomeSetting(data InserNewHomeSettingRequest) (interface{}, error) {
	o := orm.NewOrm()
	var hsetting = HomeSetting{
		Section:   data.Section,
		Type:      data.Type,
		Key:       data.Key,
		Value:     data.Value,
		CreatedAt: time.Now(),
	}
	num, err := o.Insert(&hsetting)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return nil, errors.New("error : error in insert the data")
	}
	return hsetting, nil
}

func UpdateHomeSeting(data UpdateHomeSetingRequest) (interface{}, error) {
	o := orm.NewOrm()
	var hsetting = HomeSetting{
		Id:       data.Id,
		Section:  data.Section,
		Type:     data.Type,
		Key:      data.Key,
		Value:    data.Value,
		UpdateAt: time.Now(),
	}
	num, err := o.Update(&hsetting, "section", "type", "key", "value", "update_at")
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return nil, errors.New("error : error in insert the data")
	}
	return hsetting, nil
}
