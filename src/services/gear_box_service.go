package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type GearBoxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearBoxRequest, dto.UpdateGearBoxRequest, dto.GearBoxResponse]
}

func NewGearBoxService(cfg *config.Config) *GearBoxService {
	return &GearBoxService{
		base: &BaseService[models.Gearbox, dto.CreateGearBoxRequest, dto.UpdateGearBoxRequest, dto.GearBoxResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// create
func (cs *GearBoxService) GenericCreateGearBox(ctx context.Context, req *dto.CreateGearBoxRequest) (*dto.GearBoxResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *GearBoxService) GenericUpdateGearBox(ctx context.Context, id int, req *dto.UpdateGearBoxRequest) (*dto.GearBoxResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *GearBoxService) GenericDeleteGearBox(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *GearBoxService) GenericGetGearBoxById(ctx context.Context, id int) (*dto.GearBoxResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *GearBoxService) GenericGetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.GearBoxResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
