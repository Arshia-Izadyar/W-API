package dto

import (
	"mime/multipart"
	"time"
)

type CreateUpdateCountryDTO struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=30"`
}

type CountryResponse struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Cities    []CityResponse    `json:"cities,omitempty"`
	Companies []CompanyResponse `json:"companies,omitempty"`
}

type CreateCityRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCityRequest struct {
	Name string `json:"name,omitempty" binding:"alpha,min=3,max=20"`
	// CountryId int    `json:"countryId,omitempty"`
}
type CityResponse struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

// file

type FileFormRequest struct {
	File *multipart.FileHeader `json:"file" form:"file" swaggerignore:"true"`
}

type UploadFileRequest struct {
	FileFormRequest
	Description string `json:"description" form:"description" binding:"required"`
}

type CreateFileRequest struct {
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MineType    string `json:"mine_type"`
}

type UpdateFileRequest struct {
	Description string `json:"description"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MineType    string `json:"mine_type"`
}

type CreateCompanyRequest struct {
	Name      string `json:"name" binding:"max=15"`
	CountryID int    `json:"countryId" binding:"required"`
}

type UpdateCompanyRequest struct {
	Name      string `json:"name,omitempty"`
	CountryID int    `json:"countryId,omitempty"`
}
type CompanyResponse struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

type CreateColorRequest struct {
	Name string `json:"name" binding:"required,max=15,min=2,alpha"`
	Hex  string `json:"hex" binding:"required,max=7,min=2"`
}

type UpdateColorRequest struct {
	Name string `json:"name,omitempty"`
	Hex  string `json:"hex,omitempty"`
}

type ColorResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Hex  string `json:"hex,omitempty"`
}

type CreateCarModelColorRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	ColorId    int `json:"colorId" binding:"required"`
}

type UpdateCarModelColorRequest struct {
	CarModelId int `json:"carModelId,omitempty"`
	ColorId    int `json:"colorId,omitempty"`
}

type CarModelColorResponse struct {
	Id    int           `json:"id"`
	Color ColorResponse `json:"color,omitempty"`
}

type CreatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle" binding:"min=4"`
	PersianYear  int       `json:"year"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

type UpdatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle" binding:"min=4"`
	PersianYear  int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type PersianYearResponse struct {
	Id           int    `json:"id"`
	PersianTitle string `json:"persianTitle" binding:"min=4"`
	PersianYear  int    `json:"year"`
}

/*

 */
