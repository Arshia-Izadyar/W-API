package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type ColorService struct {
	base *BaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: &BaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// create
func (cs *ColorService) GenericCreateColor(ctx context.Context, req *dto.CreateColorRequest) (*dto.ColorResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *ColorService) GenericUpdateColor(ctx context.Context, id int, req *dto.UpdateColorRequest) (*dto.ColorResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *ColorService) GenericDeleteColor(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *ColorService) GenericGetColorById(ctx context.Context, id int) (*dto.ColorResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *ColorService) GenericGetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.ColorResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
