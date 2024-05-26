package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Styles struct {
	service.Service
}

// GetPage 获取Styles列表
func (e *Styles) GetPage(c *dto.StylesGetPageReq, p *actions.DataPermission, list *[]models.Styles, count *int64) error {
	var err error
	var data models.Styles

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("StylesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Styles对象
func (e *Styles) Get(d *dto.StylesGetReq, p *actions.DataPermission, model *models.Styles) error {
	var data models.Styles

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetStyles error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Styles对象
func (e *Styles) Insert(c *dto.StylesInsertReq) error {
    var err error
    var data models.Styles
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("StylesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Styles对象
func (e *Styles) Update(c *dto.StylesUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Styles{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("StylesService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Styles
func (e *Styles) Remove(d *dto.StylesDeleteReq, p *actions.DataPermission) error {
	var data models.Styles

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveStyles error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
