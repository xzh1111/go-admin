package models

import (

	"go-admin/common/models"

)

type QuotationsV1 struct {
    models.Model
    
    OrderNumber string `json:"orderNumber" gorm:"type:varchar(255);comment:单号"` 
    Style string `json:"style" gorm:"type:varchar(255);comment:款式"` 
    Color string `json:"color" gorm:"type:varchar(255);comment:颜色"` 
    Length int `json:"length" gorm:"type:int;comment:长"` 
    Width int `json:"width" gorm:"type:int;comment:宽"` 
    Height int `json:"height" gorm:"type:int;comment:高"` 
    Thickness float64 `json:"thickness" gorm:"type:float;comment:板厚"` 
    PipeWeightAndExtraWeight float64 `json:"pipeWeightAndExtraWeight" gorm:"type:float;comment:管重及额外重量"` 
    Quantity int `json:"quantity" gorm:"type:int;comment:数量"` 
    SingleIronWeight float64 `json:"singleIronWeight" gorm:"type:float;comment:单个白铁重量"`
    FinalLogisticsWeight float64 `json:"finalLogisticsWeight" gorm:"type:float;comment:最终物流重量"` 
    IronPrice float64 `json:"ironPrice" gorm:"type:float;comment:白铁价格"` 
    BakingPaintPrice float64 `json:"bakingPaintPrice" gorm:"type:float;comment:烤漆价格"` 
    NeedSilkScreen int `json:"needSilkScreen" gorm:"type:tinyint(1);comment:是否需要丝印"` 
    NumberOfOpenTemplates int `json:"numberOfOpenTemplates" gorm:"type:int;comment:开模板数"` 
    OpenBoardFee float64 `json:"openBoardFee" gorm:"type:float;comment:开板费"` 
    NumberOfSilkScreenFaces int `json:"numberOfSilkScreenFaces" gorm:"type:int;comment:丝印面数"` 
    NumberOfColors int `json:"numberOfColors" gorm:"type:int;comment:印几个颜色"` 
    PrintPrice float64 `json:"printPrice" gorm:"type:float;comment:印的价格"` 
    Freight float64 `json:"freight" gorm:"type:float;comment:运费"` 
    PackagingFee float64 `json:"packagingFee" gorm:"type:float;comment:包装费"` 
    Cost float64 `json:"cost" gorm:"type:float;comment:成本"` 
    Profit float64 `json:"profit" gorm:"type:float;comment:利润"` 
    FinalQuotation float64 `json:"finalQuotation" gorm:"type:float;comment:最终报价"` 
    PowerSource string `json:"powerSource" gorm:"type:varchar(255);comment:插电或者太阳能"` 
    LightStripCost float64 `json:"lightStripCost" gorm:"type:float;comment:灯带费用"` 
    NumberOfIlluminatedFaces int `json:"numberOfIlluminatedFaces" gorm:"type:int;comment:发光面数"` 
    WhiteBoardPrice float64 `json:"whiteBoardPrice" gorm:"type:varchar(255);comment:白板"` 
    Region string `json:"region" gorm:"type:varchar(255);comment:地区"` 
    ExtraAttributes string `json:"extraAttributes" gorm:"type:json;comment:额外属性"` 
    models.ModelTime
    models.ControlBy
}

func (QuotationsV1) TableName() string {
    return "quotations_v1"
}

func (e *QuotationsV1) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *QuotationsV1) GetId() interface{} {
	return e.Id
}