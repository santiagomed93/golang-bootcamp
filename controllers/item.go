package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/santiagomed93/golangbootcamp/models"
)

type itemController struct {
	itemIDPattern *regexp.Regexp
	ItemService
}

type ItemService interface {
	GetItem(idItem int) (models.Item, error)
	GetItems() ([]models.Item, error)
}

func (it itemController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
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
			return
		}
		idItem, err := strconv.Atoi(matches[1])
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
		}
		switch request.Method {
		case http.MethodGet:
			it.getByID(idItem, response)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (it *itemController) getAll(response http.ResponseWriter, request *http.Request) {
	items, err := it.GetItems()
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(items, response)
}

func (it *itemController) getByID(idItem int, response http.ResponseWriter) {
	item, err := it.GetItem(idItem)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	if item == (models.Item{}) {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	encodeResponseAsJSON(item, response)
}

func newItemController(itemService ItemService) *itemController {
	return &itemController{
		itemIDPattern: regexp.MustCompile(`^/items/(\d+)/?`),
		ItemService:   itemService,
	}
}
