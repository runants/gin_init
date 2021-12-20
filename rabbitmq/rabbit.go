package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 连接RabbitMQ服务器
	conn, err := amqp.Dial("amqp://admin:egoo3466@192.168.25.181:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 创建一个channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	//q, err  := ch.QueueDeclare(
	//	"qms_inbound_cases",			// 队列名称
	//	false,			// 是否持久化
	//	false,		// 是否自动删除
	//	false,			// 是否独立
	//	false,nil,
	//)

	f, err := os.Open("U:\\std\\GOLANG_DEMO\\gin_demo\\rabbitmq\\file.txt")
	if err != nil {
		log.Println("ERROR: {}",err)
	}
	ioutil.ReadAll(f)

	failOnError(err, "Failed to declare a queue")
	// 发送消息到队列中
	body := "Hello World!"
	err = ch.Publish(
		"",                  // exchange
		"qms_inbound_cases", // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	fmt.Println("send message success")
}

// 帮助函数检测每一个amqp调用
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
