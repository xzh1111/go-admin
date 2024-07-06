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

type Thickness struct {
	api.Api
}

// GetPage 获取板厚列表
// @Summary 获取板厚列表
// @Description 获取板厚列表
// @Tags 板厚
// @Param thickness query string false "板厚"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Thickness}} "{"code": 200, "data": [...]}"
// @Router /api/v1/thickness [get]
// @Security Bearer
func (e Thickness) GetPage(c *gin.Context) {
    req := dto.ThicknessGetPageReq{}
    s := service.Thickness{}
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
	list := make([]models.Thickness, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取板厚失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取板厚
// @Summary 获取板厚
// @Description 获取板厚
// @Tags 板厚
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Thickness} "{"code": 200, "data": [...]}"
// @Router /api/v1/thickness/{id} [get]
// @Security Bearer
func (e Thickness) Get(c *gin.Context) {
	req := dto.ThicknessGetReq{}
	s := service.Thickness{}
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
	var object models.Thickness

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取板厚失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建板厚
// @Summary 创建板厚
// @Description 创建板厚
// @Tags 板厚
// @Accept application/json
// @Product application/json
// @Param data body dto.ThicknessInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/thickness [post]
// @Security Bearer
func (e Thickness) Insert(c *gin.Context) {
    req := dto.ThicknessInsertReq{}
    s := service.Thickness{}
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
		e.Error(500, err, fmt.Sprintf("创建板厚失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改板厚
// @Summary 修改板厚
// @Description 修改板厚
// @Tags 板厚
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ThicknessUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/thickness/{id} [put]
// @Security Bearer
func (e Thickness) Update(c *gin.Context) {
    req := dto.ThicknessUpdateReq{}
    s := service.Thickness{}
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
		e.Error(500, err, fmt.Sprintf("修改板厚失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除板厚
// @Summary 删除板厚
// @Description 删除板厚
// @Tags 板厚
// @Param data body dto.ThicknessDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/thickness [delete]
// @Security Bearer
func (e Thickness) Delete(c *gin.Context) {
    s := service.Thickness{}
    req := dto.ThicknessDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除板厚失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
