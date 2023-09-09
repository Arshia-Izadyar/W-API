package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type CarModelFileHandler struct {
	service *services.CarModelFileService
}

func NewCarModelFileHandler(cfg *config.Config) *CarModelFileHandler {
	return &CarModelFileHandler{
		service: services.NewCarModelFileService(cfg),
	}
}

// CreateCarModelFile godoc
// @Summary Create a CarModelFile
// @Description Create a CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelFileRequest true "Create a CarModelFile"
// @Success 201 {object} helper.Response{result=dto.CarModelFileResponse} "CarModelFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-file/create [post]
// @Security AuthBearer
func (ch *CarModelFileHandler) CreateCarModelFile(ctx *gin.Context) {
	Create[dto.CreateCarModelFileRequest, dto.CarModelFileResponse](ctx, ch.service.CreateCarModelFile)
}

// UpdateCarModelFile godoc
// @Summary Update a CarModelFile
// @Description Update a CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelFileRequest true "Update a CarModelFile"
// @Success 200 {object} helper.Response{result=dto.CarModelFileResponse} "CarModelFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-file/update/{id} [put]
// @Security AuthBearer
func (ch *CarModelFileHandler) UpdateCarModelFile(ctx *gin.Context) {
	Update[dto.UpdateCarModelFileRequest, dto.CarModelFileResponse](ctx, ch.service.UpdateCarModelFile)

}

// DeleteCarModelFile godoc
// @Summary Delete a CarModelFile
// @Description Delete a CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-file/get/{id} [get]
// @Security AuthBearer
func (ch *CarModelFileHandler) GetCarModelFileById(ctx *gin.Context) {
	GetById[dto.CarModelFileResponse](ctx, ch.service.GetCarModelFileById)

}

// GetCarModelFile godoc
// @Summary Get a CarModelFile
// @Description Get a CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "CarModelFile response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/car-file/delete/{id} [delete]
// @Security AuthBearer
func (ch *CarModelFileHandler) DeleteCarModelFile(ctx *gin.Context) {
	Delete(ctx, ch.service.DeleteCarModelFile)

}

// CarModelFile godoc
// @Summary Get CarModelFile
// @Description Get CarModelFile
// @Tags CarModelFile
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "City response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-file/filter [post]
// @Security AuthBearer
func (ch *CarModelFileHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CarModelFileResponse](ctx, ch.service.GetByFilter)
}
