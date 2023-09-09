package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/constants"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: &BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{name: "User"},
			},
		},
	}
}

// create
func (cs *CarModelCommentService) CreateCarModelComment(ctx context.Context, req *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	req.UserId = int(ctx.Value(constants.UserIdKey).(float64))
	return cs.base.Create(ctx, req)

}

// Update
func (cs *CarModelCommentService) UpdateCarModelComment(ctx context.Context, id int, req *dto.UpdateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *CarModelCommentService) DeleteCarModelComment(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *CarModelCommentService) GetCarModelCommentById(ctx context.Context, id int) (*dto.CarModelCommentResponse, error) {
	return cs.base.GetById(ctx, id)

}

func (cs *CarModelCommentService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CarModelCommentResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
