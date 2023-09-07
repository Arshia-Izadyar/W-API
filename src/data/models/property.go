package models

type PropertyCategory struct {
	BaseModel
	Name       string `gorm:"size:15;type:string;not null"`
	Icon       string `gorm:"size:250;type:string;null"`
	Properties []Property
}

type Property struct {
	BaseModel
	Description string           `gorm:"size:550;type:string;null"`
	DataType    string           `gorm:"size:15;type:string;not null"`
	Unit        string           `gorm:"size:15;type:string;not null"`
	Name        string           `gorm:"size:15;type:string;not null"`
	Icon        string           `gorm:"size:250;type:string;null"`
	Category    PropertyCategory `gorm:"foreignKey:CategoryId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	CategoryId  int
}
