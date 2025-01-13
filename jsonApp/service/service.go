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
	var company_datail *model.Company
	var admin_detail *model.Admin
	var hr_detail *model.HR
	var employees []*model.Employee

	// Loop over each company and insert into MongoDB
	for _, companyData := range compData.Companies {
		companyId := companyData.CompanyID
		company := model.Company{Name: companyData.Name, Address: companyData.Address, CompanyID: companyData.CompanyID,
			Industry: companyData.Industry, CreatedAt: companyData.CreatedAt, UpdatedAt: companyData.UpdatedAt}

		companyDetail, err := handler.ValidateCompanyData(&company)
		if err != nil {
			log.Fatal(err)
			return
		}
		company_datail = companyDetail

		for _, adminData := range companyData.Admins {
			adminData.CompanyID = companyId
			admin, err := handler.ValidateAdminFields(&adminData)
			if err != nil {
				fmt.Printf("Admin Validation fialed for Admin ID: %v Reason : %s \n", admin.AdminID, err.Error())
			}
			admin_detail = admin
		}

		for _, hrData := range companyData.HR {
			hrData.CompanyID = companyId
			hr, err := handler.ValidateHRFields(&hrData)
			if err != nil {
				fmt.Printf("HR Validation fialed for HR ID: %v Reason : %s \n", hr.HRID, err.Error())
			}
			hr_detail = hr
		}

		for _, employeeData := range companyData.Employees {
			employeeData.CompanyID = companyId
			employee, err := handler.ValidateEmployeeFields(&employeeData)
			if err != nil {
				fmt.Printf("Employee Validation fialed for Employee ID: %v Reason : %s \n", employee.EmployeeID, err.Error())
			} else {
				employees = append(employees, employee)
			}
		}
	}

	SaveCompanyDataToDB(company_datail)
	SaveAdminDataToDB(admin_detail)
	SaveHRDataToDB(hr_detail)
	SaveEmployeesDataToDB(employees)

	fmt.Println("All Operatio completed")

}
func SaveCompanyDataToDB(company *model.Company) {
	err, _ := dao.SaveCompanyData(company)
	if err != nil {
		fmt.Printf("Company data not stored due to Error %s", err.Error())
		return
	} else {
		fmt.Println("Company data stored successfully")
	}
}

func SaveEmployeesDataToDB(employees []*model.Employee) {
	err, _ := dao.SaveEmployeesData(employees)
	if err != nil {
		fmt.Printf("Employee data not stored due to Error %s", err.Error())
		return
	} else {
		fmt.Println("Employee data stored successfully")
	}
}

func SaveAdminDataToDB(admin *model.Admin) {
	err, _ := dao.SaveAdminData(admin)
	if err != nil {
		fmt.Printf("Admin data not stored due to Error %s", err.Error())
		return
	} else {
		fmt.Println("Admin data stored successfully")
	}
}
func SaveHRDataToDB(hr *model.HR) {
	err, _ := dao.SaveHRData(hr)
	if err != nil {
		fmt.Printf("HR data not stored due to Error %s", err.Error())
		return
	} else {
		fmt.Println("HR data stored successfully")
	}
}
