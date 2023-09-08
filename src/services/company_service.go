package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{name: "Country"},
			},
		},
	}
}

// create
func (cs *CompanyService) GenericCreateCompany(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CompanyService) GenericUpdateCompany(ctx context.Context, id int, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CompanyService) GenericDeleteCompany(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CompanyService) GenericGetCompanyById(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CompanyService) GenericGetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CompanyResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
