package services

import (
	"time"
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
	"wapi/src/pkg/service_errors"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type TokenDTO struct {
	UserId   int
	FullName string
	UserName string
	Phone    string
	Email    string
	Roles    []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	var logger = logging.NewLogger(cfg)
	return &TokenService{
		cfg:    cfg,
		logger: logger,
	}
}

func (ts *TokenService) GenerateToken(td *TokenDTO) (*dto.TokenDetail, error) {
	tokenDetail := &dto.TokenDetail{}
	tokenDetail.AccessTokenExpireTime = time.Now().Add(ts.cfg.Jwt.AccessTokenExpireDuration * time.Minute).Unix()
	tokenDetail.RefreshTokenExpireTime = time.Now().Add(ts.cfg.Jwt.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{}

	accessTokenClaims["user_id"] = td.UserId
	accessTokenClaims["full_name"] = td.FullName
	accessTokenClaims["username"] = td.UserName
	accessTokenClaims["phone"] = td.Phone
	accessTokenClaims["email"] = td.Email
	accessTokenClaims["roles"] = td.Roles
	accessTokenClaims["exp"] = tokenDetail.AccessTokenExpireTime

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	var err error
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(ts.cfg.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}

	refreshTokenClaims["user_id"] = td.UserId
	refreshTokenClaims["exp"] = tokenDetail.RefreshTokenExpireTime

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(ts.cfg.Jwt.Secret))
	if err != nil {
		return nil, err
	}
	return tokenDetail, nil
}

func (ts *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: "cant verify Token"}
		}
		return []byte(ts.cfg.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (ts *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}
	verifyToken, err := ts.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimNotFound}
}
