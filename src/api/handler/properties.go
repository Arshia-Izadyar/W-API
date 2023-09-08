package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	service *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	return &PropertyHandler{
		service: services.NewPropertyService(cfg),
	}
}

// CreateProperty godoc
// @Summary Create Property Category
// @Description Create Property Category
// @Tags Categories
// @Accept json
// @produces json
// @Param Request body dto.CreatePropertyRequest true "Create Property Category"
// @Success 201 {object} helper.Response "Category response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property/create [post]
// @Security AuthBearer
func (pch *PropertyHandler) CreateProperty(ctx *gin.Context) {
	Create[dto.CreatePropertyRequest, dto.PropertyResponse](ctx, pch.service.CreateProperty)
	// req := dto.CreatePropertyRequest{}
	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {

	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
	// 	return
	// }
	// res, err := pch.service.CreateProperty(ctx, &req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, -1, err))
	// 	return
	// }
	// ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

// UpdateProperty godoc
// @Summary Update Property Category
// @Description Update Property Category
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyRequest true "update a property"
// @Success 200 {object} helper.Response "property response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property/update/{id} [put]
// @Security AuthBearer
func (pch *PropertyHandler) UpdateProperty(ctx *gin.Context) {
	Update[dto.UpdatePropertyRequest, dto.PropertyResponse](ctx, pch.service.UpdateProperty)
}

// GetProperty godoc
// @Summary Get Property Category
// @Description Get Property Category
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "property response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property/get/{id} [get]
// @Security AuthBearer
func (pch *PropertyHandler) GetProperty(ctx *gin.Context) {
	GetById[dto.PropertyResponse](ctx, pch.service.GetPropertyById)
}

// DeleteProperty godoc
// @Summary Delete a property
// @Description Delete a property
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 204 {object} helper.Response "property response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property/delete/{id} [delete]
// @Security AuthBearer
func (pch *PropertyHandler) DeleteProperty(ctx *gin.Context) {
	Delete(ctx, pch.service.DeleteProperty)
}

// GetPropertyByFilter godoc
// @Summary Get Property Filter
// @Description Get Property By Filter
// @Tags Categories
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get property By Filter"
// @Success 200 {object} helper.Response "property response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property/filter [post]
// @Security AuthBearer
func (pch *PropertyHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.PropertyResponse](ctx, pch.service.GetPropertyByFilter)
}
