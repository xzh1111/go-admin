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

type QuotationsV1 struct {
	service.Service
}

// GetPage 获取QuotationsV1列表
func (e *QuotationsV1) GetPage(c *dto.QuotationsV1GetPageReq, p *actions.DataPermission, list *[]models.QuotationsV1, count *int64) error {
	var err error
	var data models.QuotationsV1

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("QuotationsV1Service GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取QuotationsV1对象
func (e *QuotationsV1) Get(d *dto.QuotationsV1GetReq, p *actions.DataPermission, model *models.QuotationsV1) error {
	var data models.QuotationsV1

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetQuotationsV1 error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建QuotationsV1对象
func (e *QuotationsV1) Insert(c *dto.QuotationsV1InsertReq) error {
    var err error
    var data models.QuotationsV1
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("QuotationsV1Service Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改QuotationsV1对象
func (e *QuotationsV1) Update(c *dto.QuotationsV1UpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.QuotationsV1{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("QuotationsV1Service Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除QuotationsV1
func (e *QuotationsV1) Remove(d *dto.QuotationsV1DeleteReq, p *actions.DataPermission) error {
	var data models.QuotationsV1

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveQuotationsV1 error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
