package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func GetDynamodbConn(ctx context.Context) *dynamodb.DynamoDB {
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
	return dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:4566")})
}

func Close(db *dynamodb.DynamoDB) {
	db = nil
}
