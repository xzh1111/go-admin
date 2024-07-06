package models

import (

	"go-admin/common/models"

)

type Thickness struct {
    models.Model
    
    Thickness string `json:"thickness" gorm:"type:varchar(255);comment:板厚"` 
    TkKg float64 `json:"tkKg" gorm:"type:float;comment:每平方米重量"` 
    models.ModelTime
    models.ControlBy
}

func (Thickness) TableName() string {
    return "thickness"
}

func (e *Thickness) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Thickness) GetId() interface{} {
	return e.Id
}