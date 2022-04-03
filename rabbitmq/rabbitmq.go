package rabbitmq

import (
	"email/email"
	"email/structs"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)
const RABBITMQ_DEFAULT_USER = "impanda"
const RABBITMQ_DEFAULT_PASS =  "I8UMY6wNaevm1MQ3QzfxjilfnBeXfb22NpbDDqxHCla6bBFxuRLsb5t8myXSYtmZ"

// failOnErrStudent 检查异常并终断程序输出错误
func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
var conn *amqp.Connection

func InitRabbitmq() {
	for  i := 0; i < 100; i++ {
		time.Sleep(time.Second * 1)
		var err error
		url := fmt.Sprintf("amqp://%s:%s@127.0.0.1:5672/", RABBITMQ_DEFAULT_USER, RABBITMQ_DEFAULT_PASS)
		conn, err = amqp.Dial(url)
		if err != nil {
			log.Println("无法链接rabbitmq服务:" + err.Error())
			continue
		}
		log.Println("链接rabbitmq服务成功")
		return
	}
	panic("经过了100次尝试链接rabbitmq都失败了")
}

func Consume() {
	ch, err := conn.Channel()
	failOnErr(err, "无法打开信道")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"activation_info",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "无法声明队列")
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "注册消费者失败")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var a structs.ActivationEx
			err = json.Unmarshal(d.Body, &a)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = email.SendEmail(a)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}()
	<-forever
}


