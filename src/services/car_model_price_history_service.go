package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CarModelPriceService struct {
	base *BaseService[models.CarModelPrice, dto.CreateCarModelPriceRequest, dto.UpdateCarModelPriceRequest, dto.CarModelPriceResponse]
}

func NewCarModelPriceService(cfg *config.Config) *CarModelPriceService {
	return &CarModelPriceService{
		base: &BaseService[models.CarModelPrice, dto.CreateCarModelPriceRequest, dto.UpdateCarModelPriceRequest, dto.CarModelPriceResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// create
func (cs *CarModelPriceService) CreateCarModelPrice(ctx context.Context, req *dto.CreateCarModelPriceRequest) (*dto.CarModelPriceResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CarModelPriceService) UpdateCarModelPrice(ctx context.Context, id int, req *dto.UpdateCarModelPriceRequest) (*dto.CarModelPriceResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CarModelPriceService) DeleteCarModelPrice(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CarModelPriceService) GetCarModelPriceById(ctx context.Context, id int) (*dto.CarModelPriceResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CarModelPriceService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CarModelPriceResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
