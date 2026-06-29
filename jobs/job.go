package jobs

import (
	"encoding/json"
	"fmt"

	rdb "github.com/beyond3800/hawk/redis"
)

func Dispatch(queue string, payload interface{}) error{
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Unable to marsh payload")
	}
	redis := rdb.Rdb
	ctx := rdb.Ctx
	return redis.LPush(ctx,queue,data).Err()
}

func DispatchWithTime() {

}
func worker(queue string) error{
	redis := rdb.Rdb
	ctx := rdb.Ctx
	for{
		result,err := redis.BLPop(ctx,0,queue).Result()
		if err != nil {
			return err
		}
		fmt.Println(result)
	}
}

func StartWorker(workerNos int, queue string) error{
	for i := 0; i < workerNos; i++ {
		go worker(queue)
	}
	return nil
}