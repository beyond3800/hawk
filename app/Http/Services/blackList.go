package Services

import (
	"fmt"
	"time"

	"github.com/beyond3800/hawk/config"
	"github.com/redis/go-redis/v9"
)


func BlacklistToken(rdb *redis.Client, token string, ttl time.Duration) error{
	return rdb.Set(config.Ctx,"blacklist:"+token, token, ttl).Err()
}


func IsBlackListed(rdb *redis.Client, token string) (bool, error){
	_, err := rdb.Get(config.Ctx,"blacklist:"+token).Result()
	// fmt.Println(val)
	if err == redis.Nil{
		fmt.Println("Not blacklisted")
		return false, nil
	}
	if err != nil{
		fmt.Println("Redis error:", err)
		return false, err
	}
		// fmt.Println("blacklisted")
	return true, nil
	
}