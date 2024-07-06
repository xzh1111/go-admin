package dto

import (

	"go-admin/app/bj/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RegionsGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:regions" comment:"地区"`
    RegionsOrder
}

type RegionsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:regions"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:regions"`
    Base1 string `form:"base1Order"  search:"type:order;column:base1;table:regions"`
    Base2 string `form:"base2Order"  search:"type:order;column:base2;table:regions"`
    
}

func (m *RegionsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RegionsInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"地区"`
    Base1 float64 `json:"base1" comment:"基数1"`
    Base2 float64 `json:"base2" comment:"基数2"`
    common.ControlBy
}

func (s *RegionsInsertReq) Generate(model *models.Regions)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Base1 = s.Base1
    model.Base2 = s.Base2
}

func (s *RegionsInsertReq) GetId() interface{} {
	return s.Id
}

type RegionsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"地区"`
    Base1 float64 `json:"base1" comment:"基数1"`
    Base2 float64 `json:"base2" comment:"基数2"`
    common.ControlBy
}

func (s *RegionsUpdateReq) Generate(model *models.Regions)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Base1 = s.Base1
    model.Base2 = s.Base2
}

func (s *RegionsUpdateReq) GetId() interface{} {
	return s.Id
}

// RegionsGetReq 功能获取请求参数
type RegionsGetReq struct {
     Id int `uri:"id"`
}
func (s *RegionsGetReq) GetId() interface{} {
	return s.Id
}

// RegionsDeleteReq 功能删除请求参数
type RegionsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RegionsDeleteReq) GetId() interface{} {
	return s.Ids
}
