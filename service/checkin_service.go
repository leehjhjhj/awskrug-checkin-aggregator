package service

import (
	"fmt"

	"meetup_checkin/handler"
	"meetup_checkin/model"
	"meetup_checkin/store"
)

type CheckinService struct {
	store *store.CheckinRepository
}

func NewCheckinService(store *store.CheckinRepository) *CheckinService {
	return &CheckinService{
		store: store,
	}
}

func (s *CheckinService) GenerateAttendanceReport(eventVersion string) {
	eventCheckins := s.store.GetEventCheckin(eventVersion)
	fmt.Printf("%d개의 레코드 찾았다.\n", len(eventCheckins))

	// 핸드폰-이름 매핑 생성하고 기존 참가자들의 참가 횟수를 계산
	phoneNameMap := s.createPhoneNameMap(eventCheckins)
	existEventCheckins := s.calculateExistingCheckins(eventCheckins)

	// CSV에서 새로운 체크인 데이터 추출
	newCheckinPhoneNumbers, csvPhoneNameMap := handler.GetNewEventCheckinFromCSV()

	// CSV의 이름 정보를 기존 매핑에 병합 + 최종 참가자 목록 생성
	s.mergePhoneNameMaps(phoneNameMap, csvPhoneNameMap)
	finalCheckins := s.mergeFinalCheckins(newCheckinPhoneNumbers, existEventCheckins)

	// CSV 파일 생성
	handler.MakeNewCheckinCSV(finalCheckins, phoneNameMap)
}

func (s *CheckinService) createPhoneNameMap(eventCheckins []model.EventCheckin) map[string]string {
	phoneNameMap := make(map[string]string)
	for _, checkin := range eventCheckins {
		phoneNameMap[checkin.Phone] = checkin.Name
	}
	return phoneNameMap
}

func (s *CheckinService) calculateExistingCheckins(eventCheckins []model.EventCheckin) map[string]int {
	existEventCheckins := make(map[string]int)
	for _, checkin := range eventCheckins {
		existEventCheckins[checkin.Phone]++
	}
	return existEventCheckins
}

func (s *CheckinService) mergePhoneNameMaps(phoneNameMap, csvPhoneNameMap map[string]string) {
	for phone, name := range csvPhoneNameMap {
		phoneNameMap[phone] = name
	}
}

func (s *CheckinService) mergeFinalCheckins(newCheckins, existingCheckins map[string]int) map[string]int {
	for phoneNumber, count := range newCheckins {
		if phoneNumber != "" {
			newCheckins[phoneNumber] = count + existingCheckins[phoneNumber]
		}
	}
	return newCheckins
}
