package rabbitmq

import (
	"log"
	"strconv"

	"github.com/streadway/amqp"

	"github.com/gauravbansal74/mlserver/pkg/logger"
)

// Server stores the hostname and port number
type Configuration struct {
	ListenHost string
	ListenPort int
	Username   string
	Password   string
	Exchange   string
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Deliveries <-chan amqp.Delivery
}

var (
	conf *Configuration
)

func GetDeliveryChannel() (*RabbitMQ, error) {
	rmq, err := Dail()
	if err != nil {
		return rmq, err
	}
	rmq.Deliveries, err = rmq.Channel.Consume(
		conf.Exchange,
		"",
		false, //autoAck
		false, // exclusive,
		false, // noLocal,
		false, //noWait,
		nil,
	)
	if err != nil {
		return rmq, err
	}
	return rmq, nil
}

// Close Connection
func (r *RabbitMQ) Close() error {
	if r.Channel == nil {
		if r.Connection == nil {
			return nil
		}
	}
	return r.Connection.Close()
}

func Init(r *Configuration) {
	connection, err := amqp.Dial("amqp://" + r.Username + ":" + r.Password + "@" + r.ListenHost + ":" + strconv.Itoa(r.ListenPort) + "/")
	if err != nil {
		log.Fatal(err.Error(), "Error while connecting with RabbitMQ")
	}
	defer connection.Close()
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err.Error(), "Error while creating channel with RabbitMQ")
	}
	defer channel.Close()
	err = channel.ExchangeDeclare(
		r.Exchange, // name of the exchange
		"fanout",   // type
		true,       // durable
		false,      // delete when complete
		false,      // internal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		log.Fatal(err.Error(), "Error while declaring Exchange RabbitMQ")
	}
	_, err = channel.QueueDeclare(
		r.Exchange, // name of the queue
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // noWait
		nil,        // arguments
	)

	if err != nil {
		log.Fatal(err.Error(), "Error while declaring Queue RabbitMQ")
	}

	err = channel.QueueBind(r.Exchange, "", r.Exchange, false, nil)
	if err != nil {
		log.Fatal(err.Error(), "Error while declaring Queue RabbitMQ")
	}

	conf = &Configuration{
		ListenHost: r.ListenHost,
		ListenPort: r.ListenPort,
		Username:   r.Username,
		Password:   r.Password,
		Exchange:   r.Exchange,
	}
	logger.Info("RabbitMQ connection tested successfully.")
}

func ReadConfig() *Configuration {
	return conf
}

// Dail Connection
func Dail() (*RabbitMQ, error) {
	rabbitMQ := &RabbitMQ{}
	connection, err := amqp.Dial("amqp://" + conf.Username + ":" + conf.Password + "@" + conf.ListenHost + ":" + strconv.Itoa(conf.ListenPort) + "/")
	if err != nil {
		logger.Error(err, "Error while connecting rabbitmq")
		return rabbitMQ, err
	}

	channel, err := connection.Channel()
	if err != nil {
		connection.Close()
		logger.Error(err, "Error while creating rabbitmq channel")
		return rabbitMQ, err
	}
	rabbitMQ = &RabbitMQ{
		Connection: connection,
		Channel:    channel,
	}
	return rabbitMQ, nil
}

// Publish Message to
func Publish(message []byte) error {
	rmq, err := Dail()
	if err != nil {
		logger.Error(err, "Error while dialing rabbitMQ")
		return err
	}
	err = rmq.Channel.Publish(
		conf.Exchange, "",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}
	rmq.Close()
	return nil
}
