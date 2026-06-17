package controllers

import (
	// "github.com/beyond3800/hawk/config"
	// "github.com/beyond3800/hawk/jobs"
	"fmt"
	"time"

	"github.com/beyond3800/hawk/jobs"
	"github.com/gin-gonic/gin"
)

type TecterController struct {}
func (Tecter * TecterController) Index(c *gin.Context) {
	// TODO: implement index
}
func (Tecter * TecterController) SendEmail(c *gin.Context) {
	// TODO: implement index
	// var message = jobs.Message{
	// 	Subject: "Payment of salaries",
	// 	Body: "Pay my salary oloriburoku",
	// 	To: "foolishboss@gmail.com",
	// }
	// ctx := config.Ctx
	// jobs.Dispatch(ctx,message)
	// jobs.StartWorker(3,ctx)
	type Message struct{
		Subject  string `json:"subject"`
		Body     string `json:"body"`
		To      string   `json:"to"`
	}
	type Email struct{
		ID       string `json:"id"`
		Msg      Message `json:"msg"`
		To       string `json:"to"`
		From     string `json:"from"`
		SentAt   time.Time `json:"sentAt"`
	}
	message := Message{
		Subject: "SignUp to beyond",
		Body: "Welcome on Board",
		To: "Folly123",
	}
	email := Email{
		ID:     fmt.Sprintf("%d", time.Now().UnixNano()),
		Msg:    message,
		To:     message.To,
		From:   "noreply@example.com",
		SentAt: time.Now(),
	}
	for i := 0; i < 6; i++ {
		if err := jobs.Dispatch("email",email); err != nil{
			fmt.Println(err)
		}
	}
	jobs.StartWorker(3,"email")
}


func (Tecter * TecterController) Store(c *gin.Context) {
	// TODO: implement store
}

func (Tecter * TecterController) Show(c *gin.Context) {
	// TODO: implement show
}

func (Tecter * TecterController) Update(c *gin.Context) {
	// TODO: implement update
}

func (Tecter * TecterController) Destroy(c *gin.Context) {
	// TODO: implement destroy
}