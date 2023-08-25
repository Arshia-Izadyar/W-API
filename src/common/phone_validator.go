package common

import (
	"log"
	"regexp"
)

const IranPhoneNumberValidatorPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func PhoneValidator(phoneNumber string) bool {
	res, err := regexp.MatchString(IranPhoneNumberValidatorPattern, phoneNumber)
	if err != nil {
		log.Print(err)
	}
	return res
}
