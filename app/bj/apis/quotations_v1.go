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

type QuotationsV1 struct {
	api.Api
}

// GetPage 获取QuotationsV1列表
// @Summary 获取QuotationsV1列表
// @Description 获取QuotationsV1列表
// @Tags QuotationsV1
// @Param orderNumber query string false "单号"
// @Param style query string false "款式"
// @Param color query string false "颜色"
// @Param region query string false "地区"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.QuotationsV1}} "{"code": 200, "data": [...]}"
// @Router /api/v1/quotations [get]
// @Security Bearer
func (e QuotationsV1) GetPage(c *gin.Context) {
	req := dto.QuotationsV1GetPageReq{}
	s := service.QuotationsV1{}
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
	userId := user.GetUserIdStr(c)
	req.UserId = userId

	p := actions.GetPermissionFromContext(c)
	list := make([]models.QuotationsV1, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取QuotationsV1失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取QuotationsV1
// @Summary 获取QuotationsV1
// @Description 获取QuotationsV1
// @Tags QuotationsV1
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.QuotationsV1} "{"code": 200, "data": [...]}"
// @Router /api/v1/quotations/{id} [get]
// @Security Bearer
func (e QuotationsV1) Get(c *gin.Context) {
	req := dto.QuotationsV1GetReq{}
	s := service.QuotationsV1{}
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
	var object models.QuotationsV1

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取QuotationsV1失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建QuotationsV1
// @Summary 创建QuotationsV1
// @Description 创建QuotationsV1
// @Tags QuotationsV1
// @Accept application/json
// @Product application/json
// @Param data body dto.QuotationsV1InsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/quotations [post]
// @Security Bearer
func (e QuotationsV1) Insert(c *gin.Context) {
	req := dto.QuotationsV1InsertReq{}
	s := service.QuotationsV1{}
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
		e.Error(500, err, fmt.Sprintf("创建QuotationsV1失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改QuotationsV1
// @Summary 修改QuotationsV1
// @Description 修改QuotationsV1
// @Tags QuotationsV1
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.QuotationsV1UpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/quotations/{id} [put]
// @Security Bearer
func (e QuotationsV1) Update(c *gin.Context) {
	req := dto.QuotationsV1UpdateReq{}
	s := service.QuotationsV1{}
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
		e.Error(500, err, fmt.Sprintf("修改QuotationsV1失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除QuotationsV1
// @Summary 删除QuotationsV1
// @Description 删除QuotationsV1
// @Tags QuotationsV1
// @Param data body dto.QuotationsV1DeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/quotations [delete]
// @Security Bearer
func (e QuotationsV1) Delete(c *gin.Context) {
	s := service.QuotationsV1{}
	req := dto.QuotationsV1DeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除QuotationsV1失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
