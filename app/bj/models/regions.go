package models

import (

	"go-admin/common/models"

)

type Regions struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:地区"` 
    Base1 float64 `json:"base1" gorm:"type:decimal(10,2);comment:基数1"` 
    Base2 float64 `json:"base2" gorm:"type:decimal(10,2);comment:基数2"` 
    models.ModelTime
    models.ControlBy
}

func (Regions) TableName() string {
    return "regions"
}

func (e *Regions) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Regions) GetId() interface{} {
	return e.Id
}