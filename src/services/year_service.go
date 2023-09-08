package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// create
func (cs *PersianYearService) CreatePersianYear(ctx context.Context, req *dto.CreatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *PersianYearService) UpdatePersianYear(ctx context.Context, id int, req *dto.UpdatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *PersianYearService) DeletePersianYear(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *PersianYearService) GetPersianYearById(ctx context.Context, id int) (*dto.PersianYearResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *PersianYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.PersianYearResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
