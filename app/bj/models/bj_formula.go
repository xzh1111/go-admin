package models

import (

	"go-admin/common/models"

)

type BjFormula struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:公式名"` 
    Formula string `json:"formula" gorm:"type:varchar(800);comment:公式"` 
    FormulaKey string `json:"formulaKey" gorm:"type:varchar(255);comment:公式标识"` 
    models.ModelTime
    models.ControlBy
}

func (BjFormula) TableName() string {
    return "bj_formula"
}

func (e *BjFormula) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BjFormula) GetId() interface{} {
	return e.Id
}