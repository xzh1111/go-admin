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

type Paints struct {
	api.Api
}

// GetPage 获取Paints列表
// @Summary 获取Paints列表
// @Description 获取Paints列表
// @Tags Paints
// @Param name query string false ""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Paints}} "{"code": 200, "data": [...]}"
// @Router /api/v1/paints [get]
// @Security Bearer
func (e Paints) GetPage(c *gin.Context) {
    req := dto.PaintsGetPageReq{}
    s := service.Paints{}
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
	list := make([]models.Paints, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Paints失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Paints
// @Summary 获取Paints
// @Description 获取Paints
// @Tags Paints
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Paints} "{"code": 200, "data": [...]}"
// @Router /api/v1/paints/{id} [get]
// @Security Bearer
func (e Paints) Get(c *gin.Context) {
	req := dto.PaintsGetReq{}
	s := service.Paints{}
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
	var object models.Paints

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Paints失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Paints
// @Summary 创建Paints
// @Description 创建Paints
// @Tags Paints
// @Accept application/json
// @Product application/json
// @Param data body dto.PaintsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/paints [post]
// @Security Bearer
func (e Paints) Insert(c *gin.Context) {
    req := dto.PaintsInsertReq{}
    s := service.Paints{}
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
		e.Error(500, err, fmt.Sprintf("创建Paints失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Paints
// @Summary 修改Paints
// @Description 修改Paints
// @Tags Paints
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.PaintsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/paints/{id} [put]
// @Security Bearer
func (e Paints) Update(c *gin.Context) {
    req := dto.PaintsUpdateReq{}
    s := service.Paints{}
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
		e.Error(500, err, fmt.Sprintf("修改Paints失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Paints
// @Summary 删除Paints
// @Description 删除Paints
// @Tags Paints
// @Param data body dto.PaintsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/paints [delete]
// @Security Bearer
func (e Paints) Delete(c *gin.Context) {
    s := service.Paints{}
    req := dto.PaintsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Paints失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
