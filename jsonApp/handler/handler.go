package handler

import (
	"jsonApp/model"
)

func ValidateCompanyData(company *model.Company) (*model.Company, error) {
	com, err := company.ValidateCompanyData()
	return com, err
}

func ValidateAdminFields(admin *model.Admin) (*model.Admin, error) {
	admin, err := admin.ValidateAdminFields()
	return admin, err
}

// validateHRFields validates HR's fields
func ValidateHRFields(hr *model.HR) (*model.HR, error) {
	hr, err := hr.ValidateHRFields()
	return hr, err
}

// validateEmployeeFields validates an employee's fields
func ValidateEmployeeFields(employee *model.Employee) (*model.Employee, error) {
	err, employee := employee.ValidateEmployeeFields()
	return employee, err
}
