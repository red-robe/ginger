package model

import (
	"errors"
	"fmt"
	"github.com/gofuncchan/ginger/common"
)

var filterCondition = []string{"=", ">", "<", "=", "<=", ">=", "!=", "<>", "in", "not in", "like", "not like", "between", "not between"}

// 转换 filter过滤字段成gendry要求的格式
func ConvertFilterToMap(fieldSliceSlice [][]interface{}) (map[string]interface{}, error) {
	filterMap := make(map[string]interface{})
	for _, fieldSlice := range fieldSliceSlice {
		fieldName, ok := fieldSlice[0].(string)
		if !ok {
			return nil, errors.New("filter field slice first element must be string")
		}

		l := len(fieldSlice)
		if l < 2 {
			return nil, errors.New("filter field slice length must be great 1")
		}
		switch l {
		case 2:
			filterMap[fieldName] = fieldSlice[1]
		case 3:
			isExsitCondition := common.IsExsitItem(fieldSlice[1], filterCondition)
			if !isExsitCondition {
				return nil, errors.New("filter field slice second element should be one of:" + fmt.Sprintf("%s", filterCondition))
			}
			mapKey := fieldSlice[0].(string) + " " + fieldSlice[1].(string)
			filterMap[mapKey] = fieldSlice[2]
		default:
			isExsitCondition := common.IsExsitItem(fieldSlice[1], []string{"in", "not in"})
			if !isExsitCondition {
				return nil, errors.New("filter field slice length great 3,the  second element should be one of ['in', 'not in']" )
			}
			mapKey := fieldSlice[0].(string) + " " + fieldSlice[1].(string)
			filterMap[mapKey] = fieldSlice[2:]

		}
	}

	return filterMap, nil
}
