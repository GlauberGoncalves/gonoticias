package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	redisExpire = 60
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	HandleError(err)
	return c
}

func Set(key string, value []byte) error {

	conn := RedisConnect()
	defer conn.Close()

	_, err := conn.Do("SET", key, []byte(value))
	HandleError(err)

	conn.Do("EXPIRE", key, redisExpire)

	return err
}

func Get(key string) ([]byte, error) {

	conn := RedisConnect()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func Flush(key string) ([]byte, error) {

	conn := RedisConnect()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("DEL", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
