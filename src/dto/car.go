package dto

import "time"

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
	Id             int                     `json:"id"`
	Name           string                  `json:"name"`
	Company        CompanyResponse         `json:"company"`
	Gearbox        GearBoxResponse         `json:"gearbox"`
	CarType        CarTypeResponse         `json:"carType"`
	CarModelColors []CarModelColorResponse `json:"carModelColors,omitempty"`
	CarModelYears  []CarModelYearResponse  `json:"carModelYears,omitempty"`
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
