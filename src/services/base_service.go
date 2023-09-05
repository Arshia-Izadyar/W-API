package services

import (
	"context"
	"database/sql"
	"time"
	"wapi/src/common"
	"wapi/src/config"
	"wapi/src/constants"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/pkg/logging"
	"wapi/src/pkg/service_errors"

	"gorm.io/gorm"
)

//			     model  create  update  response
type BaseService[T any, TC any, TU any, TR any] struct {
	Database *gorm.DB
	Logger logging.Logger
}

func NewBaseService[T any, TC any, TU any, TR any](cfg *config.Config) *BaseService[T, TC, TU, TR] {
	return &BaseService[T, TC, TU, TR]{
		Database: db.GetDB(),
		Logger: logging.NewLogger(cfg),
	}
}

func (bs *BaseService[T, TC, TU, TR]) Create(ctx context.Context, req *TC) (*TR, error) {
	model, err := common.TypeConverter[T](req)
	if err != nil {
		return nil, err
	}
	tx := bs.Database.WithContext(ctx).Begin()
	err = tx.Create(&model).Error 
	if err != nil {
		tx.Rollback()
		bs.Logger.Error(err, logging.Postgres, logging.Insert, "cant add country", nil)
		return nil, err
	}
	tx.Commit()
	return common.TypeConverter[TR](model)
}

func (bs *BaseService[T, TC, TU, TR]) Update(ctx context.Context, id int, req *TC) (*TR, error) {
	updateMap, err := common.TypeConverter[map[string]interface{}](req)
	if err != nil{
		return nil, err
	}
	(*updateMap)["modified_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}
	(*updateMap)["modified_at"] = &sql.NullTime{Valid: true, Time: time.Now().UTC()}
	model := new(T)
	tx := bs.Database.WithContext(ctx).Begin()
	err = tx.Model(model).Where("id = ? AND deleted_by IS NULL ", id).Updates(*updateMap).Error
	if err != nil{
		tx.Rollback()
		bs.Logger.Error(err, logging.Postgres, logging.Insert, "cant update country", nil)
		return nil, err
	}
	tx.Commit()
	return bs.GetById(ctx, id)
}

func (bs *BaseService[T, TC, TU, TR]) Delete(ctx context.Context, id int) error {
	tx := bs.Database.WithContext(ctx).Begin()
	model := new(T)

	deleteMap := map[string]interface{}{
		"deleted_at": &sql.NullTime{Valid:true, Time:time.Now().UTC()},
		"deleted_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))},
	}

	if ctx.Value(constants.UserIdKey) == nil {
		tx.Rollback()
		
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	if cnt:= tx.Model(model).Where("id = ?", id).Updates(deleteMap).RowsAffected; cnt == 0 {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: "record not found"}
	}
	tx.Commit()
	return nil
}

func (bs *BaseService[T, TC, TU, TR]) GetById(ctx context.Context, id int) (*TR, error) {
	model := new(T)
	err := bs.Database.Model(&models.Country{}).
	Where("id = ? AND deleted_by IS NULL", id).
	First(&model).
	Error
	if err != nil {
		return nil, err
	}

	return common.TypeConverter[TR](model)

}
