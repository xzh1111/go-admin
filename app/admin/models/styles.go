package models

import (
    "go-admin/common/models"
)

type Styles struct {
    models.Model

    Name  string `json:"name" gorm:"type:varchar(255);comment:款式"`
    Price float64 `json:"price" gorm:"type:decimal(10,2);comment:价格"` // 修复了这里的标签值
    Image string `json:"image" gorm:"type:blob;comment:图片"`
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
