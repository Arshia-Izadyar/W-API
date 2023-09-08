package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type CarModelHandler struct {
	service *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{
		service: services.NewCarModelService(cfg),
	}
}

// CreateCarModel godoc
// @Summary Create a CarModel
// @Description Create a CarModel
// @Tags CarModel
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelRequest true "Create a CarModel"
// @Success 201 {object} helper.Response{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-model/create [post]
// @Security AuthBearer
func (ch *CarModelHandler) CreateCarModel(ctx *gin.Context) {
	Create[dto.CreateCarModelRequest, dto.CarModelResponse](ctx, ch.service.GenericCreateCarModel)
}

// UpdateCarModel godoc
// @Summary Update a CarModel
// @Description Update a CarModel
// @Tags CarModel
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelRequest true "Update a CarModel"
// @Success 200 {object} helper.Response{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-model/update/{id} [put]
// @Security AuthBearer
func (ch *CarModelHandler) UpdateCarModel(ctx *gin.Context) {
	Update[dto.UpdateCarModelRequest, dto.CarModelResponse](ctx, ch.service.GenericUpdateCarModel)

}

// DeleteCarModel godoc
// @Summary Delete a CarModel
// @Description Delete a CarModel
// @Tags CarModel
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-model/get/{id} [delete]
// @Security AuthBearer
func (ch *CarModelHandler) GetCarModelById(ctx *gin.Context) {
	GetById[dto.CarModelResponse](ctx, ch.service.GenericGetCarModelById)

}

// GetCarModel godoc
// @Summary Get a CarModel
// @Description Get a CarModel
// @Tags CarModel
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModel response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-model/delete/{id} [get]
// @Security AuthBearer
func (ch *CarModelHandler) DeleteCarModel(ctx *gin.Context) {
	Delete(ctx, ch.service.GenericDeleteCarModel)

}

// CarModel godoc
// @Summary Get CarModel
// @Description Get CarModel
// @Tags CarModel
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-model/filter [post]
// @Security AuthBearer
func (ch *CarModelHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CarModelResponse](ctx, ch.service.GenericGetByFilter)
}
