package dto

import (

	"go-admin/app/bj/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BjFormulaGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:bj_formula" comment:"公式名"`
    BjFormulaOrder
}

type BjFormulaOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bj_formula"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:bj_formula"`
    Formula string `form:"formulaOrder"  search:"type:order;column:formula;table:bj_formula"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bj_formula"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bj_formula"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bj_formula"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bj_formula"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bj_formula"`
    FormulaKey string `form:"formulaKeyOrder"  search:"type:order;column:formula_key;table:bj_formula"`
    Priority string `form:"priorityOrder"  search:"type:order;column:priority;table:bj_formula"`
    
}

func (m *BjFormulaGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BjFormulaInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"公式名"`
    Formula string `json:"formula" comment:"公式"`
    FormulaKey string `json:"formulaKey" comment:"公式标识(关联计算)"`
    Priority int `json:"priority" comment:"公式计算优先级"`
    common.ControlBy
}

func (s *BjFormulaInsertReq) Generate(model *models.BjFormula)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Formula = s.Formula
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.FormulaKey = s.FormulaKey
    model.Priority = s.Priority
}

func (s *BjFormulaInsertReq) GetId() interface{} {
	return s.Id
}

type BjFormulaUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"公式名"`
    Formula string `json:"formula" comment:"公式"`
    FormulaKey string `json:"formulaKey" comment:"公式标识(关联计算)"`
    Priority int `json:"priority" comment:"公式计算优先级"`
    common.ControlBy
}

func (s *BjFormulaUpdateReq) Generate(model *models.BjFormula)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Formula = s.Formula
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.FormulaKey = s.FormulaKey
    model.Priority = s.Priority
}

func (s *BjFormulaUpdateReq) GetId() interface{} {
	return s.Id
}

// BjFormulaGetReq 功能获取请求参数
type BjFormulaGetReq struct {
     Id int `uri:"id"`
}
func (s *BjFormulaGetReq) GetId() interface{} {
	return s.Id
}

// BjFormulaDeleteReq 功能删除请求参数
type BjFormulaDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BjFormulaDeleteReq) GetId() interface{} {
	return s.Ids
}
