package dto

type CreateUpdateCountryDTO struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=30"`
}

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}