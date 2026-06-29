package Services

import (
	"encoding/json"
	"time"

	rdb "github.com/beyond3800/hawk/redis"
)

func RedisSet(name string, data any, ttl time.Duration) error{
	return rdb.Rdb.Set(rdb.Ctx, name, data, ttl).Err()
}
func RedisGet(name string) (string, error){
	return rdb.Rdb.Get(rdb.Ctx,name).Result()
}
func RedisExists(key string)(bool,error){
	count,err := rdb.Rdb.Exists(rdb.Ctx, key).Result()
	return count == 1, err
}
func RedisDelete(key string) error{
	return rdb.Rdb.Del(rdb.Ctx, key).Err()
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
