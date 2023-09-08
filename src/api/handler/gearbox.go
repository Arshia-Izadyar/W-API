package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type GearBoxHandler struct {
	service *services.GearBoxService
}

func NewGearBoxHandler(cfg *config.Config) *GearBoxHandler {
	return &GearBoxHandler{
		service: services.NewGearBoxService(cfg),
	}
}

// CreateGearBox godoc
// @Summary Create a GearBox
// @Description Create a GearBox
// @Tags GearBox
// @Accept json
// @produces json
// @Param Request body dto.CreateGearBoxRequest true "Create a GearBox"
// @Success 201 {object} helper.Response{result=dto.GearBoxResponse} "GearBox response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/gearbox/create [post]
// @Security AuthBearer
func (ch *GearBoxHandler) CreateGearBox(ctx *gin.Context) {
	Create[dto.CreateGearBoxRequest, dto.GearBoxResponse](ctx, ch.service.GenericCreateGearBox)
}

// UpdateGearBox godoc
// @Summary Update a GearBox
// @Description Update a GearBox
// @Tags GearBox
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearBoxRequest true "Update a GearBox"
// @Success 200 {object} helper.Response{result=dto.GearBoxResponse} "GearBox response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/gearbox/update/{id} [put]
// @Security AuthBearer
func (ch *GearBoxHandler) UpdateGearBox(ctx *gin.Context) {
	Update[dto.UpdateGearBoxRequest, dto.GearBoxResponse](ctx, ch.service.GenericUpdateGearBox)

}

// DeleteGearBox godoc
// @Summary Delete a GearBox
// @Description Delete a GearBox
// @Tags GearBox
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/gearbox/get/{id} [get]
// @Security AuthBearer
func (ch *GearBoxHandler) GetGearBoxById(ctx *gin.Context) {
	GetById[dto.GearBoxResponse](ctx, ch.service.GenericGetGearBoxById)

}

// GetGearBox godoc
// @Summary Get a GearBox
// @Description Get a GearBox
// @Tags GearBox
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "GearBox response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/gearbox/delete/{id} [delete]
// @Security AuthBearer
func (ch *GearBoxHandler) DeleteGearBox(ctx *gin.Context) {
	Delete(ctx, ch.service.GenericDeleteGearBox)

}

// GearBox godoc
// @Summary Get GearBox
// @Description Get GearBox
// @Tags GearBox
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "gearbox response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/gearbox/filter [post]
// @Security AuthBearer
func (ch *GearBoxHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.GearBoxResponse](ctx, ch.service.GenericGetByFilter)
}
