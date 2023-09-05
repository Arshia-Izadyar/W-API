package common

import (
	"regexp"
	"wapi/src/config"
	"wapi/src/pkg/logging"
)

var cfg = config.LoadCfg()
var logger = logging.NewLogger(cfg)

const IranPhoneNumberValidatorPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func PhoneValidator(phoneNumber string) bool {
	res, err := regexp.MatchString(IranPhoneNumberValidatorPattern, phoneNumber)
	if err != nil {
		logger.Error(err, logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return res
}
