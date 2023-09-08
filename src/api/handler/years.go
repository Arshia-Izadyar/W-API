package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type PersianYearHandler struct {
	service *services.PersianYearService
}

func NewPersianYearHandler(cfg *config.Config) *PersianYearHandler {
	return &PersianYearHandler{
		service: services.NewPersianYearService(cfg),
	}
}

// CreatePersianYear godoc
// @Summary Create a PersianYear
// @Description Create a PersianYear
// @Tags PersianYear
// @Accept json
// @produces json
// @Param Request body dto.CreatePersianYearRequest true "Create a PersianYear"
// @Success 201 {object} helper.Response{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/year/create [post]
// @Security AuthBearer
func (ch *PersianYearHandler) CreatePersianYear(ctx *gin.Context) {
	Create[dto.CreatePersianYearRequest, dto.PersianYearResponse](ctx, ch.service.CreatePersianYear)
}

// UpdatePersianYear godoc
// @Summary Update a PersianYear
// @Description Update a PersianYear
// @Tags PersianYear
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePersianYearRequest true "Update a PersianYear"
// @Success 200 {object} helper.Response{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/year/update/{id} [put]
// @Security AuthBearer
func (ch *PersianYearHandler) UpdatePersianYear(ctx *gin.Context) {
	Update[dto.UpdatePersianYearRequest, dto.PersianYearResponse](ctx, ch.service.UpdatePersianYear)

}

// DeletePersianYear godoc
// @Summary Delete a PersianYear
// @Description Delete a PersianYear
// @Tags PersianYear
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/year/get/{id} [get]
// @Security AuthBearer
func (ch *PersianYearHandler) GetPersianYearById(ctx *gin.Context) {
	GetById[dto.PersianYearResponse](ctx, ch.service.GetPersianYearById)

}

// GetPersianYear godoc
// @Summary Get a PersianYear
// @Description Get a PersianYear
// @Tags PersianYear
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "PersianYear response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/year/delete/{id} [delete]
// @Security AuthBearer
func (ch *PersianYearHandler) DeletePersianYear(ctx *gin.Context) {
	Delete(ctx, ch.service.DeletePersianYear)

}

// PersianYear godoc
// @Summary Get PersianYear
// @Description Get PersianYear
// @Tags PersianYear
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/year/filter [post]
// @Security AuthBearer
func (ch *PersianYearHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.PersianYearResponse](ctx, ch.service.GetByFilter)
}
