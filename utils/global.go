package utils

import (
	"github.com/devilpython/devil-tools/goroutine_local"
	"github.com/emirpasic/gods/maps/treemap"
)

// 获得全局数据
func GetGlobalData(key int) (interface{}, bool) {
	gl := goroutine_local.GetGoroutineLocal()
	dataMapObj := gl.Get()
	if dataMapObj != nil {
		globalMap, ok := dataMapObj.(*treemap.Map)
		if ok {
			return globalMap.Get(key)
		}
	}
	return nil, false
}

// 设置全局数据
func SetGlobalData(key int, value interface{}) {
	gl := goroutine_local.GetGoroutineLocal()
	dataMapObj := gl.Get()
	if dataMapObj != nil {
		globalMap, ok := dataMapObj.(*treemap.Map)
		if ok {
			globalMap.Put(key, value)
		}
	} else {
		globalMap := treemap.NewWithIntComparator()
		globalMap.Put(key, value)
		gl.Set(globalMap)
	}
}

// 设置全局数据
func RemoveGlobalData(key int) {
	gl := goroutine_local.GetGoroutineLocal()
	dataMapObj := gl.Get()
	if dataMapObj != nil {
		globalMap, ok := dataMapObj.(*treemap.Map)
		if ok {
			globalMap.Remove(key)
		}
	}
}

// 删除全部全局数据
func RemoveAllGlobalData() {
	gl := goroutine_local.GetGoroutineLocal()
	gl.Remove()
}

// 获得来自于映射表的字符串值
func GetStringFromMap(dataMap map[string]interface{}, key string) string {
	obj, ok := dataMap[key]
	if ok {
		strValue, strOk := obj.(string)
		if strOk {
			return strValue
		}
	}
	return ""
}
