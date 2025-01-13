package model

import (
	"time"
)

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

type CustomError struct {
	Message string
}

// Error implements error.
func (c *CustomError) Error() string {
	panic("unimplemented")
}
