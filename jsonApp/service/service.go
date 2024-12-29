package service

import (
	"encoding/json"
	"fmt"
	"jsonApp/dao"
	handler "jsonApp/handler"
	model "jsonApp/model"
	"log"
)

func GetDataFromJson(data []byte) {

	var compData model.JsonCompaniesData

	err := json.Unmarshal(data, &compData)
	if err != nil {
		log.Fatal(err)
	}

	var myAdmins *model.Admin
	var myHr *model.HR
	var myEmployee []model.Employee
	var mycompanyData *model.Company
	// Loop over each company and insert into MongoDB
	for _, companyData := range compData.Companies {
		companyId := companyData.CompanyID
		companyData1, err := handler.ValidateCompanyData(companyData.Name, companyData.Address, companyData.Industry, companyData.CompanyID, companyData.CreatedAt, companyData.UpdatedAt)
		if err != nil {
			log.Fatal(err)
			return
		} else {
			mycompanyData = companyData1
		}

		// var wg sync.WaitGroup
		// errChan := make(chan error, 10)
		// fmt.Println("Company Data", company)
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	// Insert company data into the "companies" collection
		// 	companyResult, err := companyCollection.InsertOne(context.Background(), company)
		// 	if err != nil {
		// 		errChan <- err
		// 		//log.Fatal("Failed to insert company:", err)
		// 	}
		// 	// Retrieve company ID from the insert result
		// 	companyID := companyResult.InsertedID
		// 	companyID1 := company.CompanyID
		// 	fmt.Printf("Inserted company with ID: %v and companyid: %v\n", companyID, companyID1)
		// }()

		// // Insert admins data into the "admins" collection
		for _, admin := range companyData.Admins {
			admin.CompanyID = companyId
			err, newAdmin := handler.ValidateAdminFields(admin)
			if err != nil {
				fmt.Printf("Admin Validation fialed for Admin ID: %v Reason : %s \n", admin.AdminID, err.Message)
			} else {
				myAdmins = newAdmin
			}
			fmt.Println("Admin data", admin)
			// _, err := adminCollection.InsertOne(context.Background(), admin)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// fmt.Printf("Inserted admin: %v\n", admin.AdminID)
		}

		// // // Insert hr data into the "hrs" collection
		for _, hr := range companyData.HR {
			hr.CompanyID = companyId
			err, newHR := handler.ValidateHRFields(hr)
			if err != nil {
				fmt.Printf("HR Validation fialed for HR ID: %v Reason : %s \n", hr.HRID, err.Message)
			} else {
				myHr = newHR
			}
			// _, err := hrCollection.InsertOne(context.Background(), hr)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// fmt.Printf("Inserted HR: %v\n", hr.HRID)
		}

		// // Insert employees data into the "employees" collection
		for _, employee := range companyData.Employees {
			employee.CompanyID = companyId
			err, newEmployee := handler.ValidateEmployeeFields(employee)
			if err != nil {
				fmt.Printf("Employee Validation fialed for Employee ID: %v Reason : %s \n", employee.EmployeeID, err.Message)
			} else {
				myEmployee = append(myEmployee, *newEmployee)
			}
			// _, err := employeeCollection.InsertOne(context.Background(), employee)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// fmt.Printf("Inserted employee: %v\n", employee.EmployeeID)
		}

		fmt.Println("All data inserted successfully.")
	}
	dao.SaveCompanyData(mycompanyData)
	dao.SaveAdminData(myAdmins)
	dao.SaveHRData(myHr)
	dao.SaveEmployeesData(myEmployee)
}
