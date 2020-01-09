package common

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func MapToStruct(mapData map[string]interface{},structData interface{}) error {
	return mapstructure.Decode(mapData, &structData)
}

func StructToMap(structData interface{}) map[string]interface{} {
	objType := reflect.TypeOf(structData)
	objValue := reflect.ValueOf(structData)

	var data = make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Name] = objValue.Field(i).Interface()
	}
	return data
}