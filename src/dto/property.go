package dto

type CreatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
	Icon string `json:"icon" binding:"min=1,max=250"`
}

type UpdatePropertyCategoryRequest struct {
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
}

type PropertyCategoryResponse struct {
	Id         int                `json:"id"`
	Name       string             `json:"name"`
	Icon       string             `json:"icon"`
	Properties []PropertyResponse `json:"properties,omitempty"`
}

type CreatePropertyRequest struct {
	Name        string `json:"name" binding:"required,alpha,min=3,max=15"`
	CategoryId  int    `json:"categoryId" binding:"required"`
	Icon        string `json:"icon" binding:"max=250"`
	Description string `json:"description" binding:"max=550"`
	DataType    string `json:"data_type" binding:"max=15"`
	Unit        string `json:"unit" binding:"max=15"`
}

type UpdatePropertyRequest struct {
	Name        string `json:"name,omitempty" `
	CategoryId  int    `json:"categoryId,omitempty"`
	Icon        string `json:"icon,omitempty" binding:"max=250"`
	Description string `json:"description,omitempty" binding:"max=550"`
	DataType    string `json:"data_type,omitempty" binding:"max=15"`
	Unit        string `json:"unit,omitempty" binding:"max=15"`
}

type PropertyResponse struct {
	Id          int                      `json:"id"`
	Name        string                   `json:"name,omitempty" `
	Category    PropertyCategoryResponse `json:"category,omitempty"`
	Icon        string                   `json:"icon,omitempty"`
	Description string                   `json:"description,omitempty"`
	DataType    string                   `json:"data_type,omitempty"`
	Unit        string                   `json:"unit,omitempty"`
}
