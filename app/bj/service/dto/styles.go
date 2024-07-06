package dto

import (

	"go-admin/app/bj/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type StylesGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:styles" comment:"款式"`
    StylesOrder
}

type StylesOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:styles"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:styles"`
    Price string `form:"priceOrder"  search:"type:order;column:price;table:styles"`
    Image string `form:"imageOrder"  search:"type:order;column:image;table:styles"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:styles"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:styles"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:styles"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:styles"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:styles"`
    ImageExt string `form:"imageExtOrder"  search:"type:order;column:image_ext;table:styles"`
    WeightFormula string `form:"weightFormulaOrder"  search:"type:order;column:weight_formula;table:styles"`
    
}

func (m *StylesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type StylesInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"款式"`
    Price float64 `json:"price" comment:"款式单价"`
    Image string `json:"image" comment:""`
    ImageExt string `json:"imageExt" comment:""`
    WeightFormula string `json:"weightFormula" comment:"款式重量公式"`
    common.ControlBy
}

func (s *StylesInsertReq) Generate(model *models.Styles)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Price = s.Price
    model.Image = s.Image
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.ImageExt = s.ImageExt
    model.WeightFormula = s.WeightFormula
}

func (s *StylesInsertReq) GetId() interface{} {
	return s.Id
}

type StylesUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"款式"`
    Price float64 `json:"price" comment:"款式单价"`
    Image string `json:"image" comment:""`
    ImageExt string `json:"imageExt" comment:""`
    WeightFormula string `json:"weightFormula" comment:"款式重量公式"`
    common.ControlBy
}

func (s *StylesUpdateReq) Generate(model *models.Styles)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Price = s.Price
    model.Image = s.Image
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.ImageExt = s.ImageExt
    model.WeightFormula = s.WeightFormula
}

func (s *StylesUpdateReq) GetId() interface{} {
	return s.Id
}

// StylesGetReq 功能获取请求参数
type StylesGetReq struct {
     Id int `uri:"id"`
}
func (s *StylesGetReq) GetId() interface{} {
	return s.Id
}

// StylesDeleteReq 功能删除请求参数
type StylesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *StylesDeleteReq) GetId() interface{} {
	return s.Ids
}
