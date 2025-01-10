package model

import (
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

// HR struct with regex validation directly in the tags
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

type Employee struct {
	EmployeeID      int       `json:"employeeId" bson:"employeeId" validate:"required,min=1"`
	Name            string    `json:"name" bson:"name" validate:"required,regex=^[A-Za-z\s]{1,100}$"`
	Email           string    `json:"email" bson:"email" validate:"required,regex=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"`
	Phone           string    `json:"phone" bson:"phone" validate:"required,regex=^\+?\d{1,3}-\d{3}-\d{3}-\d{4}$"`
	Position        string    `json:"position" bson:"position" validate:"required,regex=^[A-Za-z\s]{1,100}$"`
	Salary          float64   `json:"salary" bson:"salary" validate:"required,gt=0"`
	HireDate        string    `json:"hireDate" bson:"hireDate" validate:"required,regex=^\\d{4}-\\d{2}-\\d{2}$"`
	CreatedAt       time.Time `json:"createdAt" bson:"-" validate:"required"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"-" validate:"required"`
	Epoch_CreatedAt int64     `bson:"epoch_createdAt"`
	Epoch_UpdatedAt int64     `bson:"epoch_updatedAt"`
	CompanyID       int       `json:"companyId" bson:"companyId" validate:"required,min=1"`
}

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

type JsonCompaniesData struct {
	Companies []struct {
		CompanyID int        `json:"companyId"`
		Name      string     `json:"name"`
		Address   string     `json:"address"`
		Industry  string     `json:"industry"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt time.Time  `json:"updatedAt"`
		Admins    []Admin    `json:"admins"`
		HR        []HR       `json:"hr"`
		Employees []Employee `json:"employees"`
	}
}

type MyError struct {
	Message string
}
