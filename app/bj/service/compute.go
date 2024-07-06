package service

import (
	"encoding/json"
	"errors"
	"go-admin/app/bj/models"
	"go-admin/app/bj/service/dto"

	"github.com/Knetic/govaluate"
)

func (s *QuotationsV1) Compute(c *dto.QuotationsV1Req) error {

	reqByte, _ := json.Marshal(c)
	environment := make(map[string]interface{}, 0)
	json.Unmarshal(reqByte, &environment)
	formulas := models.GetFormulas()

	formulas = append([]models.BjFormula{
		{FormulaKey: "SingleIronWeight",
			Formula: c.WeightFormula},
	}, formulas...)
	// 使用govaluate 对formula.Formula中对表达式进行计算
	// 使用的数据结构为c QuotationsV1Req, 评估之后，使用 c.SetByFieldName(formula.FormulaKey)
	// 设置c中对应字段的值
	// 遍历formulas并对每个公式进行评估
	for _, formula := range formulas {
		// 创建一个新的表达式
		expression, err := govaluate.NewEvaluableExpression(formula.Formula)
		if err != nil {
			s.Log.Errorf("解析表达式失败: %v\n", err)
			return errors.New("解析表达式失败")
		}

		// 评估表达式
		result, err := expression.Evaluate(environment)
		if err != nil {
			s.Log.Errorf("评估表达式 %s 失败: %v", formula.Name, err)
			return errors.New("评估表达式失败")
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
