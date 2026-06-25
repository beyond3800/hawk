package rdb

import (
	"context"

	"github.com/redis/go-redis/v9"
)


var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func SetRedis( rdb *redis.Client)*redis.Client {
	Rdb = rdb
	return rdb
}

