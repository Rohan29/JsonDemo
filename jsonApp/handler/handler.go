package handler

import (
	"fmt"
	model "jsonApp/model"
	utility "jsonApp/utility"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/go-playground/validator"
)

func regexValidator(fl validator.FieldLevel) bool {
	tag := fl.Param()
	match, _ := regexp.MatchString(tag, fl.Field().String())
	return match
}

// ValidateStruct validates any struct based on its tags
func validateStruct(s interface{}) *model.MyError {
	validate := validator.New()
	validate.RegisterValidation("regex", regexValidator)

	if err := validate.Struct(s); err != nil {
		errorMessages := ""
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages += fmt.Sprintf("Field '%s' failed validation with tag '%s'. ", err.StructField(), err.Tag())
		}
		return &model.MyError{Message: errorMessages}
	}
	return nil
}

func ValidateCompanyData(name string, address string, industry string, companyID int, createdAt time.Time, updatedAt time.Time) (*model.Company, *model.MyError) {
	CreatedAt, UpdatedAt, err := utility.ParseTimestamps(createdAt.String(), updatedAt.String())
	if err != nil {
		log.Fatalf("Error parsing timestamps: %v", err)
		return nil, err
	}

	company := model.Company{Name: name, Address: address, CompanyID: companyID,
		Industry: industry, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt}

	if err := validateStruct(company); err != nil {
		fmt.Println("Company validation failed")
		return nil, err
	} else {
		fmt.Println("Company validation successful!")
	}
	company.Epoch_CreatedAt = createdAt.Unix()
	company.Epoch_UpdatedAt = UpdatedAt.Unix()
	fmt.Printf("Company %v \n", company)
	return &company, nil
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
	if err := validateStruct(admin); err != nil {
		fmt.Println("Admin validation failed")
		return err, nil
	} else {
		fmt.Println("Admin Validation successful!")
	}
	admin.Epoch_CreatedAt = createdAt.Unix()
	admin.Epoch_UpdatedAt = updatedAt.Unix()

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

	if err := validateStruct(hr); err != nil {
		fmt.Println("HR validation failed")
		return err, nil
	} else {
		fmt.Println("HR Validation successful!")
	}
	hr.Epoch_CreatedAt = createdAt.Unix()
	hr.Epoch_UpdatedAt = updatedAt.Unix()

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

	employee.Epoch_CreatedAt = createdAt.Unix()
	employee.Epoch_UpdatedAt = updatedAt.Unix()

	if err := validateStruct(employee); err != nil {
		fmt.Println("HEmployeeR validation failed")
		return err, nil
	} else {
		fmt.Println("Employee Validation successful!")
	}

	employee.Epoch_CreatedAt = createdAt.Unix()
	employee.Epoch_UpdatedAt = updatedAt.Unix()
	fmt.Println("Data validated for EmployeedID", strconv.Itoa(employee.EmployeeID))

	return nil, &employee
}
