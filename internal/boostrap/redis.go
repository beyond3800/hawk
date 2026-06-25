package bootstrap

import (
	"fmt"
	"log"
	"os"
	"strconv"

	rdb "github.com/beyond3800/hawk/core/redis"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb* redis.Client
)
func ConnectRedis() error { 
	if err := godotenv.Load(); err != nil{
		return fmt.Errorf("Unable to load the neccessery file")
	}
	
	dbInt, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil{
		return fmt.Errorf("A number is needed not a string")
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
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