package dto

import (
	"time"
)

type CreateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=50"`
}

type UpdateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=50"`
}

type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateGearBoxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=50"`
}
type UpdateGearBoxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=50"`
}

type GearBoxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required,max=15"`
	CompanyId int    `json:"companyId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
}

type UpdateCarModelRequest struct {
	Name      string `json:"name" binding:"required,max=15"`
	CompanyId int    `json:"companyId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
}

type CarModelResponse struct {
	Id                 int                        `json:"id"`
	Name               string                     `json:"name"`
	Company            CompanyResponse            `json:"company"`
	Gearbox            GearBoxResponse            `json:"gearbox"`
	CarType            CarTypeResponse            `json:"carType"`
	CarModelColors     []CarModelColorResponse    `json:"carModelColors,omitempty"`
	CarModelYears      []CarModelYearResponse     `json:"carModelYears,omitempty"`
	CarModelFiles      []CarModelFileResponse     `json:"carModelFiles"`
	CarModelProperties []CarModelPropertyResponse `json:"carModelProperties"`
	CarModelComments   []CarModelCommentResponse  `json:"carModelComments"`
}

type CreateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId" binding:"required"`
	PersianYearId int `json:"persianYearId" binding:"required"`
}

type UpdateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId"`
	PersianYearId int `json:"persianYearId"`
}

type CarModelYearResponse struct {
	CarModelId    int                     `json:"carModelId"`
	Id            int                     `json:"id"`
	PersianYear   PersianYearResponse     `json:"persianYear"`
	CarModelPrice []CarModelPriceResponse `json:"carModelPrice"`
}

type CreateCarModelPriceRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	Price          float64   `json:"price"`
	PriceAt        time.Time `json:"priceAt"`
}

type UpdateCarModelPriceRequest struct {
	Price   float64   `json:"price,omitempty"`
	PriceAt time.Time `json:"priceAt,omitempty"`
}

type CarModelPriceResponse struct {
	Id             int       `json:"id"`
	CarModelYearId int       `json:"carModelYearId"`
	Price          float64   `json:"price"`
	PriceAt        time.Time `json:"priceAt"`
}

type CreateCarModelFileRequest struct {
	CarModelId  int  `json:"carModelId" binding:"required"`
	FileId      int  `json:"fileId" binding:"required"`
	IsMainImage bool `json:"isMainImage"`
}

type UpdateCarModelFileRequest struct {
	IsMainImage bool `json:"isMainImage"`
}

type CarModelFileResponse struct {
	Id          int          `json:"id"`
	File        FileResponse `json:"file"`
	IsMainImage bool         `json:"isMainImage"`
}

type CreateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId"`
	Value      string `json:"value"`
	PropertyId int    `json:"propertyId"`
}

type UpdateCarModelPropertyRequest struct {
	Value string `json:"value"`
}

type CarModelPropertyResponse struct {
	Id         int              `json:"id"`
	CarModelId int              `json:"carModelId"`
	Value      string           `json:"value"`
	Property   PropertyResponse `json:"property"`
}

type CreateCarModelCommentRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	UserId     int    `json:"userId"`
	Message    string `json:"message" binding:"required,max=1000"`
}

type UpdateCarModelCommentRequest struct {
	Message string `json:"message"`
}

type CarModelCommentResponse struct {
	Id         int          `json:"id"`
	CarModelId int          `json:"carModelId"`
	User       UserResponse `json:"user"`
	Message    string       `json:"message"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
}
