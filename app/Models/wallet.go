package Models

import (
	"context"
)

type Wallet struct {
	ID string
	Balance int
	Type string
}



func (w *Wallet) Deposit(ctx context.Context, amount string) error{


	return nil

}



// local processed = redis.call("HEXISTS", KEYS[2], ARGV[1])
// if processed == 1 then
//     return redis.call("HGET", KEYS[2], ARGV[1])
// end

// local newbal = redis.call("INCRBY", KEYS[1], ARGV[2])

// redis.call("HSET", KEYS[2], ARGV[1], newbal)
// redis.call("RPUSH", KEYS[3], ARGV[3])

// return newbal
// the issue is the lua I don't fully understand it