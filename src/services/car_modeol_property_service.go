package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CarModelPropertyService struct {
	base *BaseService[models.CarModelProperty, dto.CreateCarModelPropertyRequest, dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]
}

func NewCarModelPropertyService(cfg *config.Config) *CarModelPropertyService {
	return &CarModelPropertyService{
		base: &BaseService[models.CarModelProperty, dto.CreateCarModelPropertyRequest, dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{name: "Property.Category"},
			},
		},
	}
}

// create
func (cs *CarModelPropertyService) CreateCarModelProperty(ctx context.Context, req *dto.CreateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return cs.base.Create(ctx, req)

}

// car_model_property
// Update
func (cs *CarModelPropertyService) UpdateCarModelProperty(ctx context.Context, id int, req *dto.UpdateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CarModelPropertyService) DeleteCarModelProperty(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CarModelPropertyService) GetCarModelPropertyById(ctx context.Context, id int) (*dto.CarModelPropertyResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CarModelPropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CarModelPropertyResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
