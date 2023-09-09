package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type CarModelPriceHandler struct {
	service *services.CarModelPriceService
}

func NewCarModelPriceHandler(cfg *config.Config) *CarModelPriceHandler {
	return &CarModelPriceHandler{
		service: services.NewCarModelPriceService(cfg),
	}
}

// CreateCarModelPrice godoc
// @Summary Create a CarModelPrice
// @Description Create a CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelPriceRequest true "Create a CarModelPrice"
// @Success 201 {object} helper.Response{result=dto.CarModelPriceResponse} "CarModelPrice response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-price/create [post]
// @Security AuthBearer
func (ch *CarModelPriceHandler) CreateCarModelPrice(ctx *gin.Context) {
	Create[dto.CreateCarModelPriceRequest, dto.CarModelPriceResponse](ctx, ch.service.CreateCarModelPrice)
}

// UpdateCarModelPrice godoc
// @Summary Update a CarModelPrice
// @Description Update a CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPriceRequest true "Update a CarModelPrice"
// @Success 200 {object} helper.Response{result=dto.CarModelPriceResponse} "CarModelPrice response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-price/update/{id} [put]
// @Security AuthBearer
func (ch *CarModelPriceHandler) UpdateCarModelPrice(ctx *gin.Context) {
	Update[dto.UpdateCarModelPriceRequest, dto.CarModelPriceResponse](ctx, ch.service.UpdateCarModelPrice)

}

// DeleteCarModelPrice godoc
// @Summary Delete a CarModelPrice
// @Description Delete a CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-price/get/{id} [get]
// @Security AuthBearer
func (ch *CarModelPriceHandler) GetCarModelPriceById(ctx *gin.Context) {
	GetById[dto.CarModelPriceResponse](ctx, ch.service.GetCarModelPriceById)

}

// GetCarModelPrice godoc
// @Summary Get a CarModelPrice
// @Description Get a CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModelPrice response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-price/delete/{id} [delete]
// @Security AuthBearer
func (ch *CarModelPriceHandler) DeleteCarModelPrice(ctx *gin.Context) {
	Delete(ctx, ch.service.DeleteCarModelPrice)

}

// CarModelPrice godoc
// @Summary Get CarModelPrice
// @Description Get CarModelPrice
// @Tags CarModelPrice
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-price/filter [post]
// @Security AuthBearer
func (ch *CarModelPriceHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CarModelPriceResponse](ctx, ch.service.GetByFilter)
}
