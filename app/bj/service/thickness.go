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

type Thickness struct {
	service.Service
}

// GetPage 获取Thickness列表
func (e *Thickness) GetPage(c *dto.ThicknessGetPageReq, p *actions.DataPermission, list *[]models.Thickness, count *int64) error {
	var err error
	var data models.Thickness

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ThicknessService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Thickness对象
func (e *Thickness) Get(d *dto.ThicknessGetReq, p *actions.DataPermission, model *models.Thickness) error {
	var data models.Thickness

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetThickness error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Thickness对象
func (e *Thickness) Insert(c *dto.ThicknessInsertReq) error {
    var err error
    var data models.Thickness
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ThicknessService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Thickness对象
func (e *Thickness) Update(c *dto.ThicknessUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Thickness{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("ThicknessService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Thickness
func (e *Thickness) Remove(d *dto.ThicknessDeleteReq, p *actions.DataPermission) error {
	var data models.Thickness

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveThickness error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
