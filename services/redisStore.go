package services

import (
	"encoding/json"
	"time"

	"github.com/beyond3800/hawk/config"
)

func RedisSet(name string, data any, ttl time.Duration) error{
	return config.Rdb.Set(config.Ctx, name, data, ttl).Err()
}
func RedisGet(name string) (string, error){
	return config.Rdb.Get(config.Ctx,name).Result()
}
func RedisExists(key string)(bool,error){
	count,err := config.Rdb.Exists(config.Ctx, key).Result()
	return count == 1, err
}
func RedisDelete(key string) error{
	return config.Rdb.Del(config.Ctx, key).Err()
}
func RedisSetJSON(key string, data any, ttl time.Duration) error{
	b,_ := json.Marshal(data)
	return  RedisSet(key,b,ttl)
}

func RedisGetJSON(key string, data any) error{
	val, err := RedisGet(key)
	if err != nil{
		return err
	}
	return json.Unmarshal([]byte(val),data)
}
