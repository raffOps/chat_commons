package dynamodb

import (
	"context"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

func GetDynamoDbTestContainer(ctx context.Context) (*localstack.LocalStackContainer, error) {
	return localstack.Run(ctx, "localstack/localstack:1.4.0")
}
