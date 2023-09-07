package models

type Gearbox struct {
	BaseModel
	Name string `gorm:"size:15;type:string;not null"`
}

type CarType struct {
	BaseModel
	Name string `gorm:"size:50;type:string;not null"`
}

type Company struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null"`
	CountryID int
	Country   Country //`gorm:"foreignKey:CountryID"`
}
