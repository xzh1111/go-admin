package dto

import (
    "go-admin/app/bj/models"
    "go-admin/common/dto"
    common "go-admin/common/models"
)

type QuotationsV1GetPageReq struct {
    dto.Pagination `search:"-"`
    OrderNumber    string `form:"orderNumber"  search:"type:exact;column:order_number;table:quotations_v1" comment:"单号"`
    Style          string `form:"style"  search:"type:exact;column:style;table:quotations_v1" comment:"款式"`
    Color          string `form:"color"  search:"type:exact;column:color;table:quotations_v1" comment:"颜色"`
    Region         string `form:"region"  search:"type:exact;column:region;table:quotations_v1" comment:"地区"`
    QuotationsV1Order
}

type QuotationsV1Order struct {
    Id                       string `form:"idOrder"  search:"type:order;column:id;table:quotations_v1"`
    OrderNumber              string `form:"orderNumberOrder"  search:"type:order;column:order_number;table:quotations_v1"`
    Style                    string `form:"styleOrder"  search:"type:order;column:style;table:quotations_v1"`
    Color                    string `form:"colorOrder"  search:"type:order;column:color;table:quotations_v1"`
    Length                   string `form:"lengthOrder"  search:"type:order;column:length;table:quotations_v1"`
    Width                    string `form:"widthOrder"  search:"type:order;column:width;table:quotations_v1"`
    Height                   string `form:"heightOrder"  search:"type:order;column:height;table:quotations_v1"`
    Thickness                string `form:"thicknessOrder"  search:"type:order;column:thickness;table:quotations_v1"`
    PipeWeightAndExtraWeight string `form:"pipeWeightAndExtraWeightOrder"  search:"type:order;column:pipe_weight_and_extra_weight;table:quotations_v1"`
    Quantity                 string `form:"quantityOrder"  search:"type:order;column:quantity;table:quotations_v1"`
    SingleIronWeight         string `form:"singleIronWeightOrder"  search:"type:order;column:single_iron_weight;table:quotations_v1"`
    IronPrice                string `form:"ironPriceOrder"  search:"type:order;column:iron_price;table:quotations_v1"`
    BakingPaintPrice         string `form:"bakingPaintPriceOrder"  search:"type:order;column:baking_paint_price;table:quotations_v1"`
    NeedSilkScreen           string `form:"needSilkScreenOrder"  search:"type:order;column:need_silk_screen;table:quotations_v1"`
    NumberOfOpenTemplates    string `form:"numberOfOpenTemplatesOrder"  search:"type:order;column:number_of_open_templates;table:quotations_v1"`
    OpenBoardFee             string `form:"openBoardFeeOrder"  search:"type:order;column:open_board_fee;table:quotations_v1"`
    NumberOfSilkScreenFaces  string `form:"numberOfSilkScreenFacesOrder"  search:"type:order;column:number_of_silk_screen_faces;table:quotations_v1"`
    NumberOfColors           string `form:"numberOfColorsOrder"  search:"type:order;column:number_of_colors;table:quotations_v1"`
    PrintPrice               string `form:"printPriceOrder"  search:"type:order;column:print_price;table:quotations_v1"`
    Freight                  string `form:"freightOrder"  search:"type:order;column:freight;table:quotations_v1"`
    PackagingFee             string `form:"packagingFeeOrder"  search:"type:order;column:packaging_fee;table:quotations_v1"`
    Cost                     string `form:"costOrder"  search:"type:order;column:cost;table:quotations_v1"`
    Profit                   string `form:"profitOrder"  search:"type:order;column:profit;table:quotations_v1"`
    FinalQuotation           string `form:"finalQuotationOrder"  search:"type:order;column:final_quotation;table:quotations_v1"`
    CreatedAt                string `form:"createdAtOrder"  search:"type:order;column:created_at;table:quotations_v1"`
    UpdatedAt                string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:quotations_v1"`
    DeletedAt                string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:quotations_v1"`
    CreateBy                 string `form:"createByOrder"  search:"type:order;column:create_by;table:quotations_v1"`
    UpdateBy                 string `form:"updateByOrder"  search:"type:order;column:update_by;table:quotations_v1"`
    PowerSource              string `form:"powerSourceOrder"  search:"type:order;column:power_source;table:quotations_v1"`
    LightStripCost           string `form:"lightStripCostOrder"  search:"type:order;column:light_strip_cost;table:quotations_v1"`
    NumberOfIlluminatedFaces string `form:"numberOfIlluminatedFacesOrder"  search:"type:order;column:number_of_illuminated_faces;table:quotations_v1"`
    WhiteBoardPrice          string `form:"whiteBoardOrder"  search:"type:order;column:white_board_price;table:quotations_v1"`
    Region                   string `form:"regionOrder"  search:"type:order;column:region;table:quotations_v1"`
    ExtraAttributes          string `form:"extraAttributesOrder"  search:"type:order;column:extra_attributes;table:quotations_v1"`
}

