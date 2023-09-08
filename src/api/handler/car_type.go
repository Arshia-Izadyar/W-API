package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type CarTypeHandler struct {
	service *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{
		service: services.NewCarTypeService(cfg),
	}
}

// CreateCarType godoc
// @Summary Create a CarType
// @Description Create a CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param Request body dto.CreateCarTypeRequest true "Create a CarType"
// @Success 201 {object} helper.Response{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-type/create [post]
// @Security AuthBearer
func (ch *CarTypeHandler) CreateCarType(ctx *gin.Context) {
	Create[dto.CreateCarTypeRequest, dto.CarTypeResponse](ctx, ch.service.GenericCreateCarType)
}

// UpdateCarType godoc
// @Summary Update a CarType
// @Description Update a CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarTypeRequest true "Update a CarType"
// @Success 200 {object} helper.Response{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-type/update/{id} [put]
// @Security AuthBearer
func (ch *CarTypeHandler) UpdateCarType(ctx *gin.Context) {
	Update[dto.UpdateCarTypeRequest, dto.CarTypeResponse](ctx, ch.service.GenericUpdateCarType)

}

// DeleteCarType godoc
// @Summary Delete a CarType
// @Description Delete a CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-type/get/{id} [get]
// @Security AuthBearer
func (ch *CarTypeHandler) GetCarTypeById(ctx *gin.Context) {
	GetById[dto.CarTypeResponse](ctx, ch.service.GenericGetCarTypeById)

}

// GetCarType godoc
// @Summary Get a CarType
// @Description Get a CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarType response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-type/delete/{id} [delete]
// @Security AuthBearer
func (ch *CarTypeHandler) DeleteCarType(ctx *gin.Context) {
	Delete(ctx, ch.service.GenericDeleteCarType)

}

// CarType godoc
// @Summary Get CarType
// @Description Get CarType
// @Tags CarType
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-type/filter [post]
// @Security AuthBearer
func (ch *CarTypeHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CarTypeResponse](ctx, ch.service.GenericGetByFilter)
}
