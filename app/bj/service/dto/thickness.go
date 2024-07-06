package dto

import (

	"go-admin/app/bj/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ThicknessGetPageReq struct {
	dto.Pagination     `search:"-"`
    Thickness string `form:"thickness"  search:"type:exact;column:thickness;table:thickness" comment:"板厚"`
    ThicknessOrder
}

type ThicknessOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:thickness"`
    Thickness string `form:"thicknessOrder"  search:"type:order;column:thickness;table:thickness"`
    TkKg string `form:"tkKgOrder"  search:"type:order;column:tk_kg;table:thickness"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:thickness"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:thickness"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:thickness"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:thickness"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:thickness"`
    
}

func (m *ThicknessGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ThicknessInsertReq struct {
    Id int `json:"-" comment:""` // 
    Thickness string `json:"thickness" comment:"板厚"`
    TkKg float64 `json:"tkKg" comment:"每平方米重量"`
    common.ControlBy
}

func (s *ThicknessInsertReq) Generate(model *models.Thickness)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Thickness = s.Thickness
    model.TkKg = s.TkKg
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ThicknessInsertReq) GetId() interface{} {
	return s.Id
}

type ThicknessUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Thickness string `json:"thickness" comment:"板厚"`
    TkKg float64 `json:"tkKg" comment:"每平方米重量"`
    common.ControlBy
}

func (s *ThicknessUpdateReq) Generate(model *models.Thickness)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Thickness = s.Thickness
    model.TkKg = s.TkKg
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ThicknessUpdateReq) GetId() interface{} {
	return s.Id
}

// ThicknessGetReq 功能获取请求参数
type ThicknessGetReq struct {
     Id int `uri:"id"`
}
func (s *ThicknessGetReq) GetId() interface{} {
	return s.Id
}

// ThicknessDeleteReq 功能删除请求参数
type ThicknessDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ThicknessDeleteReq) GetId() interface{} {
	return s.Ids
}
