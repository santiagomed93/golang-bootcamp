package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/santiagomed93/golangbootcamp/services"
)

func RegisterControllers(db *sql.DB) {
	RegisterItemControllers(db)
}

func RegisterItemControllers(db *sql.DB) {
	itemService := services.NewItemService()
	itemController := newItemController(itemService)
	http.Handle("/items", *itemController)
	http.Handle("/items/", *itemController)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
