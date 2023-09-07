package dto

type CreateUpdateCountryDTO struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=30"`
}

type CountryResponse struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Cities []CityResponse `json:"cities,omitempty"`
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
