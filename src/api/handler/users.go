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
func (h *UsersHandler) SendOtp(ctx *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.SendOtp(req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}
