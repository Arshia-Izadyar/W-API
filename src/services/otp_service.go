package services

import (
	"fmt"
	"time"
	"wapi/src/config"
	"wapi/src/constants"
	"wapi/src/data/cache"
	"wapi/src/pkg/logging"
	"wapi/src/pkg/service_errors"

	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redis  *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	rd := cache.GetRedis()
	return &OtpService{logger: logger, cfg: cfg, redis: rd}

}

func (s *OtpService) SetOtp(mobileNumber string, otp string) *service_errors.ServiceError {
	key := fmt.Sprintf("%s:%s", constants.DefaultRedisKey, mobileNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}
	result, err := cache.Get[OtpDto](s.redis, key)
	if err == nil && !result.Used {
		return &service_errors.ServiceError{Err: err, EndUserMessage: service_errors.OtpExists}
	} else if err == nil && result.Used {
		return &service_errors.ServiceError{Err: err, EndUserMessage: service_errors.OtpUsed}
	}
	err = cache.Set(s.redis, key, *val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return &service_errors.ServiceError{Err: err, EndUserMessage: "cant get btw"}
	}
	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.DefaultRedisKey, mobileNumber)
	res, err := cache.Get[OtpDto](s.redis, key)
	if err != nil {
		return err
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{Err: err, EndUserMessage: service_errors.OtpExists}
	} else if err == nil && !res.Used && res.Value != otp {
		return &service_errors.ServiceError{Err: err, EndUserMessage: "otp not valid"}
	} else if err == nil && !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redis, key, res, s.cfg.Otp.ExpireTime*time.Second)
		if err != nil {
			return err
		}

	}
	return nil
}
