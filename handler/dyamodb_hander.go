package handler

import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamoDBClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-2"))
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return client
}