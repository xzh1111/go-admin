package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysContentSearch struct {
	dto.Pagination `search:"-"`
}

func (m *SysContentSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *SysContentSearch) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("ShouldBind error: %s", err.Error())
	}
	return err
}

func (m *SysContentSearch) Generate() dto.Index {
	o := *m
	return &o
}

type SysContentControl struct {
	ID      int    `uri:"ID" comment:""` //
	CateId  string `json:"cateId" comment:""`
	Name    string `json:"name" comment:""`
	Status  string `json:"status" comment:""`
	Img     string `json:"img" comment:""`
	Content string `json:"content" comment:""`
	Remark  string `json:"remark" comment:""`
	Sort    string `json:"sort" comment:""`
}

func (s *SysContentControl) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
	err := ctx.ShouldBindUri(s)
	if err != nil {
		log.Debugf("ShouldBindUri error: %s", err.Error())
		return err
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		log.Debugf("ShouldBind error: %s", err.Error())
	}
	return err
}

func (s *SysContentControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *SysContentControl) GenerateM() (common.ActiveRecord, error) {
	return &models.SysContent{
		Id:      s.ID,
		CateId:  s.CateId,
		Name:    s.Name,
		Status:  s.Status,
		Img:     s.Img,
		Content: s.Content,
		Remark:  s.Remark,
		Sort:    s.Sort,
	}, nil
}

func (s *SysContentControl) GetId() interface{} {
	return s.ID
}

type SysContentById struct {
	dto.ObjectById
}

func (s *SysContentById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *SysContentById) GenerateM() (common.ActiveRecord, error) {
	return &models.SysContent{}, nil
}
