package utils

import (
	"math/rand"
	"time"
)

func GenerateOTP(length int) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := make([]byte, length)

	for i := range otp {
		otp[i] = '0' + byte(rand.Intn(10))
	}

	return string(otp)
}
