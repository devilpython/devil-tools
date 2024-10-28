package config

import (
	"devil-tools/utils"
	"log"
	"sync"
)

//配置结构体
type Configuration struct {
	RedisType      			string //redis类型: single(单台redis服务器) sentinel(哨兵集群)
	RedisSentinelMasterName string //redis哨兵集群主服务器名
	RedisSentinelList    	string //redis哨兵服务器列表
	RedisServer      		string //redis服务器
	RedisPort        		int16  //redis端口号
	RedisPassword    		string //redis密码
	LogPath          		string //日志路径
	RegisterKey	     		string //注册码
}

//配置对象
var config Configuration
var once sync.Once
var configOk = false

//获得配置结构体实例
func GetConfigInstance() (Configuration, bool) {
	once.Do(func() {
		dataMap := utils.GetConfigMap("config.json")
		if dataMap != nil {
			_config, err := utils.ConvertMapToStruct(*dataMap, Configuration{})
			if err != nil {
				log.Println("config error:", err)
			} else {
				c, ok := _config.(*Configuration)
				if ok {
					config = *c
					configOk = true
				}
			}
		}
	})
	return config, configOk
}
