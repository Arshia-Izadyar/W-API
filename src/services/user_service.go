package services

import (
	"database/sql"
	"wapi/src/common"
	"wapi/src/config"
	"wapi/src/constants"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
	"wapi/src/pkg/service_errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices struct {
	Logger     *logging.Logger
	Cfg        *config.Config
	OtpService *OtpService
	Token      *TokenService
	Db         *gorm.DB
}

func NewUserService(cfg *config.Config) *UserServices {
	var db = db.GetDB()
	var logger = logging.NewLogger(cfg)
	var otp = NewOtpService(cfg)
	var token = NewTokenService(cfg)
	return &UserServices{
		Logger:     &logger,
		Cfg:        cfg,
		OtpService: otp,
		Token:      token,
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

func (u *UserServices) existsByEmail(email string) (bool, error) {
	var exists bool
	err := u.Db.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", email).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserServices) existsByUsername(username string) (bool, error) {
	var exists bool
	err := u.Db.Model(&models.User{}).Select("count(*) > 1").Where("user_name = ?", username).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserServices) existsByPhone(phone string) (bool, error) {
	var exists bool
	err := u.Db.Model(&models.User{}).Select("count(*) > 1").Where("phone_number = ?", phone).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserServices) getDefaultRole() (roleId int, err error) {

	if err = u.Db.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}

func (u *UserServices) RegisterByUsername(req *dto.RegisterUserByUsername) error {
	usr := models.User{
		UserName:  req.Username,
		Email:     sql.NullString{String: req.Email, Valid: true},
		FirstName: sql.NullString{String: req.FirstName, Valid: true},
		LastName:  sql.NullString{String: req.LastName, Valid: true},
		Password:  req.Password,
	}
	exists, err := u.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = u.existsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	usr.Password = string(bs)
	roleID, err := u.getDefaultRole()
	if err != nil {
		return err
	}
	tx := u.Db.Begin()

	err = tx.Create(&usr).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleID, UserId: usr.Id}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (u *UserServices) RegisterLoginByPhone(req *dto.RegisterLoginByPhone) (*dto.TokenDetail, error) {
	err := u.OtpService.ValidateOtp(req.Phone, req.Otp)
	if err != nil {
		return nil, err
	}
	exists, err := u.existsByPhone(req.Phone)
	if err != nil {
		return nil, err
	}

	usr := models.User{PhoneNumber: sql.NullString{Valid: true, String: req.Phone}, UserName: req.Phone}

	if exists {
		var user models.User
		err = u.Db.Model(&models.User{}).Where("user_name = ?", usr.UserName).Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("role")
		}).Find(&user).Error
		if err != nil {
			return nil, err
		}
		tDTO := TokenDTO{
			UserId:   user.Id,
			FullName: user.FirstName.String + user.LastName.String,
			UserName: user.UserName,
			Email:    user.Email.String,
			Phone:    user.PhoneNumber.String,
		}
		if len(*user.UserRoles) > 0 {
			for _, role := range *user.UserRoles {
				tDTO.Roles = append(tDTO.Roles, role.Role.Name)
			}
		}
		token, err := u.Token.GenerateToken(&tDTO)
		if err != nil {
			return nil, err
		}
		return token, nil
	} else {
		bs, err := bcrypt.GenerateFromPassword([]byte("changeme"), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}
		usr.Password = string(bs)
		roleId, err := u.getDefaultRole()
		if err != nil {
			return nil, err
		}
		tx := u.Db.Begin()
		err = tx.Create(&usr).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		err = tx.Create(&models.UserRole{UserId: usr.Id, RoleId: roleId}).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()

		var user models.User
		err = u.Db.Model(&models.User{}).Where("user_name = ?", usr.UserName).Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).Find(&user).Error
		if err != nil {
			return nil, err
		}
		tdto := TokenDTO{
			UserId:   user.Id,
			UserName: user.UserName,
			Phone:    user.PhoneNumber.String,
			Email:    user.Email.String,
			FullName: user.FirstName.String + user.LastName.String,
		}
		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				tdto.Roles = append(tdto.Roles, ur.Role.Name)
			}
		}
		token, err := u.Token.GenerateToken(&tdto)
		if err != nil {
			return nil, err
		}
		return token, nil

	}
}

func (u *UserServices) LoginByUsername(req *dto.LoginByUsername) (*dto.TokenDetail, error) {
	var user models.User
	err := u.Db.Model(&models.User{}).Where("user_name = ?", req.Username).Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("Role")
	}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.WrongPassword}
	}
	tdto := TokenDTO{
		UserId:   user.Id,
		UserName: user.UserName,
		Phone:    user.PhoneNumber.String,
		Email:    user.Email.String,
		FullName: user.FirstName.String + user.LastName.String,
	}
	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
	}
	token, err := u.Token.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil
}
