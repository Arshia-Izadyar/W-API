package validators

import (
	"wapi/src/common"

	"github.com/go-playground/validator/v10"
)

func PassWordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}
	return common.CheckPassword(value)
}
