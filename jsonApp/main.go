package main

import (
	db "jsonApp/db"
	"jsonApp/routes"
	service "jsonApp/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db.InitMongo()

	//for getting jsonDAta and storing in Mongodb
	service.GetDataFromJson()

	r := gin.Default()

	// Register all routes
	routes.RegisterRoutes(r)

	// Start the server
	log.Println("Server running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	//close DB connection
	defer db.CloseDB()
}
