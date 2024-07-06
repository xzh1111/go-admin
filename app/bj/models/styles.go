package models

import (

	"go-admin/common/models"

)

type Styles struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:款式"` 
    Price float64 `json:"price" gorm:"type:decimal(10,2);comment:款式单价"` 
    Image string `json:"image" gorm:"type:mediumblob;comment:Image"` 
    ImageExt string `json:"imageExt" gorm:"type:mediumblob;comment:ImageExt"` 
    WeightFormula string `json:"weightFormula" gorm:"type:varchar(800);comment:款式重量公式"` 
    models.ModelTime
    models.ControlBy
}

func (Styles) TableName() string {
    return "styles"
}

func (e *Styles) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Styles) GetId() interface{} {
	return e.Id
}