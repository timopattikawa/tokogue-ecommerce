package util

import (
	"fmt"
	"math/rand"
)

func GenerateOTP() string {
	result := ""

	for i := 0; i < 4; i++ {
		r := rand.Intn(10)
		tmp := fmt.Sprintf("%d", r)
		result += tmp
	}

	return result
}
