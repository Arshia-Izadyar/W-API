package handler

import (
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

// CreateCity godoc
// @Summary Create a City
// @Description Create a City
// @Tags Cities
// @Accept json
// @produces json
// @Param Request body dto.CreateCityRequest true "Create a City"
// @Success 201 {object} helper.Response{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/city/create [post]
// @Security AuthBearer
func (ch *CityHandler) CreateCity(ctx *gin.Context) {
	Create[dto.CreateCityRequest, dto.CityResponse](ctx, ch.service.GenericCreateCity)
	/*
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
	*/
}

// UpdateCity godoc
// @Summary Update a City
// @Description Update a City
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCityRequest true "Update a City"
// @Success 200 {object} helper.Response{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/city/update/{id} [put]
// @Security AuthBearer
func (ch *CityHandler) UpdateCity(ctx *gin.Context) {
	Update[dto.UpdateCityRequest, dto.CityResponse](ctx, ch.service.GenericUpdateCity)

	/*
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
	*/
}

// DeleteCity godoc
// @Summary Delete a City
// @Description Delete a City
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/city/get/{id} [get]
// @Security AuthBearer
func (ch *CityHandler) GetCityById(ctx *gin.Context) {
	GetById[dto.CityResponse](ctx, ch.service.GenericGetCityById)

	// id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// req := dto.CityResponse{}
	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
	// 	return
	// }
	// res, err := ch.service.GenericGetCityById(ctx, id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
	// 	return
	// }
	// ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))

}

// GetCity godoc
// @Summary Get a City
// @Description Get a City
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/cities/delete/{id} [delete]
// @Security AuthBearer
func (ch *CityHandler) DeleteCity(ctx *gin.Context) {
	Delete(ctx, ch.service.GenericDeleteCity)
	/*
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
	*/
}

// GetCities godoc
// @Summary Get Cities
// @Description Get Cities
// @Tags Cities
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/cities/filter [post]
// @Security AuthBearer
func (ch *CityHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CityResponse](ctx, ch.service.GenericGetByFilter)
	/*
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
	*/

}
