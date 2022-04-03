package main

import "email/rabbitmq"

func main() {
	rabbitmq.InitRabbitmq()
	rabbitmq.Consume()
}
