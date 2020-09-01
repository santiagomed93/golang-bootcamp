package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/santiagomed93/golangbootcamp/models"
)

type itemController struct {
	itemIDPattern *regexp.Regexp
}

func (it itemController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/items" {
		switch request.Method {
		case http.MethodGet:
			it.getAll(response, request)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := it.itemIDPattern.FindStringSubmatch(request.URL.Path)
		if len(matches) == 0 {
			response.WriteHeader(http.StatusNotFound)
		}
		idItem, err := strconv.Atoi(matches[1])
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
		}
		switch request.Method {
		case http.MethodGet:
			it.get(idItem, response)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (it *itemController) getAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	encodeResponseAsJSON(models.GetItems(), response)
}

func (it *itemController) get(idItem int, response http.ResponseWriter) {
	item, err := models.GetItemById(idItem)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	encodeResponseAsJSON(item, response)
}

func newItemController() *itemController {
	return &itemController{
		itemIDPattern: regexp.MustCompile(`^/items/(\d+)/?`),
	}
}
