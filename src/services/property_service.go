package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: &BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{name: "Category"}},
		},
	}
}

// create
func (pc *PropertyService) CreateProperty(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {
	return pc.base.Create(ctx, req)
}

// func (cs *PropertyService) CreateProperty(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {
// 	property := &models.Property{

// 		Description: req.Description,
// 		DataType:    req.DataType,
// 		Unit:        req.Unit,
// 		Name:        req.Name,
// 		Icon:        req.Icon,
// 		CategoryId:  req.CategoryId,
// 	}
// 	model, _ := common.TypeConverter[models.Property](req)
// 	fmt.Println(model)

// 	tx := cs.base.Database.WithContext(ctx).Begin()
// 	err := tx.Model(&models.Property{}).Create(&property).Error
// 	if err != nil {
// 		tx.Rollback()
// 		cs.base.Logger.Error(err, logging.Postgres, logging.Insert, "cant add country", nil)
// 		return nil, err
// 	}
// 	tx.Commit()
// 	c := dto.PropertyResponse{
// 		Id:          property.Id,
// 		Name:        property.Name,
// 		Icon:        property.Icon,
// 		Description: property.Description,
// 		DataType:    property.DataType,
// 		Unit:        property.Unit,
// 	}
// 	return &c, nil

// }

// update
func (pc *PropertyService) UpdateProperty(ctx context.Context, id int, req *dto.UpdatePropertyRequest) (*dto.PropertyResponse, error) {
	return pc.base.Update(ctx, id, req)
}

// delete
func (pc *PropertyService) DeleteProperty(ctx context.Context, id int) error {
	return pc.base.Delete(ctx, id)
}

// get
func (pc *PropertyService) GetPropertyById(ctx context.Context, id int) (*dto.PropertyResponse, error) {
	return pc.base.GetById(ctx, id)
}

// GetByFilter
func (pc *PropertyService) GetPropertyByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.PropertyResponse], error) {
	return pc.base.GetByFilter(ctx, req)
}
