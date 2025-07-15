package store

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"meetup_checkin/model"
)

type CheckinRepository struct {
	client *dynamodb.Client
}

func NewCheckinRepository(client *dynamodb.Client) *CheckinRepository {
	return &CheckinRepository{
		client: client,
	}
}

func (r *CheckinRepository) GetEventCheckin(eventVersion string) []model.EventCheckin {
	result, err := r.client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:        aws.String("prod-event-checkin"),
		FilterExpression: aws.String("event_version = :event_version"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":event_version": &types.AttributeValueMemberS{Value: eventVersion},
		},
	})
	if err != nil {
		log.Fatalf("failed to scan: %v", err)
	}
	var checkins []model.EventCheckin
	err = attributevalue.UnmarshalListOfMaps(result.Items, &checkins)
	if err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}
	return checkins
}