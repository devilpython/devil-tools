package utils

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"log"
	"reflect"
	"strings"
	"time"
)

//转换映射表到结构体
func ConvertMapToStruct(mapData interface{}, structData interface{}) (interface{}, error) {
	ConvertToCamelCaseName(mapData)
	switch data := mapData.(type) {
	case map[string]interface{}:
		return mapToStruct(data, structData)
	case *map[string]interface{}:
		return mapToStruct(data, structData)
	case []map[string]interface{}:
		var dataArray []interface{}
		for index := range data {
			newData, err := mapToStruct(data[index], structData)
			if err != nil {
				return nil, err
			} else {
				dataArray = append(dataArray, newData)
			}
		}
		return dataArray, nil
	case []*map[string]interface{}:
		var dataArray []interface{}
		for index := range data {
			newData, err := mapToStruct(*data[index], structData)
			if err != nil {
				return nil, err
			} else {
				dataArray = append(dataArray, newData)
			}
		}
		return dataArray, nil
	}
	return nil, errors.New("映射表参数错误")
}

//结构体转换成"snake_case"键值的哈希表
func ConvertStructToSnakeCaseMap(structDataPointer interface{}, mapDataPointer interface{}) error {
	jsonData, err := ConvertDataToJson(structDataPointer)
	if err == nil {
		err = ConvertJsonToData(jsonData, mapDataPointer)
		if err == nil {
			ConvertToSnakeCaseName(mapDataPointer)
		}
	}
	return err
}

//结构体转换成"camelCase"键值的哈希表
func ConvertStructToCamelCaseMap(structDataPointer interface{}, mapDataPointer interface{}) error {
	jsonData, err := ConvertDataToJson(structDataPointer)
	if err == nil {
		err = ConvertJsonToData(jsonData, mapDataPointer)
		if err == nil {
			ConvertToCamelCaseName(mapDataPointer)
		}
	}
	return err
}

//转换json数据
func ConvertDataToJson(data interface{}) (string, error) {
	//ConvertToSnakeCaseName(data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	} else {
		return string(jsonData), nil
	}
}

//转换json数据
func ConvertJsonToData(jsonData string, dataPointer interface{}) error {
	err := json.Unmarshal([]byte(jsonData), dataPointer)
	if err != nil {
		log.Println("json err:", err)
		return err
	} else {
		return nil
	}
}

//转换到"camelCase"样式的键值
func ConvertToCamelCaseName(mapData interface{}) {
	switch data := mapData.(type) {
	case map[string]interface{}:
		convertMapName(data)
	case *map[string]interface{}:
		convertMapName(*data)
	case []map[string]interface{}:
		for index := range data {
			convertMapName(data[index])
		}
	case []*map[string]interface{}:
		for index := range data {
			convertMapName(*data[index])
		}
	case []interface{}:
		for index := range data {
			mapData, ok := data[index].(map[string]interface{})
			if ok {
				convertMapName(mapData)
			}
		}
	case *[]interface{}:
		for index := range *data {
			mapData, ok := (*data)[index].(map[string]interface{})
			if ok {
				convertMapName(mapData)
			}
		}
	}
}

//转换到"snake_case"样式名称
func ConvertToSnakeCaseName(mapData interface{}) {
	switch data := mapData.(type) {
	case map[string]interface{}:
		convertSnakeName(data)
	case *map[string]interface{}:
		convertSnakeName(*data)
	case []map[string]interface{}:
		for index := range data {
			convertSnakeName(data[index])
		}
	case []*map[string]interface{}:
		for index := range data {
			convertSnakeName(*data[index])
		}
	case []interface{}:
		for index := range data {
			mapData, ok := data[index].(map[string]interface{})
			if ok {
				convertSnakeName(mapData)
			}
		}
	case *[]interface{}:
		for index := range *data {
			mapData, ok := (*data)[index].(map[string]interface{})
			if ok {
				convertSnakeName(mapData)
			}
		}
	}
}

//转换时间字符串到时间戳
func ConvertTimeToTimestamp(value string) int64 {
	timeArray := strings.Split(value, "+")
	value = strings.Replace(timeArray[0], "T", " ", 1)
	stamp, _ := time.ParseInLocation(timeTemplate, value, time.Local)
	return stamp.Unix() * 1000
}

//时间模板
var timeTemplate = "2006-01-02 15:04:05"

//转换到映射表名字【"dataMap"的样式】
func convertMapName(mapData map[string]interface{}) {
	for key, value := range mapData {
		delete(mapData, key)
		key = ToHumpCase(key)
		mapData[key] = value
		switch data := value.(type) {
		case map[string]interface{}:
			convertMapName(data)
		case *map[string]interface{}:
			convertMapName(*data)
		case []interface{}:
			for index := range data {
				arrayData, ok := data[index].(map[string]interface{})
				if ok {
					convertMapName(arrayData)
				}
			}
		case []*interface{}:
			for index := range data {
				arrayData, ok := (*data[index]).(*map[string]interface{})
				if ok {
					convertMapName(*arrayData)
				}
			}
		}
	}
}

//转换到Json名字【"snake_case"的样式】
func convertSnakeName(mapData map[string]interface{}) {
	for key, value := range mapData {
		delete(mapData, key)
		key = ToSnakeCase(key)
		mapData[key] = value
		switch data := value.(type) {
		case map[string]interface{}:
			convertSnakeName(data)
		case *map[string]interface{}:
			convertSnakeName(*data)
		case []interface{}:
			for index := range data {
				arrayData, ok := data[index].(map[string]interface{})
				if ok {
					convertSnakeName(arrayData)
				}
			}
		case []*interface{}:
			for index := range data {
				arrayData, ok := (*data[index]).(map[string]interface{})
				if ok {
					convertSnakeName(arrayData)
				}
			}
		}
	}
}

//反射创建新对象。
func newObject(target interface{}) (interface{}, error) {
	if target == nil {
		return nil, errors.New("参数不能未空")
	}

	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr { //指针类型获取真正type需要调用Elem
		t = t.Elem()
	}

	return reflect.New(t).Interface(), nil
}

//映射表转结构体
func mapToStruct(mapData interface{}, structData interface{}) (interface{}, error) {
	newData, err := newObject(structData)
	if err == nil {
		err = mapstructure.WeakDecode(mapData, &newData)
	}
	return newData, err
}
