package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{name: "Company.Country"},
				{name: "CarType"},
				{name: "Gearbox"},
				{name: "CarModelColors.Color"},
				{name: "CarModelYears.PersianYear"},
				{name: "CarModelYears.PersianYear.CarModelPriceHistories"},
				{name: "CarModelProperties.Property.Category"},
				{name: "CarModelFiles.File"},
				{name: "CarModelComments.User"},
			},
		},
	}
}

// create
func (cs *CarModelService) GenericCreateCarModel(ctx context.Context, req *dto.CreateCarModelRequest) (*dto.CarModelResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CarModelService) GenericUpdateCarModel(ctx context.Context, id int, req *dto.UpdateCarModelRequest) (*dto.CarModelResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CarModelService) GenericDeleteCarModel(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CarModelService) GenericGetCarModelById(ctx context.Context, id int) (*dto.CarModelResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CarModelService) GenericGetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CarModelResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
