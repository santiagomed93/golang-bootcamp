package main

import (
	"net/http"

	"github.com/santiagomed93/golangbootcamp/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":8080", nil)
}
