package services

import (
	"context"
	"wapi/src/config"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
)

type FileService struct {
	base *BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]
}

func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		// base: NewBaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest,dto.FileResponse](cfg),
		base: &BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]{
			Database: db.GetDB(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// create
func (cs *FileService) GenericCreateFile(ctx context.Context, req *dto.CreateFileRequest) (*dto.FileResponse, error) {
	return cs.base.Create(ctx, req)

}

// Update
func (cs *FileService) GenericUpdateFile(ctx context.Context, id int, req *dto.UpdateFileRequest) (*dto.FileResponse, error) {
	return cs.base.Update(ctx, id, req)

}

// Delete
func (cs *FileService) GenericDeleteFile(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)

}

// get by id
func (cs *FileService) GenericGetFileById(ctx context.Context, id int) (*dto.FileResponse, error) {
	return cs.base.GetById(ctx, id)

}
