// package main

// import (
// 	"BankingApp/models"
// 	"BankingApp/routes"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {

// }

package main

import (
	"BankingApp/models"
	"BankingApp/routes"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/bankingapp?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	models.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	models.DB.AutoMigrate(&models.User{})

	// Initialize routes
	router := routes.InitRoutes()

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
