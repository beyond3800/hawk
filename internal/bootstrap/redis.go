package bootstrap

import (
	"fmt"
	"log"
	"strconv"

	"github.com/beyond3800/hawk/internal/env"
	rdb "github.com/beyond3800/hawk/redis"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb* redis.Client
)
func ConnectRedis() error { 
	
	redisDb,_:= env.Get("REDIS_DB")
	dbInt, err := strconv.Atoi(redisDb)
	redisAddr,_ := env.Get("REDIS_ADDR")
	redisPwd,_ := env.Get("REDIS_PASSWORD")
	if err != nil{
		return fmt.Errorf("A number is needed not a string")
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		Password: redisPwd,
		DB: dbInt,
	})
	pong, err := Rdb.Ping(rdb.Ctx).Result()

	if err != nil{
		log.Println("Redis is not working")
	}
	rdb.SetRedis(Rdb)
	fmt.Println(pong)
	return nil
}