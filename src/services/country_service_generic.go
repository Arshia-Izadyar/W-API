package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type GenericCountryService struct {
	base *BaseService[models.Country, dto.CreateUpdateCountryDTO, dto.CreateUpdateCountryDTO, dto.CountryResponse]
}

func NewGenericCountryService(cfg *config.Config) *GenericCountryService {
	return &GenericCountryService{
		base: &BaseService[models.Country, dto.CreateUpdateCountryDTO, dto.CreateUpdateCountryDTO, dto.CountryResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{name: "Cities"}, {name: "Companies"}}, // []preload{{string: "Cities.Region"}} Chain preload
		},
	}
}

// create
func (cs *GenericCountryService) GenericCreateCountry(ctx context.Context, req *dto.CreateUpdateCountryDTO) (*dto.CountryResponse, error) {
	res, err := cs.base.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Update
func (cs *GenericCountryService) GenericUpdateCountry(ctx context.Context, id int, req *dto.CreateUpdateCountryDTO) (*dto.CountryResponse, error) {
	res, err := cs.base.Update(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete
func (cs *GenericCountryService) GenericDeleteCountry(ctx context.Context, id int) error {
	err := cs.base.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil

}

// get by id
func (cs *GenericCountryService) GenericGetCountryById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *GenericCountryService) GenericGetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CountryResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
