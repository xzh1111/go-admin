package models

import (
	"sync"
	"time"

	log "github.com/go-admin-team/go-admin-core/logger"

	"gorm.io/gorm"
)

type FormulaMapper struct {
	FormulaKeyToFormula map[string]string
	NameToFormula       map[string]string
	formulas            *[]BjFormula
	mu                  sync.RWMutex
	db                  *gorm.DB
}

var DefaultFormulaMapper *FormulaMapper

func InitFormulaMapper(db *gorm.DB) {
	DefaultFormulaMapper = &FormulaMapper{
		db: db,
	}

	err := DefaultFormulaMapper.sync()
	if err != nil {
		log.Errorf("FormulaMapper init error:%s", err.Error())
		return
	}
	log.Infof("FormulaMapper init success")
	go func() {
		var ticker = time.NewTicker(2 * time.Minute)
		for range ticker.C {
			err := DefaultFormulaMapper.sync()
			if err != nil {
				log.Errorf("FormulaMapper sync error:%s", err.Error())
			} else {
				log.Infof("FormulaMapper sync success")
			}
		}
	}()
}

func (m *FormulaMapper) sync() error {
	var data BjFormula

	formulas := &[]BjFormula{}

	err := m.db.Order("priority ASC").Model(&data).Find(formulas).Limit(-1).Offset(-1).Error
	if err != nil {
		log.Errorf("FormulaMapper sync error:%s, %s", err.Error())
		return err
	}
	m.UpdateMaps(formulas)
	return nil
}

func (m *FormulaMapper) UpdateMaps(formulas *[]BjFormula) {
	tmpFormulaKeyToFormula := make(map[string]string)
	tmpNameToFormula := make(map[string]string)

	for _, formula := range *formulas {
		tmpFormulaKeyToFormula[formula.FormulaKey] = formula.Formula
		tmpNameToFormula[formula.Name] = formula.Formula
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.FormulaKeyToFormula = tmpFormulaKeyToFormula
	m.NameToFormula = tmpNameToFormula
	m.formulas = formulas
	log.Debugf("FormulaMapper UpdateMaps %v", len(*formulas))
}

func GetFormulaByKey(key string) (string, bool) {
	DefaultFormulaMapper.mu.RLock()
	defer DefaultFormulaMapper.mu.RUnlock()
	formula, ok := DefaultFormulaMapper.FormulaKeyToFormula[key]
	return formula, ok
}

func GetFormulaByName(name string) (string, bool) {
	DefaultFormulaMapper.mu.RLock()
	defer DefaultFormulaMapper.mu.RUnlock()
	formula, ok := DefaultFormulaMapper.NameToFormula[name]
	return formula, ok
}

func GetFormulas() []BjFormula {
	DefaultFormulaMapper.mu.RLock()
	defer DefaultFormulaMapper.mu.RUnlock()
	return *DefaultFormulaMapper.formulas
}
