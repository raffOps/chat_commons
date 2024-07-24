package rabbitmq

import (
	"fmt"
	"github.com/raffops/chat/pkg"
	"github.com/raffops/chat/pkg/logger"
	"go.uber.org/zap"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func GetRabbitMQConnection() *amqp.Connection {
	envVariables := []string{"RABBITMQ_HOST", "RABBITMQ_PORT", "RABBITMQ_USER", "RABBITMQ_PASSWORD"}
	pkg.SanityCheck(logger.Logger, envVariables)

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Fatal("Connection to rabbitmq failed", zap.Error(err))
	}
	return conn
}

func CloseRabbitMQConnection(connection *amqp.Connection) error {
	err := connection.Close()
	return err
}
