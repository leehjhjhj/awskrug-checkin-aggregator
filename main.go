package main

import (
	"log"

	"meetup_checkin/handler"
	"meetup_checkin/service"
	"meetup_checkin/store"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env 파일 로드 실패 (무시하고 계속): %v", err)
	}

	dynamoClient := handler.GetDynamoDBClient()
	checkinStore := store.NewCheckinRepository(dynamoClient)
	checkinService := service.NewCheckinService(checkinStore)

	checkinService.GenerateAttendanceReport("1")
}
