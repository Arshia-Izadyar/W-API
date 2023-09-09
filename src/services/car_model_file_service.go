package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CarModelFileService struct {
	base *BaseService[models.CarModelFile, dto.CreateCarModelFileRequest, dto.UpdateCarModelFileRequest, dto.CarModelFileResponse]
}

func NewCarModelFileService(cfg *config.Config) *CarModelFileService {
	return &CarModelFileService{
		base: &BaseService[models.CarModelFile, dto.CreateCarModelFileRequest, dto.UpdateCarModelFileRequest, dto.CarModelFileResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{name: "CarModel"},
				{name: "File"},
				{name: "CarModel.Company"},
				{name: "CarModel.Company.Country"},
				{name: "CarModel.CarType"},
				{name: "CarModel.Gearbox"},
			},
		},
	}
}

// create
func (cs *CarModelFileService) CreateCarModelFile(ctx context.Context, req *dto.CreateCarModelFileRequest) (*dto.CarModelFileResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CarModelFileService) UpdateCarModelFile(ctx context.Context, id int, req *dto.UpdateCarModelFileRequest) (*dto.CarModelFileResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CarModelFileService) DeleteCarModelFile(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CarModelFileService) GetCarModelFileById(ctx context.Context, id int) (*dto.CarModelFileResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CarModelFileService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CarModelFileResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
