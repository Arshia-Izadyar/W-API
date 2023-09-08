package handler

import (
	"wapi/src/config"
	"wapi/src/dto"
	"wapi/src/services"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{
		service: services.NewCompanyService(cfg),
	}
}

// CreateCompany godoc
// @Summary Create a Company
// @Description Create a Company
// @Tags Company
// @Accept json
// @produces json
// @Param Request body dto.CreateCompanyRequest true "Create a Company"
// @Success 201 {object} helper.Response{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/car-type/create [post]
// @Security AuthBearer
func (ch *CompanyHandler) CreateCompany(ctx *gin.Context) {
	Create[dto.CreateCompanyRequest, dto.CompanyResponse](ctx, ch.service.GenericCreateCompany)
}

// UpdateCompany godoc
// @Summary Update a Company
// @Description Update a Company
// @Tags Company
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCompanyRequest true "Update a Company"
// @Success 200 {object} helper.Response{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/company/update/{id} [put]
// @Security AuthBearer
func (ch *CompanyHandler) UpdateCompany(ctx *gin.Context) {
	Update[dto.UpdateCompanyRequest, dto.CompanyResponse](ctx, ch.service.GenericUpdateCompany)

}

// DeleteCompany godoc
// @Summary Delete a Company
// @Description Delete a Company
// @Tags Company
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/company/get/{id} [get]
// @Security AuthBearer
func (ch *CompanyHandler) GetCompanyById(ctx *gin.Context) {
	GetById[dto.CompanyResponse](ctx, ch.service.GenericGetCompanyById)

}

// GetCompany godoc
// @Summary Get a Company
// @Description Get a Company
// @Tags Company
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.Response "Company response"
// @Failure 400 {object} helper.Response "Bad request"
// @Failure 404 {object} helper.Response "Not found"
// @Router /v1/company/delete/{id} [delete]
// @Security AuthBearer
func (ch *CompanyHandler) DeleteCompany(ctx *gin.Context) {
	Delete(ctx, ch.service.GenericDeleteCompany)

}

// company godoc
// @Summary Get company
// @Description Get company
// @Tags Company
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.Response "company response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/company/filter [post]
// @Security AuthBearer
func (ch *CompanyHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CompanyResponse](ctx, ch.service.GenericGetByFilter)
}
