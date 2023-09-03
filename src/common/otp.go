package common

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GenerateOtp() string {
	rand.Seed(time.Now().UnixNano())

	min := int(math.Pow10(cfg.Otp.Digits - 1))
	max := int(math.Pow10(cfg.Otp.Digits) - 1)
	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
