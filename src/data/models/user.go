package models

import "database/sql"

type User struct {
	BaseModel
	UserName    string         `gorm:"size:25;type:string;not null;unique"`
	FirstName   sql.NullString `gorm:"size:20;type:string;null"`
	LastName    sql.NullString `gorm:"size:20;type:string;null"`
	PhoneNumber sql.NullString `gorm:"size:11;type:string;null;unique;default:null"`
	Email       sql.NullString `gorm:"size:64;type:string;null;unique;default:null"`
	Password    string         `gorm:"size:64;type:string;not null"`
	Enable      bool           `gorm:"default:true"`
	userRoles   *[]UserRole
}

type Role struct {
	BaseModel
	Name  string `gorm:"size:20;not null;unique;type:string"`
	Users *[]UserRole
}

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId"` // constraint:onUpdate:NO ACTION;OnDelete:NO ACTION;
	Role   Role `gorm:"foreignKey:RoleId"`
	UserId int
	RoleId int
}
