package helper

import "wapi/src/api/validators"

type Response struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       int                           `json:"result_code"`
	ValidationErrors *[]validators.ValidationError `json:"validation_error"`
	Error            any                           `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *Response {
	return &Response{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *Response {
	return &Response{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode int, err error) *Response {
	ve := validators.GetValidationErrors(err)
	return &Response{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: ve,
	}
}
