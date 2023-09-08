package dto

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
