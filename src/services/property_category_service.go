package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type PropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	return &PropertyCategoryService{
		base: &BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{name: "Properties"}},
		},
	}
}

// create
func (pc *PropertyCategoryService) CreatePropertyCategoryService(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return pc.base.Create(ctx, req)
}

// update
func (pc *PropertyCategoryService) UpdatePropertyCategoryService(ctx context.Context, id int, req *dto.UpdatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return pc.base.Update(ctx, id, req)
}

// delete
func (pc *PropertyCategoryService) DeletePropertyCategoryService(ctx context.Context, id int) error {
	return pc.base.Delete(ctx, id)
}

// get
func (pc *PropertyCategoryService) GetPropertyCategoryServiceById(ctx context.Context, id int) (*dto.PropertyCategoryResponse, error) {
	return pc.base.GetById(ctx, id)
}

// GetByFilter
func (pc *PropertyCategoryService) GetPropertyCategoryServiceByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.PropertyCategoryResponse], error) {
	return pc.base.GetByFilter(ctx, req)
}
