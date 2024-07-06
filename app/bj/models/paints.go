package models

import (
	"go-admin/common/models"
)

type Paints struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:Name"` 
    Price float64 `json:"price" gorm:"type:decimal(10,2);comment:Price"` 
    Image string `json:"image" gorm:"type:blob;comment:Image"` 
    models.ModelTime
    models.ControlBy
}

func (Paints) TableName() string {
    return "paints"
}

func (e *Paints) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Paints) GetId() interface{} {
	return e.Id
}