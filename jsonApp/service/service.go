package service

import (
	"encoding/json"
	"fmt"
	"jsonApp/config"
	"jsonApp/dao"
	handler "jsonApp/handler"
	model "jsonApp/model"
	"jsonApp/utility"
	"log"
)

func GetDataFromJson() {
	fileContents, filerr := utility.ReadCompaniesFromFile(config.FilePath)
	if filerr != nil {
		log.Fatalf("Error: %v", filerr)
		return
	}

	var compData model.JsonCompaniesData

	err := json.Unmarshal(fileContents, &compData)
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

		for _, admin := range companyData.Admins {
			admin.CompanyID = companyId
			err, newAdmin := handler.ValidateAdminFields(admin)
			if err != nil {
				fmt.Printf("Admin Validation fialed for Admin ID: %v Reason : %s \n", admin.AdminID, err.Message)
			} else {
				myAdmins = newAdmin
			}
			fmt.Println("Admin data", admin)
		}

		for _, hr := range companyData.HR {
			hr.CompanyID = companyId
			err, newHR := handler.ValidateHRFields(hr)
			if err != nil {
				fmt.Printf("HR Validation fialed for HR ID: %v Reason : %s \n", hr.HRID, err.Message)
			} else {
				myHr = newHR
			}
		}

		for _, employee := range companyData.Employees {
			employee.CompanyID = companyId
			err, newEmployee := handler.ValidateEmployeeFields(employee)
			if err != nil {
				fmt.Printf("Employee Validation fialed for Employee ID: %v Reason : %s \n", employee.EmployeeID, err.Message)
			} else {
				myEmployee = append(myEmployee, *newEmployee)
			}
		}

	}
	dao.SaveCompanyData(mycompanyData)
	dao.SaveAdminData(myAdmins)
	dao.SaveHRData(myHr)
	dao.SaveEmployeesData(myEmployee)

	fmt.Println("All data inserted successfully.")

}
