package dto

import (
	"fmt"
	"reflect"
)

// SetByFieldName 根据字段名设置结构体的值
func (q *QuotationsV1Req) SetByFieldName(fieldName string, value interface{}) error {
	// 获取结构体的反射类型对象
	structType := reflect.TypeOf(q)
	// 获取字段
	_, found := structType.Elem().FieldByName(fieldName)
	if !found {
		return fmt.Errorf("field %s not found", fieldName)
	}

	// 获取字段的反射值对象
	fieldValue := reflect.ValueOf(q).Elem().FieldByName(fieldName)

	// 设置字段的值
	if fieldValue.CanSet() {
		fieldValue.Set(reflect.ValueOf(value))
		return nil
	} else {
		return fmt.Errorf("cannot set field %s", fieldName)
	}
}

// GetByFieldName 根据字段名获取结构体的值
func (q *QuotationsV1Req) GetByFieldName(fieldName string) (interface{}, error) {
	// 获取结构体的反射类型对象
	structType := reflect.TypeOf(q)
	// 获取字段
	_, found := structType.Elem().FieldByName(fieldName)
	if !found {
		return nil, fmt.Errorf("field %s not found", fieldName)
	}

	// 获取字段的反射值对象
	fieldValue := reflect.ValueOf(q).Elem().FieldByName(fieldName)

	// 返回字段的值
	return fieldValue.Interface(), nil
}
