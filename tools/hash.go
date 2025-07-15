package tools

import (
	"crypto/sha256"
	"fmt"
	"os"
	"regexp"
)

func HashPhoneNumber(phone string) string {
    salt := os.Getenv("PHONE_HASH_SALT")
    if salt == "" {
        panic("PHONE_HASH_SALT 환경변수가 설정되지 않았습니다")
    }
    
    // 숫자만 추출
    reg := regexp.MustCompile(`\D`)
    cleanPhone := reg.ReplaceAllString(phone, "")
    
    // salt와 결합
    hashInput := cleanPhone + salt
    
    // SHA256 해시 생성
    hash := sha256.Sum256([]byte(hashInput))
    
    // 16진수 문자열로 반환
    return fmt.Sprintf("%x", hash)
}