func (m *QuotationsV1GetPageReq) GetNeedSearch() interface{} {
    return *m
}

type QuotationsV1Req struct {
    Id                       int     `uri:"id" comment:"主键ID"` // 主键ID
    OrderNumber              string  `json:"orderNumber" comment:"单号"`
    Style                    string  `json:"style" comment:"款式"`
    Color                    string  `json:"color" comment:"颜色"`
    Length                   int     `json:"length" comment:"长"`
    Width                    int     `json:"width" comment:"宽"`
    Height                   int     `json:"height" comment:"高"`
    Thickness                float64 `json:"thickness" comment:"板厚"`
    PipeWeightAndExtraWeight float64 `json:"pipeWeightAndExtraWeight" comment:"管重及额外重量"`
    Quantity                 int     `json:"quantity" comment:"数量"`
    SingleIronWeight         float64 `json:"singleIronWeight" comment:"单个白铁重量"`
    IronPrice                float64 `json:"ironPrice" comment:"白铁价格"`
    BakingPaintPrice         float64 `json:"bakingPaintPrice" comment:"烤漆价格"`
    NeedSilkScreen           int     `json:"needSilkScreen" comment:"是否需要丝印"`
    NumberOfOpenTemplates    int     `json:"numberOfOpenTemplates" comment:"开模板数"`
    OpenBoardFee             float64 `json:"openBoardFee" comment:"开板费"`
    NumberOfSilkScreenFaces  int     `json:"numberOfSilkScreenFaces" comment:"丝印面数"`
    NumberOfColors           int     `json:"numberOfColors" comment:"印几个颜色"`
    PrintPrice               float64 `json:"printPrice" comment:"印的价格"`
    Freight                  float64 `json:"freight" comment:"运费"`
    PackagingFee             float64 `json:"packagingFee" comment:"包装费"`
    Cost                     float64 `json:"cost" comment:"成本"`
    Profit                   float64 `json:"profit" comment:"利润"`
    FinalQuotation           float64 `json:"finalQuotation" comment:"最终报价"`
    PowerSource              string  `json:"powerSource" comment:"插电或者太阳能"`
    LightStripCost           float64 `json:"lightStripCost" comment:"灯带费用"`
    NumberOfIlluminatedFaces int     `json:"numberOfIlluminatedFaces" comment:"发光面数"`
    WhiteBoardPrice          float64 `json:"whiteBoardPrice" comment:"白板"`
    Region                   string  `json:"region" comment:"地区"`
    ExtraAttributes          string  `json:"extraAttributes" comment:"额外属性"`
    // --额外附加
    ThicknessWeight float64 `json:"thickness_weight" comment:"每平方重量"`
    StylePrice      float64 `json:"style_price" comment:"款式价格"`
    ColorPrice      float64 `json:"color_price" comment:"颜色价格"`
    WeightFormula   string  `json:"weight_formula" comment:"重量公式"`
    LocationBase1   float64 `json:"location_base1" comment:"位置1"`
    LocationBase2   float64 `json:"location_base2" comment:"位置2"`
    // --^

    common.ControlBy
}

