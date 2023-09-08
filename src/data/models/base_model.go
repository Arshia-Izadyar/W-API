package models

import (
	"database/sql"
	"time"
	"wapi/src/constants"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id         int           `gorm:"primaryKey"`
	CreatedAt  time.Time     `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime  `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime  `gorm:"type:TIMESTAMP with time zone;null"`
	CreatedBy  int           `gorm:"not null"`
	ModifiedBy sql.NullInt64 `gorm:"null"`
	DeletedBy  sql.NullInt64 `gorm:"null"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value(constants.UserIdKey)
	var userId int = -1
	if value != nil {
		userId = int(value.(float64))
	}
	b.CreatedAt = time.Now()
	b.CreatedBy = userId
	return
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value(constants.UserIdKey)
	var userId = &sql.NullInt64{
		Valid: false,
	}
	if value != nil {
		userId = &sql.NullInt64{
			Valid: true,
			Int64: int64(value.(float64)),
		}
	}
	b.ModifiedBy = *userId
	b.ModifiedAt = sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}
	return
}
