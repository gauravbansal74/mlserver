package watcher

import (
	"encoding/json"

	"github.com/gauravbansal74/mlserver/pkg/logger"
	"github.com/gauravbansal74/mlserver/pkg/rabbitmq"
)

const (
	fileFound   = "NEW"
	fileError   = "FAILED"
	fileSuccess = "SUCCESS"
)

// Consumer - RabbitMQ message consumer
func Consumer(rmq *rabbitmq.RabbitMQ, filePath string) {
	defer rmq.Close()
	forever := make(chan bool)
	go func() {
		for d := range rmq.Deliveries {
			logger.Info("Message Received", logger.Fields{"message": string(d.Body)})
			err := processMessage(d.Body, filePath)
			if err != nil {
				logger.Error(err, "Error while processing message Received", logger.Fields{"message": string(d.Body)})
				d.Reject(false)
			} else {
				logger.Info("Message Processed succesfully", logger.Fields{"message": string(d.Body)})
				d.Ack(false)
			}
		}
	}()
	logger.Info("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// processMessage Data
func processMessage(message []byte, filePath string) error {
	fileMsg := FileMsg{}
	err := json.Unmarshal(message, &fileMsg)
	if err != nil {
		return err
	}
	err = Create(fileMsg.Name, fileFound, fileFound)
	if err != nil {
		return err
	}

	lineCount, err := FileDataValidation(filePath + fileMsg.Name)
	if err != nil {
		err = Create(fileMsg.Name, "Line Number: "+string(lineCount)+" > "+err.Error(), fileError)
		if err != nil {
			return err
		}
	}
	err = Create(fileMsg.Name, "File Processed", fileSuccess)
	if err != nil {
		return err
	}
	return nil
}
