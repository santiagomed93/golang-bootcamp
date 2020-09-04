package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/santiagomed93/golangbootcamp/database"

	"github.com/santiagomed93/golangbootcamp/repositories"
	"github.com/santiagomed93/golangbootcamp/services"
)

func RegisterControllers(db *database.Database) {
	RegisterItemControllers()
	RegisterCartController(db)
}

func RegisterItemControllers() {
	itemService := services.NewItemService()
	itemController := NewItemController(itemService)
	http.Handle("/items", *itemController)
	http.Handle("/items/", *itemController)
}

func RegisterCartController(db *database.Database) {
	cartRepository := repositories.NewCartRepository(db)
	cartService := services.NewCartService(cartRepository)
	cartController := NewCartController(cartService)
	http.Handle("/carts", *cartController)
	http.Handle("/carts/", *cartController)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
