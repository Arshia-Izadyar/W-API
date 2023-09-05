package middleware

import (
	"net/http"
	"strings"
	"wapi/src/api/helper"
	"wapi/src/config"
	"wapi/src/constants"
	"wapi/src/pkg/service_errors"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenServise = services.NewTokenService(cfg)
	return func(ctx *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		key := ctx.GetHeader(constants.AuthenTicationHeaderKey)
		if key == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenNotPresent}
		} else {
			token := strings.Split(key, " ")[1]
			claimMap, err = tokenServise.GetClaims(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(nil, false, -2, err))
			return
		}
		ctx.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		ctx.Set(constants.FullNameKey, claimMap[constants.FullNameKey])
		ctx.Set(constants.UserNameKey, claimMap[constants.UserNameKey])
		ctx.Set(constants.PhoneKey, claimMap[constants.PhoneKey])
		ctx.Set(constants.EmailKey, claimMap[constants.EmailKey])
		ctx.Set(constants.RolesKey, claimMap[constants.RolesKey])
		ctx.Set(constants.ExpKey, claimMap[constants.ExpKey])

		ctx.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -1))
			return
		}
		rolesV, ok := ctx.Keys[constants.RolesKey]
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -1))
			return
		}
		roles := rolesV.([]interface{})
		val := map[string]int{}
		for _, v := range roles {
			val[v.(string)] = 0
		}
		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				ctx.Next()
				return
			}
			ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -1))
			return
		}
	}
}
