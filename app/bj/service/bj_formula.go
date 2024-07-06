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

type BjFormula struct {
	service.Service
}

// GetPage 获取BjFormula列表
func (e *BjFormula) GetPage(c *dto.BjFormulaGetPageReq, p *actions.DataPermission, list *[]models.BjFormula, count *int64) error {
	var err error
	var data models.BjFormula

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BjFormulaService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取BjFormula对象
func (e *BjFormula) Get(d *dto.BjFormulaGetReq, p *actions.DataPermission, model *models.BjFormula) error {
	var data models.BjFormula

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBjFormula error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建BjFormula对象
func (e *BjFormula) Insert(c *dto.BjFormulaInsertReq) error {
    var err error
    var data models.BjFormula
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BjFormulaService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改BjFormula对象
func (e *BjFormula) Update(c *dto.BjFormulaUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.BjFormula{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BjFormulaService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除BjFormula
func (e *BjFormula) Remove(d *dto.BjFormulaDeleteReq, p *actions.DataPermission) error {
	var data models.BjFormula

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBjFormula error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
