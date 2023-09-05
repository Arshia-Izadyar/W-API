package handler

import (
	"net/http"
	"wapi/src/api/helper"
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *services.UserServices
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	services := services.NewUserService(cfg)
	return &UsersHandler{
		service: services,
	}
}

// SendOtp godoc
// @Summary send otp
// @Description send otp
// @Tags User
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.Response "Success"
// @Failure 400 {object} helper.Response "Failed"
// @Failure 409 {object} helper.Response "Failed"
// @Router /v1/users/send-otp [post]
func (uh *UsersHandler) SendOtp(ctx *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = uh.service.SendOtp(req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}

// LoginByUsername godoc
// @Summary LoginByUsername
// @Description LoginByUsername
// @Tags User
// @Accept  json
// @Produce  json
// @Param Request body dto.LoginByUsername true "LoginByUsername"
// @Success 201 {object} helper.Response "Success"
// @Failure 400 {object} helper.Response "Failed"
// @Failure 409 {object} helper.Response "Failed"
// @Router /v1/users/login-by-username [post]
func (uh *UsersHandler) LoginByUsername(ctx *gin.Context) {
	req := new(dto.LoginByUsername)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	token, err := uh.service.LoginByUsername(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(token, true, 0))
}

// RegisterLoginByPhone godoc
// @Summary RegisterLoginByPhone
// @Description RegisterLoginByPhone
// @Tags User
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterLoginByPhone true "RegisterLoginByPhone"
// @Success 201 {object} helper.Response "Success"
// @Failure 400 {object} helper.Response "Failed"
// @Failure 409 {object} helper.Response "Failed"
// @Router /v1/users/register-login-by-phone [post]
func (uh *UsersHandler) RegisterLoginByPhone(ctx *gin.Context) {
	req := new(dto.RegisterLoginByPhone)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	token, err := uh.service.RegisterLoginByPhone(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(token, true, 0))

}

// RegisterByUsername godoc
// @Summary RegisterByUsername
// @Description sRegisterByUsername
// @Tags User
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterUserByUsername true "RegisterUserByUsername"
// @Success 201 {object} helper.Response "Success"
// @Failure 400 {object} helper.Response "Failed"
// @Failure 409 {object} helper.Response "Failed"
// @Router /v1/users/register-by-username [post]
func (uh *UsersHandler) RegisterByUsername(ctx *gin.Context) {
	req := new(dto.RegisterUserByUsername)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	err = uh.service.RegisterByUsername(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))

}
