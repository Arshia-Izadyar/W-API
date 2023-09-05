package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,phone,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpireTime  int64  `json:"access_token_expire"`
	RefreshTokenExpireTime int64  `json:"refresh_token_expire"`
}

type RegisterUserByUsername struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password" binding:"password"`
}

type RegisterLoginByPhone struct {
	Phone string `json:"phone" binding:"required,phone,min=11,max=11"`
	Otp   string `json:"otp" binding:"required"`
}

type LoginByUsername struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
}
