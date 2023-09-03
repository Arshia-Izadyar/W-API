package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,phone,min=11,max=11"`
}
