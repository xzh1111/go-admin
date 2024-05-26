package dto

import (
	"go-admin/app/bj/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type StylesGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:styles" comment:"款式"`
    Price float64 `form:"price"  search:"type:exact;column:price;table:styles" comment:"价格"`
    StylesOrder
}

type StylesOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:styles"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:styles"`
    Price string `form:"priceOrder"  search:"type:order;column:price;table:styles"`
    Image string `form:"imageOrder"  search:"type:order;column:image;table:styles"`
    
}

func (m *StylesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type StylesInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"款式"`
    Price float64 `json:"price" comment:"价格"`
    Image string `json:"image" comment:"图片"`
    common.ControlBy
}

func (s *StylesInsertReq) Generate(model *models.Styles)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Price = s.Price
    model.Image = s.Image
}

func (s *StylesInsertReq) GetId() interface{} {
	return s.Id
}

type StylesUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"款式"`
    Price float64 `json:"price" comment:"价格"`
    Image string `json:"image" comment:"图片"`
    common.ControlBy
}

func (s *StylesUpdateReq) Generate(model *models.Styles)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Price = s.Price
    model.Image = s.Image
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
