package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: &BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// create
func (cs *CarTypeService) GenericCreateCarType(ctx context.Context, req *dto.CreateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CarTypeService) GenericUpdateCarType(ctx context.Context, id int, req *dto.UpdateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CarTypeService) GenericDeleteCarType(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CarTypeService) GenericGetCarTypeById(ctx context.Context, id int) (*dto.CarTypeResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CarTypeService) GenericGetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CarTypeResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
