package validators

import (
	"wapi/src/common"

	"github.com/go-playground/validator/v10"
)

// `^09[0-9]{9}$`

func IranPhoneNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	return common.PhoneValidator(value)
}
