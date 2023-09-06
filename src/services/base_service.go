package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"
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

type preload struct {
	string
}

//				 -dto	 +dto
//	     model -create  +update  response
type BaseService[T any, TC any, TU any, TR any] struct {
	Database *gorm.DB
	Logger   logging.Logger
	Preloads []preload
}

func NewBaseService[T any, TC any, TU any, TR any](cfg *config.Config) *BaseService[T, TC, TU, TR] {
	return &BaseService[T, TC, TU, TR]{
		Database: db.GetDB(),
		Logger:   logging.NewLogger(cfg),
		Preloads: []preload{{string: "Cities"}}, // []preload{{string: "Cities.Region"}} Chain preload
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
	bm, err := common.TypeConverter[models.BaseModel](model)
	if err != nil {
		bs.Logger.Error(err, logging.Postgres, logging.Select, "cant select country", nil)
		return nil, err
	}
	return bs.GetById(ctx, bm.Id)
}

func (bs *BaseService[T, TC, TU, TR]) Update(ctx context.Context, id int, req *TC) (*TR, error) {
	updateMap, err := common.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}
	snakeMap["modified_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}
	snakeMap["modified_at"] = &sql.NullTime{Valid: true, Time: time.Now().UTC()}

	model := new(T)
	tx := bs.Database.WithContext(ctx).Begin()

	err = tx.Model(model).Where("id = ? AND deleted_by IS NULL ", id).Updates(snakeMap).Error
	if err != nil {
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
		"deleted_at": &sql.NullTime{Valid: true, Time: time.Now().UTC()},
		"deleted_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))},
	}

	if ctx.Value(constants.UserIdKey) == nil {
		tx.Rollback()

		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	if cnt := tx.Model(model).Where("id = ?", id).Updates(deleteMap).RowsAffected; cnt == 0 {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: "record not found"}
	}
	tx.Commit()
	return nil
}

func (bs *BaseService[T, TC, TU, TR]) GetById(ctx context.Context, id int) (*TR, error) {
	model := new(T)
	db := Preload(bs.Database, bs.Preloads)
	err := db.Model(&models.Country{}).
		Where("id = ? AND deleted_by IS NULL", id).
		First(&model).
		Error
	if err != nil {
		return nil, err
	}

	return common.TypeConverter[TR](model)

}

// get query / get sort / paginator / preload

func (bs *BaseService[T, TC, TU, TR]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[TR], error) {
	return Paginate[T, TR](req, bs.Preloads, bs.Database)
}

func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	if filter.Filter != nil {
		for name, filter := range filter.Filter {
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
					query = append(query, fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From))
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

func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)
	}
	return db
}

func NewPageList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PageList[T] {
	pl := &dto.PageList[T]{
		PageNumber: pageNumber,
		TotalRows:  count,
		Items:      items,
	}
	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasNextPage = pl.PageNumber > 1
	return pl

}

func Paginate1[T any, TR any](pagination *dto.PaginationInputWithFilter, preloads []preload, db *gorm.DB) (*dto.PageList[TR], error) {
	model := new(T)
	var items *[]T
	var rItems *[]TR
	db = Preload(db, preloads)
	query := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)

	var totalRows int64 = 0
	db.Model(model).Where(query).Count(&totalRows)
	err := db.Where(query).Offset(pagination.GetOffSet()).Limit(pagination.GetPageSize()).Order(sort).Find(&items).Error
	if err != nil {
		return nil, err
	}
	rItems, err = common.TypeConverter[[]TR](items)
	if err != nil {
		return nil, err
	}
	return NewPageList(rItems, totalRows, pagination.PageNumber, int64(pagination.PageSize)), err
}
func Paginate[T any, Tr any](pagination *dto.PaginationInputWithFilter, preloads []preload, db *gorm.DB) (*dto.PageList[Tr], error) {
	model := new(T)
	var items *[]T
	var rItems *[]Tr
	db = Preload(db, preloads)
	query := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)

	var totalRows int64 = 0

	db.
		Model(model).
		Where(query).
		Count(&totalRows)

	err := db.
		Where(query).
		Offset(pagination.GetOffSet()).
		Limit(pagination.GetPageSize()).
		Order(sort).
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}
	rItems, err = common.TypeConverter[[]Tr](items)
	if err != nil {
		return nil, err
	}
	return NewPageList(rItems, totalRows, pagination.PageNumber, int64(pagination.PageSize)), err
}
