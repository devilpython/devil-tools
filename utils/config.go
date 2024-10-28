package utils

import (
	"log"
	"os"
)

//获得配置结构体实例
func GetConfigMap(path string) *map[string]interface{} {
	//配置对象
	var dataMap *map[string]interface{}

	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Open file failed [Err:%s]", err.Error())
		return nil
	}
	// 关闭文件
	defer func(){
		_ = file.Close()
	}()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	data := make(map[string]interface{})
	dataMap = &data
	err = decoder.Decode(&dataMap)
	if err != nil {
		log.Println("config parse error:", err)
	}
	return dataMap
}

//获得配置结构体实例
func GetConfigMapArray(path string) []map[string]interface{} {
	//配置对象
	var dataMapArray []map[string]interface{}

	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Open file failed [Err:%s]", err.Error())
		return nil
	}
	// 关闭文件
	defer func(){
		_ = file.Close()
	}()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dataMapArray)
	if err != nil {
		log.Println("config parse error:", err)
	}
	return dataMapArray
}
