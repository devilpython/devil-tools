package redis_utils

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"devil-tools/config"
)

func NewRedisInstance() redis.Conn {
	conf, _ := config.GetConfigInstance()
	conn, err := getRedisConnection()
	if err != nil {
		log.Println("Connect to redis_utils error", err)
	} else if _, err := conn.Do("AUTH", conf.RedisPassword); err != nil {
		log.Println("密码错误")
		conn.Close()
	} else {
		return conn
	}
	return nil
}

//获得redis连接
func getRedisConnection() (redis.Conn, error) {
	conf, _ := config.GetConfigInstance()
	server := conf.RedisServer + ":" + strconv.Itoa(int(conf.RedisPort))
	conn, err := redis.Dial("tcp", server)
	return conn, err
}
