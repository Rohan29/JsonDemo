package model

import (
	"fmt"
	"jsonApp/utility"
	"time"
)

// Company struct with methods to implement the ValidatableEntity interface
type Company struct {
	CompanyID       int       `json:"companyId" bson:"companyId" validate:"required,min=1"`
	Name            string    `json:"name" bson:"name" validate:"required,regex=^[A-Za-z0-9\s\.,&\-]{1,100}$"`
	Address         string    `json:"address" bson:"address" validate:"required,regex=^[A-Za-z0-9\s,.\-]{1,200}$"`
	Industry        string    `json:"industry" bson:"industry" validate:"required,regex=^[A-Za-z\s]{1,50}$"`
	CreatedAt       time.Time `json:"createdAt" bson:"-" validate:"required"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"-" validate:"required"`
	Epoch_CreatedAt int64     `bson:"epoch_createdAt"`
	Epoch_UpdatedAt int64     `bson:"epoch_updatedAt"`
}

func (company *Company) ValidateCompanyData() (*Company, error) {
	CreatedAt, UpdatedAt, err := utility.ParseTimestamps(company.CreatedAt.String(), company.UpdatedAt.String())
	if err != nil {
		return nil, err
	}
	company.CreatedAt = CreatedAt
	company.UpdatedAt = UpdatedAt

	if err := utility.ValidateStruct(company); err != nil {
		fmt.Println("Company validation failed")
		return nil, err
	} else {
		fmt.Println("Company validation successful!")
	}
	company.Epoch_CreatedAt = CreatedAt.Unix()
	company.Epoch_UpdatedAt = UpdatedAt.Unix()
	fmt.Printf("Company %v \n", company)
	return company, nil
}
