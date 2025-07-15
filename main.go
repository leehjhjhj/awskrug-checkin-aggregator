package main

import (
	"log"

	"meetup_checkin/container"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env 파일 로드 실패 (무시하고 계속): %v", err)
	}

	checkinService := container.GetCheckinService()
	checkinService.GenerateAttendanceReport("1")
}
