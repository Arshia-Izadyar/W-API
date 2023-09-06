package services

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
	"wapi/src/common"
	"wapi/src/config"
	"wapi/src/constants"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"
	"wapi/src/pkg/service_errors"

	"gorm.io/gorm"
)

//						 -dto	 +dto
//			     model -create  +update  response
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


// get query / get sort / paginator / preload

func getQuery[T any](filter *dto.DynamicFilter) string{
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	if filter.Filter != nil {
		for name, filter := range filter.Filter{
			fld, ok := typeT.FieldByName(name)
			if ok {
				fld.Name = common.ToSnakeCase(fld.Name)
				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "notContains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From))
				case "endsWith":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s'", fld.Name, filter.From))
				case "equals":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEquals":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lessThan":
					query = append(query, fmt.Sprintf("%s < %s", fld.Name, filter.From))
				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > '%s'", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
					}
				
				}
			}
		}
	}
	return strings.Join(query, " AND ")
}

func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range *filter.Sort {
			fld, ok := typeT.FieldByName(tp.ColID)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}