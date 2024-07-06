package dto

import (
	"go-admin/app/bj/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PaintsGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:paints" comment:""`
    PaintsOrder
}

type PaintsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:paints"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:paints"`
    Price float64 `form:"priceOrder"  search:"type:order;column:price;table:paints"`
    Image string `form:"imageOrder"  search:"type:order;column:image;table:paints"`
    
}

func (m *PaintsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type PaintsInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:""`
    Price float64 `json:"price" comment:""`
    Image string `json:"image" comment:""`
    common.ControlBy
}

func (s *PaintsInsertReq) Generate(model *models.Paints)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Price = s.Price
    model.Image = s.Image
}

func (s *PaintsInsertReq) GetId() interface{} {
	return s.Id
}

type PaintsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:""`
    Price float64 `json:"price" comment:""`
    Image string `json:"image" comment:""`
    common.ControlBy
}

func (s *PaintsUpdateReq) Generate(model *models.Paints)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Price = s.Price
    model.Image = s.Image
}

func (s *PaintsUpdateReq) GetId() interface{} {
	return s.Id
}

// PaintsGetReq 功能获取请求参数
type PaintsGetReq struct {
     Id int `uri:"id"`
}
func (s *PaintsGetReq) GetId() interface{} {
	return s.Id
}

// PaintsDeleteReq 功能删除请求参数
type PaintsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *PaintsDeleteReq) GetId() interface{} {
	return s.Ids
}
