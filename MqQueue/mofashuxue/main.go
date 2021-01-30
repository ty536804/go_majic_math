package main

import (
	"elearn100/Model/Message"
	"elearn100/MqQueue/Mq"
)

func main() {
	Mq.ConsumeEx("mofashuxue", "fanout", "", Message.AddMessage)
}
