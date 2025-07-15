package tools

import "regexp"

func CleanPhoneNumber(phone string) string {
	reg := regexp.MustCompile(`\D`)
	return reg.ReplaceAllString(phone, "")
}