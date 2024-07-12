package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/app/bj/models"
	"go-admin/app/bj/service/dto"
	"math"

	"github.com/Knetic/govaluate"
)

// 定义自定义函数
var govaluateFunctions = map[string]govaluate.ExpressionFunction{
	"MAX": func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("add function requires exactly 2 arguments")
		}
		a, ok1 := args[0].(float64)
		b, ok2 := args[1].(float64)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("add function requires float64 arguments")
		}
		return math.Max(a, b), nil
	},
	"ROUND": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("round function expects exactly 1 argument")
		}
		switch arg := args[0].(type) {
		case float64:
			return math.Round(arg), nil
		default:
			return nil, fmt.Errorf("round function expects a float64 argument")
		}
	},
	"MIN": func(args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, fmt.Errorf("min function expects at least 2 arguments")
		}
		min := args[0]
		for _, arg := range args[1:] {
			switch a := min.(type) {
			case float64:
				if b, ok := arg.(float64); ok && b < a {
					min = arg
				}
			default:
				return nil, fmt.Errorf("min function expects float64 arguments")
			}
		}
		return min, nil
	},
	// IF函数，当第一个条件满足时使用第二个参数值返回，否则使用第三参数值返回
	"IF": func(args ...interface{}) (interface{}, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("IF function expects exactly 3 arguments")
		}
		condition, ok1 := args[0].(bool)
		if !ok1 {
			return nil, fmt.Errorf("IF function expects a boolean argument")
		}
		if condition {
			return args[1], nil
		}
		return args[2], nil
	},
	
}

func (s *QuotationsV1) Compute(c *dto.QuotationsV1Req) error {

	reqByte, _ := json.Marshal(c)
	environment := make(map[string]interface{}, 0)
	json.Unmarshal(reqByte, &environment)
	formulas := models.GetFormulas()

	formulas = append([]models.BjFormula{
		{FormulaKey: "singleIronWeight",
			Name:    "单个白铁重量",
			Formula: c.WeightFormula},
	}, formulas...)
	// 使用govaluate 对formula.Formula中对表达式进行计算
	// 使用的数据结构为c QuotationsV1Req, 评估之后，使用 c.SetByFieldName(formula.FormulaKey)
	// 设置c中对应字段的值
	// 遍历formulas并对每个公式进行评估
	for _, formula := range formulas {
		// 创建一个新的表达式
		expression, err := govaluate.NewEvaluableExpressionWithFunctions(
			formula.Formula, govaluateFunctions)
		if err != nil {
			s.Log.Errorf("解析表达式失败: %v\n", err)
			return errors.New("解析表达式失败")
		}

		// 评估表达式
		result, err := expression.Evaluate(environment)
		if err != nil {
			errMsg := fmt.Sprintf("评估表达式 %s 失败: %v", formula.Name, err)
			s.Log.Errorf(errMsg)
			return errors.New(errMsg)
		}
		s.Log.Debugf("%s 计算结果: %v", formula.Name, result)
		environment[formula.FormulaKey] = result

	}
	newByte, _ := json.Marshal(environment)
	err := json.Unmarshal(newByte, c)
	if err != nil {
		s.Log.Errorf("生成报价失败: %v\n", err)
		return errors.New("生成报价失败")
	}

	return nil

}
