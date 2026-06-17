package jobs

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"time"

// 	"github.com/beyond3800/hawk/config"
// 	_"github.com/redis/go-redis/v9"
// )

// type Email struct{
// 	ID       string `json:"id"`
// 	Msg      Message `json:"msg"`
// 	To       string `json:"to"`
// 	From     string `json:"from"`
// 	SentAt   time.Time `json:"sentAt"`
// }
// type Message struct{
// 	Subject  string `json:"subject"`
// 	Body     string `json:"body"`
// 	To       string  `json:"to"`
// }

// func Dispatch(ctx context.Context , message Message)  {
// 	redis := config.Rdb
	
// 	email := Email{
// 		ID:     fmt.Sprintf("%d", time.Now().UnixNano()),
// 		Msg:    message,
// 		To:     message.To,
// 		From:   "noreply@example.com",
// 		SentAt: time.Now(),
// 	}
// 	data, err := json.Marshal(email)
// 	if err != nil {
// 		fmt.Println("Unable to marshal data")
// 	}
// 	if redis == nil {
// 		panic("REDIS IS NIL")
// 	}
// 	err = redis.LPush(ctx,"email", data).Err()
// 	if err != nil {
// 		fmt.Println("Unable to dispatch job")
// 	}
// 	fmt.Println("job dispatched succesfully")
// }

// func StartWorker(workerAmount int, ctx context.Context){
// 	for i := 0; i < workerAmount; i++ {
// 		go Worker(ctx)
// 	}
// }
// func Worker(ctx context.Context){
// 	redis := config.Rdb
// 	for{
// 		data,err := redis.BLPop(ctx,0,"email").Result()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		var email Email
// 		err = json.Unmarshal([]byte(data[1]),&email)
// 		if err != nil{
// 			fmt.Println(err)
// 		}
// 		SendEmail(email)
// 	}
// }
// func SendEmail(email Email) {
// 	fmt.Println(email.ID +"sent to " + email.To)
// }