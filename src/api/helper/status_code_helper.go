package helper

import (
	"net/http"
	"wapi/src/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	service_errors.OtpExists: 409,
	service_errors.OtpUsed:   409,
	service_errors.PermissionDenied: 403,
}

func TranslateErrorToStatusCode(err error) int {
	v, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return v
}
