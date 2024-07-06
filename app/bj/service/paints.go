package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/bj/models"
	"go-admin/app/bj/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Paints struct {
	service.Service
}

// GetPage 获取Paints列表
func (e *Paints) GetPage(c *dto.PaintsGetPageReq, p *actions.DataPermission, list *[]models.Paints, count *int64) error {
	var err error
	var data models.Paints

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("PaintsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Paints对象
func (e *Paints) Get(d *dto.PaintsGetReq, p *actions.DataPermission, model *models.Paints) error {
	var data models.Paints

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetPaints error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Paints对象
func (e *Paints) Insert(c *dto.PaintsInsertReq) error {
    var err error
    var data models.Paints
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("PaintsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Paints对象
func (e *Paints) Update(c *dto.PaintsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Paints{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("PaintsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Paints
func (e *Paints) Remove(d *dto.PaintsDeleteReq, p *actions.DataPermission) error {
	var data models.Paints

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemovePaints error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
