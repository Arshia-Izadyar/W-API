package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: &BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{name: "Color"},
			},
		},
	}
}

// create
func (cs *CarModelColorService) GenericCreateCarModelColor(ctx context.Context, req *dto.CreateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CarModelColorService) GenericUpdateCarModelColor(ctx context.Context, id int, req *dto.UpdateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CarModelColorService) GenericDeleteCarModelColor(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CarModelColorService) GenericGetCarModelColorById(ctx context.Context, id int) (*dto.CarModelColorResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CarModelColorService) GenericGetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CarModelColorResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
