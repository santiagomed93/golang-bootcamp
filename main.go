package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/santiagomed93/golangbootcamp/controllers"
	"github.com/santiagomed93/golangbootcamp/database"
)

func main() {
	dbConn := database.SetupDatabase()
	defer dbConn.Close()
	controllers.RegisterControllers(dbConn)
	http.ListenAndServe(":8080", nil)
}
