package utils

import (
	"crypto/md5"
	"fmt"
	"os"
	"reflect"
	"strings"

	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 获得字段名数组
func GetFieldNameArray(data interface{}) [][]string {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}
	if dataType.Kind() != reflect.Struct {
		return nil
	}
	fieldNum := dataType.NumField()
	result := make([][]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		fieldName := dataType.Field(i).Name
		var tag string
		tag = ""
		tags := strings.Split(string(dataType.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tag = tags[1]
		}
		record := []string{fieldName, tag}
		result = append(result, record)
	}
	return result
}

// 绑定数据
func BindingData(jsonStr string, param interface{}) (interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	} else {
		_, ok := param.(map[string]interface{})
		if ok {
			//ConvertToCamelCaseName(data)
			return data, nil
		} else {
			return ConvertMapToStruct(data, param)
		}
	}
}

// 求向量相似度
func DoSimilarity(vec1 []float32, vec2 []float32) float32 {
	var sum float32 = 0
	for i := 0; i < len(vec1); i++ {
		sum += vec1[i] * vec2[i]
	}
	return sum
}

// 获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			files, dirs, _ = GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

// 获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			files, _ = GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for index := range temp {
			files = append(files, temp[index])
		}
	}

	return files, nil
}

// 生成md5码
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 创建ID
func CreateId() string {
	return Md5(uuid.NewV4().String())
}

// 为映射表数组过滤时间戳
func FilterTimestampForMapArray(resultArray []interface{}) {
	for index := range resultArray {
		mapData, isMap := resultArray[index].(map[string]interface{})
		if isMap {
			FilterTimestamp(mapData)
		}
	}
}

// 过滤时间戳
func FilterTimestamp(dataMap map[string]interface{}) {
	for key, value := range dataMap {
		if strings.Contains(key, "timestamp") {
			valueStr, ok := value.(string)
			if ok {
				dataMap[key] = ConvertTimeToTimestamp(valueStr)
			}
		}
	}
}

// 复制数据
func CopyData(srcObj, destObjPointer interface{}) {
	json, err := ConvertDataToJson(srcObj)
	if err == nil {
		//fmt.Println("...............json:", json)
		err = ConvertJsonToData(json, destObjPointer)
	} else {
		fmt.Println("............copy error:", err)
	}
}
