package services

import (
	"wapi/src/common"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/dto"
	"wapi/src/pkg/logging"

	"gorm.io/gorm"
)

type UserServices struct {
	Logger     *logging.Logger
	Cfg        *config.Config
	OtpService *OtpService
	Db         *gorm.DB
}

func NewUserService(cfg *config.Config) *UserServices {
	var db = db.GetDB()
	var logger = logging.NewLogger(cfg)
	return &UserServices{
		Logger:     &logger,
		Cfg:        cfg,
		OtpService: NewOtpService(cfg),
		Db:         db,
	}
}

func (u *UserServices) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := u.OtpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
