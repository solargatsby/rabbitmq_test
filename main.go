package main

import "github.com/Taskon-xyz/rabbitMQ_test/rabbit_mq/hello_world"

func main() {
	go hello_world.Send()
	hello_world.Receive()
}
