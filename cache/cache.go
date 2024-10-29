package cache

import (
	"fmt"
	"log"

	"github.com/devilpython/devil-tools/redis_utils"
	"github.com/devilpython/devil-tools/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
)

// 设置数据
func Set(key string, data interface{}) bool {
	conn := redis_utils.NewRedisInstance()
	if conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		_, err := conn.Do("SET", key, data)
		if err != nil {
			log.Println(fmt.Sprintf("redis set 【%s】 failed:", key), err)
		} else {
			return true
		}
	}
	return false
}

// 设置对象数据
func SetObject(key string, data interface{}) bool {
	conn := redis_utils.NewRedisInstance()
	if conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		jsonData, err := utils.ConvertDataToJson(data)
		if err == nil {
			_, err := conn.Do("SET", key, jsonData)
			if err != nil {
				log.Println(fmt.Sprintf("redis set 【%s】 failed:", key), err)
			} else {
				return true
			}
		}
	}
	return false
}

// 带时限设置数据
func SetEx(key string, data interface{}, second int64) bool {
	conn := redis_utils.NewRedisInstance()
	if conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		_, err := conn.Do("SET", key, data, "EX", fmt.Sprintf("%d", second))
		if err != nil {
			log.Println(fmt.Sprintf("redis set 【%s】 failed:", key), err)
		} else {
			return true
		}
	}
	return false
}

// 限制时间设置对象数据
func SetObjectEx(key string, data interface{}, second int64) bool {
	conn := redis_utils.NewRedisInstance()
	if conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		jsonData, err := utils.ConvertDataToJson(data)
		if err == nil {
			_, err := conn.Do("SET", key, jsonData, "EX", fmt.Sprintf("%d", second))
			if err != nil {
				log.Println(fmt.Sprintf("redis set 【%s】 failed:", key), err)
			} else {
				return true
			}
		}
	}
	return false
}

// 获得数据
func Get(key string) (interface{}, bool) {
	conn := redis_utils.NewRedisInstance()
	if conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		msg, err := conn.Do("GET", key)
		if err != nil {
			log.Println(fmt.Sprintf("redis_utils get data[%s] failed:", key), err)
		} else {
			return msg, true
		}
		//c.Do("SELECT", REDIS_DB)
	}
	return "", false
}

// 获得数据
func GetObject(key string, objDataPointer interface{}) error {
	conn := redis_utils.NewRedisInstance()
	if conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		msg, err := redis.String(conn.Do("GET", key))
		if err != nil {
			return err
		} else {
			err = utils.ConvertJsonToData(msg, objDataPointer)
			return err
		}
	}
	return errors.New("redis连接失败")
}

// 移除数据
func Remove(key string) {
	conn := redis_utils.NewRedisInstance()
	if conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		_, _ = conn.Do("DEL", key)
	}
}
