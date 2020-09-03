package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	itemController := newItemController()
	http.Handle("/items", *itemController)
	http.Handle("/items/", *itemController)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
