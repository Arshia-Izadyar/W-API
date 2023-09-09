package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type CarModelPropertyHandler struct {
	service *services.CarModelPropertyService
}

func NewCarModelPropertyHandler(cfg *config.Config) *CarModelPropertyHandler {
	return &CarModelPropertyHandler{
		service: services.NewCarModelPropertyService(cfg),
	}
}

// CreateCarModelProperty godoc
// @Summary Create a CarModelProperty
// @Description Create a CarModelProperty
// @Tags CarModelProperty
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelPropertyRequest true "Create a CarModelProperty"
// @Success 201 {object} helper.Response{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-property/create [post]
// @Security AuthBearer
func (ch *CarModelPropertyHandler) CreateCarModelProperty(ctx *gin.Context) {
	Create[dto.CreateCarModelPropertyRequest, dto.CarModelPropertyResponse](ctx, ch.service.CreateCarModelProperty)
}

// UpdateCarModelProperty godoc
// @Summary Update a CarModelProperty
// @Description Update a CarModelProperty
// @Tags CarModelProperty
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPropertyRequest true "Update a CarModelProperty"
// @Success 200 {object} helper.Response{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-property/update/{id} [put]
// @Security AuthBearer
func (ch *CarModelPropertyHandler) UpdateCarModelProperty(ctx *gin.Context) {
	Update[dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse](ctx, ch.service.UpdateCarModelProperty)

}

// DeleteCarModelProperty godoc
// @Summary Delete a CarModelProperty
// @Description Delete a CarModelProperty
// @Tags CarModelProperty
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-property/get/{id} [get]
// @Security AuthBearer
func (ch *CarModelPropertyHandler) GetCarModelPropertyById(ctx *gin.Context) {
	GetById[dto.CarModelPropertyResponse](ctx, ch.service.GetCarModelPropertyById)

}

// GetCarModelProperty godoc
// @Summary Get a CarModelProperty
// @Description Get a CarModelProperty
// @Tags CarModelProperty
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModelProperty response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-property/delete/{id} [delete]
// @Security AuthBearer
func (ch *CarModelPropertyHandler) DeleteCarModelProperty(ctx *gin.Context) {
	Delete(ctx, ch.service.DeleteCarModelProperty)

}

// CarModelProperty godoc
// @Summary Get CarModelProperty
// @Description Get CarModelProperty
// @Tags CarModelProperty
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-property/filter [post]
// @Security AuthBearer
func (ch *CarModelPropertyHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CarModelPropertyResponse](ctx, ch.service.GetByFilter)
}
