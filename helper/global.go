package helper

import (
	"encoding/json"
	uuid2 "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	mathRand "math/rand"
)

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func RandomString(length int) string {
	bytes := make([]byte, length)

	for i := 0; i < length; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}

	return string(bytes)
}

func RandomInt(min int, max int) int {
	return min + mathRand.Intn(max-min)
}

func JSONEncode(data interface{}) string {
	jsonResult, _ := json.Marshal(data)

	return string(jsonResult)
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func ArrayStringContains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}

func ArrayInterfaceContains(arr []interface{}, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}

func IsUUID(uuid string) bool {
	u, err := uuid2.Parse(uuid)
	if err != nil {
		return false
	}

	if u.Version() == 4 {
		return true
	}

	return false
}