func (s *QuotationsV1Req) GetName() string {
    return "QuotationsV1Req"
}

type QuotationsV1InsertReq QuotationsV1Req

func (s *QuotationsV1InsertReq) Generate(model *models.QuotationsV1) {
    if s.Id == 0 {
        model.Model = common.Model{Id: s.Id}
    }
    model.OrderNumber = s.OrderNumber
    model.Style = s.Style
    model.Color = s.Color
    model.Length = s.Length
    model.Width = s.Width
    model.Height = s.Height
    model.Thickness = s.Thickness
    model.PipeWeightAndExtraWeight = s.PipeWeightAndExtraWeight
    model.Quantity = s.Quantity
    model.SingleIronWeight = s.SingleIronWeight
    model.IronPrice = s.IronPrice
    model.BakingPaintPrice = s.BakingPaintPrice
    model.NeedSilkScreen = s.NeedSilkScreen
    model.NumberOfOpenTemplates = s.NumberOfOpenTemplates
    model.OpenBoardFee = s.OpenBoardFee
    model.NumberOfSilkScreenFaces = s.NumberOfSilkScreenFaces
    model.NumberOfColors = s.NumberOfColors
    model.PrintPrice = s.PrintPrice
    model.Freight = s.Freight
    model.PackagingFee = s.PackagingFee
    model.Cost = s.Cost
    model.Profit = s.Profit
    model.FinalQuotation = s.FinalQuotation
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.PowerSource = s.PowerSource
    model.LightStripCost = s.LightStripCost
    model.NumberOfIlluminatedFaces = s.NumberOfIlluminatedFaces
    model.WhiteBoardPrice = s.WhiteBoardPrice
    model.Region = s.Region
    model.ExtraAttributes = s.ExtraAttributes
}

func (s *QuotationsV1InsertReq) GetId() interface{} {
    return s.Id
}

type QuotationsV1UpdateReq QuotationsV1Req

func (s *QuotationsV1UpdateReq) Generate(model *models.QuotationsV1) {
    if s.Id == 0 {
        model.Model = common.Model{Id: s.Id}
    }
    model.OrderNumber = s.OrderNumber
    model.Style = s.Style
    model.Color = s.Color
    model.Length = s.Length
    model.Width = s.Width
    model.Height = s.Height
    model.Thickness = s.Thickness
    model.PipeWeightAndExtraWeight = s.PipeWeightAndExtraWeight
    model.Quantity = s.Quantity
    model.SingleIronWeight = s.SingleIronWeight
    model.IronPrice = s.IronPrice
    model.BakingPaintPrice = s.BakingPaintPrice
    model.NeedSilkScreen = s.NeedSilkScreen
    model.NumberOfOpenTemplates = s.NumberOfOpenTemplates
    model.OpenBoardFee = s.OpenBoardFee
    model.NumberOfSilkScreenFaces = s.NumberOfSilkScreenFaces
    model.NumberOfColors = s.NumberOfColors
    model.PrintPrice = s.PrintPrice
    model.Freight = s.Freight
    model.PackagingFee = s.PackagingFee
    model.Cost = s.Cost
    model.Profit = s.Profit
    model.FinalQuotation = s.FinalQuotation
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.PowerSource = s.PowerSource
    model.LightStripCost = s.LightStripCost
    model.NumberOfIlluminatedFaces = s.NumberOfIlluminatedFaces
    model.WhiteBoardPrice = s.WhiteBoardPrice
    model.Region = s.Region
    model.ExtraAttributes = s.ExtraAttributes
}

func (s *QuotationsV1UpdateReq) GetId() interface{} {
    return s.Id
}

// QuotationsV1GetReq 功能获取请求参数
type QuotationsV1GetReq struct {
    Id int `uri:"id"`
}

func (s *QuotationsV1GetReq) GetId() interface{} {
    return s.Id
}

// QuotationsV1DeleteReq 功能删除请求参数
type QuotationsV1DeleteReq struct {
    Ids []int `json:"ids"`
}

func (s *QuotationsV1DeleteReq) GetId() interface{} {
    return s.Ids
}
