package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/bj/models"
	"go-admin/app/bj/service"
	"go-admin/app/bj/service/dto"
	"go-admin/common/actions"
)

type BjFormula struct {
	api.Api
}

// GetPage 获取BjFormula列表
// @Summary 获取BjFormula列表
// @Description 获取BjFormula列表
// @Tags BjFormula
// @Param name query string false "公式名"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BjFormula}} "{"code": 200, "data": [...]}"
// @Router /api/v1/bj-formula [get]
// @Security Bearer
func (e BjFormula) GetPage(c *gin.Context) {
    req := dto.BjFormulaGetPageReq{}
    s := service.BjFormula{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.BjFormula, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取BjFormula失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取BjFormula
// @Summary 获取BjFormula
// @Description 获取BjFormula
// @Tags BjFormula
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BjFormula} "{"code": 200, "data": [...]}"
// @Router /api/v1/bj-formula/{id} [get]
// @Security Bearer
func (e BjFormula) Get(c *gin.Context) {
	req := dto.BjFormulaGetReq{}
	s := service.BjFormula{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.BjFormula

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取BjFormula失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建BjFormula
// @Summary 创建BjFormula
// @Description 创建BjFormula
// @Tags BjFormula
// @Accept application/json
// @Product application/json
// @Param data body dto.BjFormulaInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/bj-formula [post]
// @Security Bearer
func (e BjFormula) Insert(c *gin.Context) {
    req := dto.BjFormulaInsertReq{}
    s := service.BjFormula{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建BjFormula失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改BjFormula
// @Summary 修改BjFormula
// @Description 修改BjFormula
// @Tags BjFormula
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BjFormulaUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/bj-formula/{id} [put]
// @Security Bearer
func (e BjFormula) Update(c *gin.Context) {
    req := dto.BjFormulaUpdateReq{}
    s := service.BjFormula{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改BjFormula失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除BjFormula
// @Summary 删除BjFormula
// @Description 删除BjFormula
// @Tags BjFormula
// @Param data body dto.BjFormulaDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/bj-formula [delete]
// @Security Bearer
func (e BjFormula) Delete(c *gin.Context) {
    s := service.BjFormula{}
    req := dto.BjFormulaDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除BjFormula失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
