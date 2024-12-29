package dao

import (
	"context"
	"fmt"
	db "jsonApp/Db"
	"jsonApp/config"
	"jsonApp/model"
)

// Get the "jsonApp" database

func SaveCompanyData(company *model.Company) (*model.MyError, bool) {
	var database = db.GetDB(config.Db_Name)

	// MongoDB collections
	var companyCollection = database.Collection(config.Company_Collection)
	// Insert company data into the "companies" collection
	_, err := companyCollection.InsertOne(context.Background(), company)
	if err != nil {

		return &model.MyError{Message: err.Error()}, false
	}
	// Retrieve company ID from the insert result
	companyID1 := company.CompanyID
	fmt.Printf("Inserted company with companyid: %v\n", companyID1)
	return nil, true
}

func SaveAdminData(admin *model.Admin) (*model.MyError, bool) {
	var database = db.GetDB(config.Db_Name)
	var adminCollection = database.Collection(config.Admin_Collection)

	// Insert company data into the "companies" collection
	_, err := adminCollection.InsertOne(context.Background(), admin)
	if err != nil {

		return &model.MyError{Message: err.Error()}, false
	}
	// Retrieve company ID from the insert result
	adminid1 := admin.AdminID
	fmt.Printf("Inserted Admin with Adminid: %v\n", adminid1)
	return nil, true
}

// func SaveAdminsData(admins []model.Admin) (*model.MyError, bool) {

// }

func SaveHRData(hr *model.HR) (*model.MyError, bool) {
	var database = db.GetDB(config.Db_Name)

	var hrCollection = database.Collection(config.HR_Collection)
	// Insert company data into the "companies" collection
	_, err := hrCollection.InsertOne(context.Background(), hr)
	if err != nil {

		return &model.MyError{Message: err.Error()}, false
	}
	// Retrieve company ID from the insert result
	hrid := hr.HRID
	fmt.Printf("Inserted hr with HRID: %v\n", hrid)
	return nil, true
}

// func SaveHRsData(hrs []model.HR) (*model.MyError, bool) {

// }

func SaveEmployeeData(employee *model.Employee) (*model.MyError, bool) {
	var database = db.GetDB(config.Db_Name)

	var employeeCollection = database.Collection(config.Employee_Collection)

	// Insert company data into the "companies" collection
	_, err := employeeCollection.InsertOne(context.Background(), employee)
	if err != nil {

		return &model.MyError{Message: err.Error()}, false
	}
	// Retrieve company ID from the insert result
	empid := employee.EmployeeID
	fmt.Printf("Inserted Employee with empid: %v\n", empid)
	return nil, true
}

func SaveEmployeesData(employees []model.Employee) (*model.MyError, bool) {
	var database = db.GetDB(config.Db_Name)

	var employeeCollection = database.Collection(config.Employee_Collection)

	var employeesInterface []interface{}
	for _, emp := range employees {
		employeesInterface = append(employeesInterface, emp)
	}

	// Insert multiple employees into the collection
	result, err := employeeCollection.InsertMany(context.Background(), employeesInterface)
	if err != nil {
		return &model.MyError{Message: err.Error()}, false
	}

	// Print the inserted employee IDs
	for _, id := range result.InsertedIDs {
		fmt.Printf("Inserted employee with ID: %v\n", id)
	}

	return nil, true
}