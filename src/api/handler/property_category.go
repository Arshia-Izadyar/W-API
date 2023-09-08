package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{
		service: services.NewPropertyCategoryService(cfg),
	}
}

// CreatePropertyCategory godoc
// @Summary Create Property Category
// @Description Create Property Category
// @Tags Categories
// @Accept json
// @produces json
// @Param Request body dto.CreatePropertyCategoryRequest true "Create Property Category"
// @Success 201 {object} helper.Response "Category response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property-category/create [post]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) CreatePropertyCategory(ctx *gin.Context) {
	Create[dto.CreatePropertyCategoryRequest, dto.PropertyCategoryResponse](ctx, pch.service.CreatePropertyCategoryService)
}

// UpdatePropertyCategory godoc
// @Summary Update Property Category
// @Description Update Property Category
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyCategoryRequest true "update a property-category"
// @Success 200 {object} helper.Response "property-category response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property-category/update/{id} [put]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) UpdatePropertyCategory(ctx *gin.Context) {
	Update[dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse](ctx, pch.service.UpdatePropertyCategoryService)
}

// GetPropertyCategory godoc
// @Summary Get Property Category
// @Description Get Property Category
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "property-category response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property-category/get/{id} [get]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) GetPropertyCategory(ctx *gin.Context) {
	GetById[dto.PropertyCategoryResponse](ctx, pch.service.GetPropertyCategoryServiceById)
}

// DeletePropertyCategory godoc
// @Summary Delete a property-category
// @Description Delete a property-category
// @Tags Categories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 204 {object} helper.Response "property-category response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property-category/delete/{id} [delete]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) DeletePropertyCategory(ctx *gin.Context) {
	Delete(ctx, pch.service.DeletePropertyCategoryService)
}

// GetPropertyByFilter godoc
// @Summary Get Property Filter
// @Description Get Property By Filter
// @Tags Categories
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get property-category By Filter"
// @Success 200 {object} helper.Response "property-category response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/property-category/filter [post]
// @Security AuthBearer
func (pch *PropertyCategoryHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.PropertyCategoryResponse](ctx, pch.service.GetPropertyCategoryServiceByFilter)
}
