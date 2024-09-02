package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/raffops/chat_commons/pkg"
	"github.com/raffops/chat_commons/pkg/logger"
	"os"
)

func GetDynamodbConn(ctx context.Context) *dynamodb.DynamoDB {
	pkg.SanityCheck(logger.Logger, []string{"LOCALSTACK_PORT"})

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	if os.Getenv("ENV") == "PRD" {
		return dynamodb.New(sess)
	}

	localstackPort := os.Getenv("LOCALSTACK_PORT")
	return dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:" + localstackPort)})
}

func Close(db *dynamodb.DynamoDB) {
	db = nil
}
