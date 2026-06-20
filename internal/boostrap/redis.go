package bootstrap



import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	Rdb* redis.Client
)
func ConnectRedis() *redis.Client { 
	if err := godotenv.Load(); err != nil{
		log.Fatal("Unable to load the neccessery file")
	}
	
	dbInt, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil{
		log.Fatal("A number is needed not a string")
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: dbInt,
	})
	pong, err := Rdb.Ping(Ctx).Result()

	if err != nil{
		log.Println("Redis is not working")
		return Rdb
	}
	fmt.Println(pong)
	return Rdb
}