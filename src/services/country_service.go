package services

import (
	"context"
	"database/sql"
	"time"
	"wapi/src/config"
	"wapi/src/constants"
	"wapi/src/data/db"
	"wapi/src/data/models"
	"wapi/src/dto"
	"wapi/src/pkg/logging"

	"gorm.io/gorm"
)

type CountryService struct {
	db *gorm.DB
	logger logging.Logger

}

func NewCountryService(cfg *config.Config) *CountryService {
	logger := logging.NewLogger(cfg)
	db := db.GetDB()
	return &CountryService{
		db: db,
		logger: logger,
	}
}

// create 
func (cs *CountryService) CreateCountry(ctx context.Context, req *dto.CreateUpdateCountryDTO) (*dto.CountryResponse, error) {
	country := &models.Country{
		Name: req.Name,
	}
	country.CreatedBy = int(ctx.Value(constants.UserIdKey).(float64))
	country.CreatedAt = time.Now()
	tx := cs.db.WithContext(ctx).Begin()
	err := tx.Create(&country).Error
	if err != nil {
		tx.Rollback()
		cs.logger.Error(err, logging.Postgres, logging.Insert, "cant add country", nil)
		return nil, err
	}
	tx.Commit()
	c := dto.CountryResponse{
		ID:   country.Id,
		Name: country.Name,
	}
	return &c, nil

}

// Update
func (cs *CountryService) UpdateCountry(ctx context.Context, id int, req *dto.CreateUpdateCountryDTO) (*dto.CountryResponse, error) {
	updateMap := map[string]interface{}{
		"name":req.Name,
		"modified_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))},
		"modified_at": &sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	tx := cs.db.WithContext(ctx).Begin()
	err := tx.Model(&models.Country{}).Where("id = ? AND deleted_by IS NULL", id).Updates(updateMap).Error
	if err != nil {
		tx.Rollback()
		cs.logger.Error(err, logging.Postgres, logging.Insert, "cant update country", nil)
		return nil, err
	}
	country := &models.Country{}
	err = tx.Model(&models.Country{}).Where("id = ? AND deleted_by IS NULL", id).First(&country).Error
	if err != nil {
		tx.Rollback()
		cs.logger.Error(err, logging.Postgres, logging.Select, "cant Select country", nil)
		return nil, err
	}
	tx.Commit()
	c := &dto.CountryResponse{Name: country.Name, ID: country.Id}
	return c, nil
}

// Delete
func (cs *CountryService) DeleteCountry(ctx context.Context, id int) error {
	tx := cs.db.WithContext(ctx).Begin()
	deleteMap := map[string]interface{}{
		"deleted_at": &sql.NullTime{Valid:true, Time:time.Now().UTC()},
		"deleted_by": &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))},
	}
	err := tx.Model(&models.Country{}).Where("id = ?", id).Updates(deleteMap).Error
	if err != nil {
		tx.Rollback()
		cs.logger.Error(err, logging.Postgres, logging.Delete, "cant delete country", nil)
		return err
	}
	tx.Commit()
	return nil

}

// get by id
func (cs *CountryService) GetCountryById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	country := &models.Country{}

	err := cs.db.Model(&models.Country{}).
	Where("id = ? AND deleted_by IS NULL", id).
	First(&country).
	Error

	if err != nil {
		cs.logger.Error(err, logging.Postgres, logging.Get, "cant delete country", nil)
		return nil, err
	}
	c := &dto.CountryResponse{Name: country.Name, ID:country.Id}
	return c, nil
}

