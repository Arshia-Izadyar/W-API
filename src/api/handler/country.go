package handler

import (
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
	ss      *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	srv := services.NewGenericCountryService(cfg)
	ss := services.NewCountryService(cfg)
	return &CountryHandler{
		service: srv,
		ss:      ss,
	}
}

// CreateCountry godoc
// @Summary Create a Country
// @Description Create a Country
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdateCountryDTO true "Create a Country"
// @Success 201 {object} helper.Response "Country response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/country/create [post]
// @Security AuthBearer
func (ch *CountryHandler) CreateCountry(ctx *gin.Context) {
	Create[dto.CreateUpdateCountryDTO, dto.CountryResponse](ctx, ch.service.GenericCreateCountry)
	/*
		req := dto.CreateUpdateCountryDTO{}
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
			return
		}
		res, err := ch.service.GenericCreateCountry(ctx, &req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
			return
		}
		ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
	*/
}

// UpdateCountry godoc
// @Summary Update a Country
// @Description Update a Country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCountryDTO true "Create a Country"
// @Success 200 {object} helper.Response "Country response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/country/update/{id} [put]
// @Security AuthBearer
func (ch *CountryHandler) UpdateCountry(ctx *gin.Context) {
	Update[dto.CreateUpdateCountryDTO, dto.CountryResponse](ctx, ch.service.GenericUpdateCountry)
	/*
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
	*/
}

// DeleteCountry godoc
// @Summary Delete a Country
// @Description Delete a Country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 204 {object} helper.Response "Country response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/country/delete/{id} [delete]
// @Security AuthBearer
func (ch *CountryHandler) DeleteCountry(ctx *gin.Context) {
	Delete(ctx, ch.service.GenericDeleteCountry)

	/*
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
		ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(gin.H{"Status": "Deleted"}, true, 0))
	*/
}

// GetCountryById godoc
// @Summary Get a Country
// @Description Get a Country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "Country response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/country/get/{id} [get]
// @Security AuthBearer
func (ch *CountryHandler) GetCountryById(ctx *gin.Context) {
	GetById[dto.CountryResponse](ctx, ch.service.GenericGetCountryById)
	/*
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
	*/

}

// GetCountryByFilter godoc
// @Summary GetCountryByFilter
// @Description GetCountryByFilter
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "GetCountryByFilter"
// @Success 200 {object} helper.Response "Country response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/country/filter [post]
// @Security AuthBearer
func (ch *CountryHandler) GetCountryByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CountryResponse](ctx, ch.service.GenericGetByFilter)
	/*
		req := dto.PaginationInputWithFilter{}
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
			return
		}
		res, err := ch.service.GenericGetByFilter(ctx, &req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
			return
		}
		ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
	*/
}

func (ch *CountryHandler) GetCitiesById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.CountryResponse{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := ch.ss.GetCitiesByCountryId(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}
