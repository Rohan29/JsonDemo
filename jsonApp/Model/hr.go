package model

import (
	"fmt"
	"jsonApp/utility"
	"strconv"
	"time"
)

// HR struct with methods to implement the ValidatableEntity interface
type HR struct {
	HRID            int       `json:"hrId" bson:"hrId" validate:"required,min=1"`
	Name            string    `json:"name" bson:"name" validate:"required,regex=^[A-Za-z\s]{1,100}$"`
	Email           string    `json:"email" bson:"email" validate:"required,regex=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"`
	Phone           string    `json:"phone" bson:"phone" validate:"required,regex=^\+?\d{1,3}-\d{3}-\d{3}-\d{4}$"`
	Department      string    `json:"department" bson:"department" validate:"required,regex=^[A-Za-z\s]{1,100}$"`
	CreatedAt       time.Time `json:"createdAt" bson:"-" validate:"required"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"-" validate:"required"`
	Epoch_CreatedAt int64     `bson:"epoch_createdAt"`
	Epoch_UpdatedAt int64     `bson:"epoch_updatedAt"`
	CompanyID       int       `json:"companyId" bson:"companyId" validate:"required,min=1"`
}

func (hr *HR) ValidateHRFields() (*HR, error) {
	createdAt, updatedAt, err := utility.ParseTimestamps(hr.CreatedAt.String(), hr.UpdatedAt.String())
	if err != nil {
		return nil, err
	}

	hr.CreatedAt = createdAt
	hr.UpdatedAt = updatedAt

	if err := utility.ValidateStruct(hr); err != nil {
		fmt.Println("HR validation failed")
		return nil, err
	} else {
		fmt.Println("HR Validation successful!")
	}
	hr.Epoch_CreatedAt = createdAt.Unix()
	hr.Epoch_UpdatedAt = updatedAt.Unix()

	fmt.Println("Data validated for HRID", strconv.Itoa(hr.HRID))
	return hr, nil
}
