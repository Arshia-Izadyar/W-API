package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: &BaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{name: "Country"},
			},
		},
	}
}

// create
func (cs *CityService) GenericCreateCity(ctx context.Context, req *dto.CreateCityRequest) (*dto.CityResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CityService) GenericUpdateCity(ctx context.Context, id int, req *dto.UpdateCityRequest) (*dto.CityResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CityService) GenericDeleteCity(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CityService) GenericGetCityById(ctx context.Context, id int) (*dto.CityResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CityService) GenericGetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CityResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
