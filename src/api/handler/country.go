package handler

import (
	"errors"
	"net/http"
	"strconv"
	"wapi/src/api/helper"
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *services.GenericCountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler{ 
	srv := services.NewGenericCountryService(cfg)
	return &CountryHandler{
		service: srv,
	}
}

func (ch *CountryHandler) CreateCountry(ctx *gin.Context) {
	req := dto.CreateUpdateCountryDTO{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res ,err := ch.service.GenericCreateCountry(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

func (ch *CountryHandler) UpdateCountry(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.CreateUpdateCountryDTO{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := ch.service.GenericUpdateCountry(ctx, id, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

func (ch *CountryHandler) DeleteCountry(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, -1, errors.New("not found id = 0")))
		return
	}

	err := ch.service.GenericDeleteCountry(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(gin.H{"Status":"Deleted"}, true, 0))

}

func (ch *CountryHandler) GetCountryById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.CountryResponse{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := ch.service.GenericGetCountryById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}