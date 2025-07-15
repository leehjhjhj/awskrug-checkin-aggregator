package handler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"meetup_checkin/model"
	"meetup_checkin/tools"
	"meetup_checkin/config"

)

func GetNewEventCheckinFromCSV() (map[string]int, map[string]string) {
	file, err := os.Open(config.CSVPath)
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read csv file: %v", err)
	}

	headers := records[0]
	targetIndex := -1
	for i, header := range headers {
		if header == config.TargetHeader {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		log.Fatalf("target header not found")
	}

	phoneNumbers := make([]string, 0)
	csvPhoneNameMap := make(map[string]string)

	for _, record := range records[1:] {
		if targetIndex < len(record) {
			checkinInfo := strings.Split(record[targetIndex], "/")
			if len(checkinInfo) >= 2 {
				name := strings.TrimSpace(checkinInfo[0])
				cleanedPhone := tools.CleanPhoneNumber(checkinInfo[len(checkinInfo)-1])
				hashedPhone := tools.HashPhoneNumber(cleanedPhone)

				if name != "" && hashedPhone != "" {
					csvPhoneNameMap[hashedPhone] = name
				}

				phoneNumbers = append(phoneNumbers, hashedPhone)
			}
		}
	}

	newCheckinPhoneNumbers := make(map[string]int)
	for _, phoneNumber := range phoneNumbers {
		if strings.TrimSpace(phoneNumber) != "" {
			newCheckinPhoneNumbers[phoneNumber]++
		} else {
			fmt.Printf("빈 핸드폰 발견하여 제외했습니다. %v\n", phoneNumber)
		}
	}

	return newCheckinPhoneNumbers, csvPhoneNameMap
}

func MakeNewCheckinCSV(newCheckinPhoneNumbers map[string]int, phoneNameMap map[string]string) {
	participants := make([]model.Result, 0)
	for phone, count := range newCheckinPhoneNumbers {
		resultModel := model.Result{Name: phoneNameMap[phone], Count: count}
		participants = append(participants, resultModel)
	}
	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Count > participants[j].Count
	})
	date := time.Now().Format("2006-01-02")
	csvFile, err := os.Create(fmt.Sprintf("data/%s_기준_출석_횟수(핸드폰기준).csv", date))
	if err != nil {
		log.Fatalf("failed to create csv file: %v", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{"이름", "참가 횟수"})

	for _, participant := range participants {
		record := []string{
			participant.Name,
			strconv.Itoa(participant.Count),
		}
		writer.Write(record)
	}
}