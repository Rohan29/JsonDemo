package main

import (
	"context"
	"encoding/json"
	"fmt"
	db "jsonApp/Db"
	utility "jsonApp/Utility"
	"jsonApp/model"
	"log"
	"sync"
	"time"
)

func main() {
	// MongoDB URI (replace with your actual URI)
	mongoURI := "mongodb://localhost:27017"

	// Connect to the database
	_, err := db.ConnectDB(mongoURI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer db.CloseDB() // Ensure the connection is closed when the main function exits

	// Get the "jsonApp" database
	database := db.GetDB("jsonApp")

	// MongoDB collections
	companyCollection := database.Collection("company")
	adminCollection := database.Collection("admin")
	hrCollection := database.Collection("hr")
	employeeCollection := database.Collection("employee")

	// Unmarshal JSON data into Go structs
	filePath := "./data.json"

	// Call the method to read the file
	fileContents, err := utility.ReadCompaniesFromFile(filePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var companiesData struct {
		Companies []struct {
			CompanyID int              `json:"companyId"`
			Name      string           `json:"name"`
			Address   string           `json:"address"`
			Industry  string           `json:"industry"`
			CreatedAt time.Time        `json:"createdAt"`
			UpdatedAt time.Time        `json:"updatedAt"`
			Admins    []model.Admin    `json:"admins"`
			HR        []model.HR       `json:"hr"`
			Employees []model.Employee `json:"employees"`
		} `json:"companies"`
	}

	err = json.Unmarshal(fileContents, &companiesData)
	if err != nil {
		log.Fatal(err)
	}
	companyId := 0
	// Loop over each company and insert into MongoDB
	for _, companyData := range companiesData.Companies {
		// Ensure `createdAt` and `updatedAt` are parsed into `time.Time`
		companyId = companyData.CompanyID
		company := new(model.Company)
		company.Name = companyData.Name
		company.Address = companyData.Address
		company.CompanyID = companyData.CompanyID
		company.Industry = companyData.Industry
		company.CreatedAt, err = utility.ParseTime(companyData.CreatedAt.String())
		if err != nil {
			log.Fatal("Error parsing CreatedAt for company:", err)

		}
		company.UpdatedAt, err = utility.ParseTime(companyData.UpdatedAt.String())
		if err != nil {
			log.Fatal("Error parsing UpdatedAt for company:", err)
		}
		var wg sync.WaitGroup
		errChan := make(chan error, 10)
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Insert company data into the "companies" collection
			companyResult, err := companyCollection.InsertOne(context.Background(), company)
			if err != nil {
				errChan <- err
				//log.Fatal("Failed to insert company:", err)
			}

			// Retrieve company ID from the insert result
			companyID := companyResult.InsertedID
			companyID1 := company.CompanyID

			fmt.Printf("Inserted company with ID: %v and companyid: %v\n", companyID, companyID1)

		}()

		// Insert admins data into the "admins" collection
		for _, admin := range companyData.Admins {
			admin.CompanyID = companyId
			wg.Add(1)
			go func() {

				defer wg.Done()
				_, err := adminCollection.InsertOne(context.Background(), admin)
				if err != nil {
					//log.Fatal(err)
					errChan <- err
				}
				fmt.Printf("Inserted admin: %v\n", admin.AdminID)
			}()
		}

		// Insert hr data into the "hrs" collection
		for _, hr := range companyData.HR {
			hr.CompanyID = companyId
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := hrCollection.InsertOne(context.Background(), hr)
				if err != nil {
					//log.Fatal(err)
					errChan <- err
				}
				fmt.Printf("Inserted HR: %v\n", hr.HRID)
			}()

		}

		// Insert employees data into the "employees" collection
		for _, employee := range companyData.Employees {
			employee.CompanyID = companyId
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := employeeCollection.InsertOne(context.Background(), employee)
				if err != nil {
					//log.Fatal(err)
					errChan <- err
				}
				fmt.Printf("Inserted employee: %v\n", employee.EmployeeID)
			}()
		}

		wg.Wait()
		// Check for any errors
		close(errChan)
		for err := range errChan {
			log.Println("Error:", err)
		}

		fmt.Println("All data inserted successfully.")
	}
}
