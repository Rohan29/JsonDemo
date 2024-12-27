package model

import "time"

type Admin struct {
	AdminID   int       `json:"adminId" bson:"adminId"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Phone     string    `json:"phone" bson:"phone"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	CompanyID int       `json:"companyId" bson:"companyId"`
}
type HR struct {
	HRID       int       `json:"hrId" bson:"hrId"`
	Name       string    `json:"name" bson:"name"`
	Email      string    `json:"email" bson:"email"`
	Phone      string    `json:"phone" bson:"phone"`
	Department string    `json:"department" bson:"department"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	CompanyID  int       `json:"companyId" bson:"companyId"`
}

type Employee struct {
	EmployeeID int       `json:"employeeId" bson:"employeeId"`
	Name       string    `json:"name" bson:"name"`
	Email      string    `json:"email" bson:"email"`
	Phone      string    `json:"phone" bson:"phone"`
	Position   string    `json:"position" bson:"position"`
	Salary     float64   `json:"salary" bson:"salary"`
	HireDate   string    `json:"hireDate" bson:"hireDate"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	CompanyID  int       `json:"companyId" bson:"companyId"`
}

type Company struct {
	CompanyID int       `json:"companyId" bson:"companyId"`
	Name      string    `json:"name" bson:"name"`
	Address   string    `json:"address" bson:"address"`
	Industry  string    `json:"industry" bson:"industry"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
