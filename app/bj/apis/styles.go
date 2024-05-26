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

type Styles struct {
	api.Api
}

// GetPage 获取Styles列表
// @Summary 获取Styles列表
// @Description 获取Styles列表
// @Tags Styles
// @Param name query string false "款式"
// @Param price query int64 false "价格"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Styles}} "{"code": 200, "data": [...]}"
// @Router /api/v1/styles [get]
// @Security Bearer
func (e Styles) GetPage(c *gin.Context) {
    req := dto.StylesGetPageReq{}
    s := service.Styles{}
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
	list := make([]models.Styles, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Styles失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Styles
// @Summary 获取Styles
// @Description 获取Styles
// @Tags Styles
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Styles} "{"code": 200, "data": [...]}"
// @Router /api/v1/styles/{id} [get]
// @Security Bearer
func (e Styles) Get(c *gin.Context) {
	req := dto.StylesGetReq{}
	s := service.Styles{}
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
	var object models.Styles

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Styles失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Styles
// @Summary 创建Styles
// @Description 创建Styles
// @Tags Styles
// @Accept application/json
// @Product application/json
// @Param data body dto.StylesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/styles [post]
// @Security Bearer
func (e Styles) Insert(c *gin.Context) {
    req := dto.StylesInsertReq{}
    s := service.Styles{}
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
		e.Error(500, err, fmt.Sprintf("创建Styles失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Styles
// @Summary 修改Styles
// @Description 修改Styles
// @Tags Styles
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.StylesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/styles/{id} [put]
// @Security Bearer
func (e Styles) Update(c *gin.Context) {
    req := dto.StylesUpdateReq{}
    s := service.Styles{}
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
		e.Error(500, err, fmt.Sprintf("修改Styles失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Styles
// @Summary 删除Styles
// @Description 删除Styles
// @Tags Styles
// @Param data body dto.StylesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/styles [delete]
// @Security Bearer
func (e Styles) Delete(c *gin.Context) {
    s := service.Styles{}
    req := dto.StylesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Styles失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
