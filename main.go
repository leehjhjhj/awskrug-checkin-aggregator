package main

import (
	"fmt"
	"log"

	"meetup_checkin/handler"
	"meetup_checkin/store"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env 파일 로드 실패 (무시하고 계속): %v", err)
	}
	store := store.NewCheckinRepository(handler.GetDynamoDBClient())
	eventCheckins := store.GetEventCheckin("1")

	fmt.Printf("%d개의 레코드 찾았다.\n", len(eventCheckins))

	phoneNameMap := make(map[string]string)
	for _, checkin := range eventCheckins {
		phoneNameMap[checkin.Phone] = checkin.Name
	}

	existEventCheckins := make(map[string]int)
	for _, checkin := range eventCheckins {
		existEventCheckins[checkin.Phone]++
	}
	newCheckinPhoneNumbers, csvPhoneNameMap := handler.GetNewEventCheckinFromCSV()

	for phone, name := range csvPhoneNameMap {
		phoneNameMap[phone] = name
	}

	/*
	최종 참가자 목록 생성
	*/
	for phoneNumber, count := range newCheckinPhoneNumbers {
		if phoneNumber != "" {
			newCheckinPhoneNumbers[phoneNumber] = count + existEventCheckins[phoneNumber]
		}
	}
	handler.MakeNewCheckinCSV(newCheckinPhoneNumbers, phoneNameMap)
}
