package model

import (
	"fmt"
	"jsonApp/utility"
	"strconv"
	"time"
)

type Admin struct {
	AdminID         int       `json:"adminId" bson:"adminId" validate:"required,min=1"`
	Name            string    `json:"name" bson:"name" validate:"required,regex=^[A-Za-z\s]{1,100}$"`
	Email           string    `json:"email" bson:"email" validate:"required,regex=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"`
	Phone           string    `json:"phone" bson:"phone" validate:"required,regex=^\+?\d{1,3}-\d{3}-\d{3}-\d{4}$"`
	CreatedAt       time.Time `json:"createdAt" bson:"-" validate:"required"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"-" validate:"required"`
	Epoch_CreatedAt int64     `bson:"epoch_createdAt"`
	Epoch_UpdatedAt int64     `bson:"epoch_updatedAt"`
	CompanyID       int       `json:"companyId" bson:"companyId" validate:"required,min=1"`
}

// validateAdminFields validates an admin's fields
func (admin *Admin) ValidateAdminFields() (*Admin, error) {
	createdAt, updatedAt, err := utility.ParseTimestamps(admin.CreatedAt.String(), admin.UpdatedAt.String())
	if err != nil {
		return nil, err
	}

	admin.CreatedAt = createdAt
	admin.UpdatedAt = updatedAt
	if err := utility.ValidateStruct(admin); err != nil {
		return nil, err
	} else {
		fmt.Println("Admin Validation successful!")
	}
	admin.Epoch_CreatedAt = createdAt.Unix()
	admin.Epoch_UpdatedAt = updatedAt.Unix()

	fmt.Println("Data validated for Adminid", strconv.Itoa(admin.AdminID))

	return admin, nil
}
