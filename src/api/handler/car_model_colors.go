package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type CarModelColorHandler struct {
	service *services.CarModelColorService
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	return &CarModelColorHandler{
		service: services.NewCarModelColorService(cfg),
	}
}

// CreateCarModelColor godoc
// @Summary Create a CarModelColor
// @Description Create a CarModelColor
// @Tags CarModelColor
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelColorRequest true "Create a CarModelColor"
// @Success 201 {object} helper.Response{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-color/create [post]
// @Security AuthBearer
func (ch *CarModelColorHandler) CreateCarModelColor(ctx *gin.Context) {
	Create[dto.CreateCarModelColorRequest, dto.CarModelColorResponse](ctx, ch.service.GenericCreateCarModelColor)
}

// UpdateCarModelColor godoc
// @Summary Update a CarModelColor
// @Description Update a CarModelColor
// @Tags CarModelColor
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelColorRequest true "Update a CarModelColor"
// @Success 200 {object} helper.Response{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-color/update/{id} [put]
// @Security AuthBearer
func (ch *CarModelColorHandler) UpdateCarModelColor(ctx *gin.Context) {
	Update[dto.UpdateCarModelColorRequest, dto.CarModelColorResponse](ctx, ch.service.GenericUpdateCarModelColor)

}

// DeleteCarModelColor godoc
// @Summary Delete a CarModelColor
// @Description Delete a CarModelColor
// @Tags CarModelColor
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-color/get/{id} [get]
// @Security AuthBearer
func (ch *CarModelColorHandler) GetCarModelColorById(ctx *gin.Context) {
	GetById[dto.CarModelColorResponse](ctx, ch.service.GenericGetCarModelColorById)

}

// GetCarModelColor godoc
// @Summary Get a CarModelColor
// @Description Get a CarModelColor
// @Tags CarModelColor
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModelColor response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-color/delete/{id} [delete]
// @Security AuthBearer
func (ch *CarModelColorHandler) DeleteCarModelColor(ctx *gin.Context) {
	Delete(ctx, ch.service.GenericDeleteCarModelColor)

}

// CarModelColor godoc
// @Summary Get CarModelColor
// @Description Get CarModelColor
// @Tags CarModelColor
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-color/filter [post]
// @Security AuthBearer
func (ch *CarModelColorHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CarModelColorResponse](ctx, ch.service.GenericGetByFilter)
}
