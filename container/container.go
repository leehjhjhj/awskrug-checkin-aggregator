package container

import (
	"context"
	"log"
	"sync"
	"meetup_checkin/service"
	"meetup_checkin/store"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	once sync.Once
	checkinService *service.CheckinService
)

func getDynamoDBClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-2"))
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return client
}

func GetCheckinService() *service.CheckinService {
	once.Do(func() {
	dynamoClient := getDynamoDBClient()
		checkinStore := store.NewCheckinRepository(dynamoClient)
		checkinService := service.NewCheckinService(checkinStore)
		checkinService = checkinService
	})
	return checkinService
}