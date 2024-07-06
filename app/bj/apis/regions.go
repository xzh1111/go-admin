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

type Regions struct {
	api.Api
}

// GetPage 获取Regions列表
// @Summary 获取Regions列表
// @Description 获取Regions列表
// @Tags Regions
// @Param name query string false "地区"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Regions}} "{"code": 200, "data": [...]}"
// @Router /api/v1/regions [get]
// @Security Bearer
func (e Regions) GetPage(c *gin.Context) {
    req := dto.RegionsGetPageReq{}
    s := service.Regions{}
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
	list := make([]models.Regions, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Regions失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Regions
// @Summary 获取Regions
// @Description 获取Regions
// @Tags Regions
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Regions} "{"code": 200, "data": [...]}"
// @Router /api/v1/regions/{id} [get]
// @Security Bearer
func (e Regions) Get(c *gin.Context) {
	req := dto.RegionsGetReq{}
	s := service.Regions{}
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
	var object models.Regions

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Regions失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Regions
// @Summary 创建Regions
// @Description 创建Regions
// @Tags Regions
// @Accept application/json
// @Product application/json
// @Param data body dto.RegionsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/regions [post]
// @Security Bearer
func (e Regions) Insert(c *gin.Context) {
    req := dto.RegionsInsertReq{}
    s := service.Regions{}
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
		e.Error(500, err, fmt.Sprintf("创建Regions失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Regions
// @Summary 修改Regions
// @Description 修改Regions
// @Tags Regions
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RegionsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/regions/{id} [put]
// @Security Bearer
func (e Regions) Update(c *gin.Context) {
    req := dto.RegionsUpdateReq{}
    s := service.Regions{}
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
		e.Error(500, err, fmt.Sprintf("修改Regions失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Regions
// @Summary 删除Regions
// @Description 删除Regions
// @Tags Regions
// @Param data body dto.RegionsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/regions [delete]
// @Security Bearer
func (e Regions) Delete(c *gin.Context) {
    s := service.Regions{}
    req := dto.RegionsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Regions失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
