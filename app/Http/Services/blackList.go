package Services

import (
	"fmt"
	"time"

	rdb "github.com/beyond3800/hawk/redis"
	"github.com/redis/go-redis/v9"
)


func BlacklistToken(rc *redis.Client, token string, ttl time.Duration) error{
	return rc.Set(rdb.Ctx,"blacklist:"+token, token, ttl).Err()
}


func IsBlackListed(rc *redis.Client, token string) (bool, error){
	_, err := rc.Get(rdb.Ctx,"blacklist:"+token).Result()
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