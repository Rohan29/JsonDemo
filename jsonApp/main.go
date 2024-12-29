package main

import (
	db "jsonApp/Db"
	config "jsonApp/config"
	service "jsonApp/service"
	utility "jsonApp/utility"
	"log"
)

func main() {
	// Connect to the database
	_, err := db.ConnectDB(config.Mongo_URI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	// Ensure the connection is closed when the main function exits
	// Unmarshal JSON data into Go structs

	// Call the method to read the file
	fileContents, err := utility.ReadCompaniesFromFile(config.FilePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	service.GetDataFromJson(fileContents)

	defer db.CloseDB()
}
