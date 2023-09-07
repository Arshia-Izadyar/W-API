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

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		service: services.NewCityService(cfg),
	}
}

func (ch *CityHandler) CreateCity(ctx *gin.Context) {
	req := dto.CreateCityRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := ch.service.GenericCreateCity(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

func (ch *CityHandler) UpdateCity(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.UpdateCityRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := ch.service.GenericUpdateCity(ctx, id, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

func (ch *CityHandler) GetCityById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.CityResponse{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := ch.service.GenericGetCityById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

func (ch *CityHandler) DeleteCity(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, -1, errors.New("not found id = 0")))
		return
	}
	err := ch.service.GenericDeleteCity(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(gin.H{"Status": "Deleted"}, true, 0))

}

func (ch *CityHandler) GetByFilter(ctx *gin.Context) {
	req := dto.PaginationInputWithFilter{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	res, err := ch.service.GenericGetByFilter(ctx, &req)
	// res, err := ch.service.GenericGetByFilter(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}
