package main

import (
	db "jsonApp/db"
	service "jsonApp/service"
)

func main() {
	// Connect to the database
	db.InitMongo()

	//for getting jsonDAta and storing in Mongodb
	service.GetDataFromJson()

	//close DB connection
	defer db.CloseDB()
}
