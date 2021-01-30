package Mq

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
)

type CallBack func(msg string)

// @Summer Mq连接
func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	return conn, err
}

// @Title PublishEx
// @DDescription 订阅模式，生产端
// @Param exchange  	string	交换机名称
// @Param types			string	类型
// @Param routingKey	string	路由key
// @Param body			string	内容
func PublishEx(exchange, types, routingKey, body string) error {
	//建立连接
	conn, err := Connect()
	if err != nil {
		fmt.Println("建立连接失败:", err)
		return err
	}
	defer conn.Close()

	//创建一个通道
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("创建通道失败:", err)
		return err
	}
	defer channel.Close()

	//创建一个交互机
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println("创建交互机失败:", err)
		return err
	}
	err = channel.Publish(exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	return err
}

// @Title ConsumeEx 消费端
// @Description 订阅模式，生产端
// @Param exchange  	string	交换机名称
// @Param types			string	类型
// @Param routingKey	string	路由key
// @Param callback		fun		回调函数
func ConsumeEx(exchange, types, routingKey string, callback CallBack) {
	//建立连接
	conn, err := Connect()
	if err != nil {
		fmt.Println("消费端 建立连接失败:", err)
		return
	}
	defer conn.Close()

	//建立通道
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("建立通道失败:", err)
		return
	}
	defer channel.Close()

	//创建交换机
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println("创建交互机失败：", err)
		return
	}
	//创建队列  临时队列
	q, err := channel.QueueDeclare("",
		false,
		false,
		true,
		false,
		nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = channel.QueueBind(q.Name, routingKey, exchange, false, nil)
	if err != nil {
		fmt.Println("绑定失败")
		return
	}

	//接收信息
	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("没有接收到信息:", err)
		return
	}
	forever := make(chan bool)
	go func() {
		for {
			for d := range msgs {
				s := BytesToString(&(d.Body))
				callback(*s)
				d.Ack(false)
			}
		}
	}()
	<-forever
}

// @Title BytesToString
// @Description 字节转字符串
// @Param b  	*[]byte	字节
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

// @Title 发送端函数
// @Param exchange	string	交换机名称
// @Param queueName	string 	队列名称
// @Param body		string	内容
func Publish(exchange, queueName, body string) error {
	//建立连接
	conn, err := Connect()
	if err != nil {
		fmt.Println("建立连接失败:", err)
		return err
	}
	defer conn.Close()

	//创建一个通道 channel
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("创建一个通道失败:", err)
		return err
	}
	defer channel.Close()

	//创建队列
	q, err := channel.QueueDeclare(
		queueName,
		true, //是否持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil
	}
	//发送消息
	err = channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain", //发布内容类型
		Body:         []byte(body),
	})
	return err
}

//接收者的方法
func Consumer(exchange, queueName string, callback CallBack) {
	//建立连接
	conn, err := Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//创建一个通道 channel
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer channel.Close()

	//创建队列
	q, err := channel.QueueDeclare(
		queueName,
		true, //是否持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	//接收消息 第三参数为true 表示自动应答
	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			d.Ack(false) //手动应答
		}
	}()
	fmt.Println("waiting for messages")
	<-forever
}
