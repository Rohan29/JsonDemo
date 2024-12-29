package handler

import (
	"fmt"
	model "jsonApp/model"
	utility "jsonApp/utility"
	"log"
	"regexp"
	"strconv"
	"time"
)

func ValidateCompanyData(name string, address string, industry string, companyID int, createdAt time.Time, updatedAt time.Time) (*model.Company, *model.MyError) {
	CreatedAt, UpdatedAt, err := utility.ParseTimestamps(createdAt.String(), updatedAt.String())
	if err != nil {
		log.Fatalf("Error parsing timestamps: %v", err)
		return nil, err
	}
	company := model.Company{Name: name, Address: address, CompanyID: companyID,
		Industry: industry, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt}

	if err := ValidateCompanyFields(&company); err != nil {
		fmt.Println("Validation failed:", err)
		return nil, err
	} else {
		fmt.Println("Company Validation successful!")
	}
	fmt.Printf("Company %v \n", company)
	return &company, nil
}

// ValidateCompanyFields validates a company's fields
func ValidateCompanyFields(company *model.Company) *model.MyError {
	// Define field validations for Company
	fields := []model.Field{
		{Name: "Company ID", Value: strconv.Itoa(company.CompanyID), Regex: nil},
		{Name: "Company Name", Value: company.Name, Regex: regexp.MustCompile(`^[A-Za-z0-9\s\.,&\-]{1,100}$`)},
		{Name: "Address", Value: company.Address, Regex: regexp.MustCompile(`^[A-Za-z0-9\s,.\-]{1,200}$`)},
		{Name: "Industry", Value: company.Industry, Regex: regexp.MustCompile(`^[A-Za-z\s]{1,50}$`)},
	}
	// Validate all fields
	if err := utility.ValidateEntity(fields); err != nil {
		return err
	}
	return nil
}

// validateAdminFields validates an admin's fields
func ValidateAdminFields(admin model.Admin) (*model.MyError, *model.Admin) {
	createdAt, updatedAt, err := utility.ParseTimestamps(admin.CreatedAt.String(), admin.UpdatedAt.String())
	if err != nil {
		log.Fatalf("Error parsing timestamps: %v", err)
		return err, nil
	}
	admin.CreatedAt = createdAt
	admin.UpdatedAt = updatedAt

	fields := []model.Field{
		{Name: "Admin ID", Value: strconv.Itoa(admin.AdminID), Regex: nil},
		{Name: "Name", Value: admin.Name, Regex: regexp.MustCompile(`^[A-Za-z\s]{1,100}$`)},
		{Name: "Email", Value: admin.Email, Regex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)},
		{Name: "Phone", Value: admin.Phone, Regex: regexp.MustCompile(`^\+?\d{1,3}-\d{3}-\d{3}-\d{4}$`)},
	}
	// Validate all fields
	if err := utility.ValidateEntity(fields); err != nil {
		return err, nil

	}

	fmt.Println("Data validated for Adminid", strconv.Itoa(admin.AdminID))

	return nil, &admin
}

// validateHRFields validates HR's fields
func ValidateHRFields(hr model.HR) (*model.MyError, *model.HR) {
	createdAt, updatedAt, err := utility.ParseTimestamps(hr.CreatedAt.String(), hr.UpdatedAt.String())
	if err != nil {
		log.Fatalf("Error parsing timestamps: %v", err)
		return err, nil
	}
	hr.CreatedAt = createdAt
	hr.UpdatedAt = updatedAt

	fields := []model.Field{
		{Name: "HR ID", Value: strconv.Itoa(hr.HRID), Regex: nil},
		{Name: "Name", Value: hr.Name, Regex: regexp.MustCompile(`^[A-Za-z\s]{1,100}$`)},
		{Name: "Email", Value: hr.Email, Regex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)},
		{Name: "Phone", Value: hr.Phone, Regex: regexp.MustCompile(`^\+?\d{1,3}-\d{3}-\d{3}-\d{4}$`)},
		{Name: "Department", Value: hr.Department, Regex: regexp.MustCompile(`^[A-Za-z\s]{1,100}$`)},
	}
	// Validate all fields
	if err := utility.ValidateEntity(fields); err != nil {
		return err, nil
	}
	fmt.Println("Data validated for HRID", strconv.Itoa(hr.HRID))
	return nil, &hr
}

// validateEmployeeFields validates an employee's fields
func ValidateEmployeeFields(employee model.Employee) (*model.MyError, *model.Employee) {
	createdAt, updatedAt, err := utility.ParseTimestamps(employee.CreatedAt.String(), employee.UpdatedAt.String())
	if err != nil {
		log.Fatalf("Error parsing timestamps: %v", err)
		return err, nil
	}
	employee.CreatedAt = createdAt
	employee.UpdatedAt = updatedAt

	fields := []model.Field{
		{Name: "Employee ID", Value: strconv.Itoa(employee.EmployeeID), Regex: nil},
		{Name: "Name", Value: employee.Name, Regex: regexp.MustCompile(`^[A-Za-z\s]{1,100}$`)},
		{Name: "Email", Value: employee.Email, Regex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)},
		{Name: "Phone", Value: employee.Phone, Regex: regexp.MustCompile(`^\+?\d{1,3}-\d{3}-\d{3}-\d{4}$`)},
		{Name: "Position", Value: employee.Position, Regex: regexp.MustCompile(`^[A-Za-z\s]{1,100}$`)},
		{Name: "Salary", Value: strconv.Itoa(int(employee.Salary)), Regex: nil},
		{Name: "HireDate", Value: employee.HireDate, Regex: nil}, // Date validation will be done separately
	}

	// Validate all fields
	if err := utility.ValidateEntity(fields); err != nil {
		return err, nil
	}
	fmt.Println("Data validated for EmployeedID", strconv.Itoa(employee.EmployeeID))

	return nil, &employee
}
