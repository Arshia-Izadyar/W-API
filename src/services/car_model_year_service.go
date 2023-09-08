package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: &BaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{name: "PersianYear"},
			},
		},
	}
}

// create
func (cs *CarModelYearService) CreateCarModelYear(ctx context.Context, req *dto.CreateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CarModelYearService) UpdateCarModelYear(ctx context.Context, id int, req *dto.UpdateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CarModelYearService) DeleteCarModelYear(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CarModelYearService) GetCarModelYearById(ctx context.Context, id int) (*dto.CarModelYearResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CarModelYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CarModelYearResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
