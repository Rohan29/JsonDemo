package model

import (
	"fmt"
	"jsonApp/utility"
	"strconv"
	"time"
)

// Employee struct with methods to implement the ValidatableEntity interface
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

func (employee *Employee) ValidateEmployeeFields() (error, *Employee) {
	createdAt, updatedAt, err := utility.ParseTimestamps(employee.CreatedAt.String(), employee.UpdatedAt.String())
	if err != nil {
		return err, nil
	}

	employee.Epoch_CreatedAt = createdAt.Unix()
	employee.Epoch_UpdatedAt = updatedAt.Unix()

	if err := utility.ValidateStruct(employee); err != nil {
		fmt.Println("HEmployeeR validation failed")
		return err, nil
	} else {
		fmt.Println("Employee Validation successful!")
	}

	employee.Epoch_CreatedAt = createdAt.Unix()
	employee.Epoch_UpdatedAt = updatedAt.Unix()
	fmt.Println("Data validated for EmployeedID", strconv.Itoa(employee.EmployeeID))

	return nil, employee
}
