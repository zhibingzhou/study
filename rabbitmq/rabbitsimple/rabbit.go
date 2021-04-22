package rabbitsimple

import (
	"bytes"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://zhou:123@127.0.0.1:5672/zhou" //amqp:// 账号 : 密码 @ 连接生产者、消费者的端口 /verhost name

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	//连接信息
	Mqurl string
}

func NewRabbitMQ(QueueName, Exchange, key string) *RabbitMQ {
	MQURL := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", Sys.RabbitMq.Admin, Sys.RabbitMq.Pwd, Sys.RabbitMq.Ip, Sys.RabbitMq.Port, Sys.RabbitMq.Verhost)
	fmt.Println(MQURL)
	rabbitmq := RabbitMQ{QueueName: QueueName, Exchange: Exchange, Key: key, Mqurl: MQURL}
	var err error
	//先连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误")
	//再获取通道
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel错误")
	return &rabbitmq
}

//中断连接
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		//错误日志打印
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//群发 Direct: default 未指定 exchange 和 bangdinkey ，可以被多个消费者消费
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

//发送者
func (r *RabbitMQ) PulishSimple(message string) {
	//1.申请队列，如果不存在队列，则会自动创建
	//name string,  队列名称
	//durable,  持久化
	//autoDelete, 自动删除，最后一个消费者，断开连接
	//exclusive, 是否有排他性，一般不用，不许其它用户使用
	//noWait bool, 是否阻塞 ， 设置false ，等待服务器响应
	// args Table
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
	//发送
	// exchange,
	// key string,
	// mandatory,
	// immediate,
	//  msg,
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, //mandatory 如果为true , 根据exchange 类型和routkey 规则， 如果无法找到，发返回消息给发送者
		false, // immediate 如果为true, exchange 发送的时候，发现没绑定消费者，返回信息给发送者
		amqp.Publishing{
			ContentType: "text/pain",
			Body:        []byte(message),
		},
	)

	fmt.Println("发送成功")
}

//消费者
func (r *RabbitMQ) ConsumeSimple() {
	//也是申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	// queue
	// consumer
	// autoAck,
	// exclusive,
	//  noLocal,
	//  noWait
	// args Table
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    // consumer 用来区分消费者
		true,  // autoAck 是否自动应答
		false, // exclusive 排他性
		false, //noLocal , 如果为true ,表示不能将同一个connection中发送的消息传给这个connection中的消费者
		false, //noWait false 为阻塞
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Receive messsage %s", *BytesToString(&d.Body))
			fmt.Println()
		}
	}()

	log.Printf("Wait For message")
	<-forever
	log.Printf("End")
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
