package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/santiagomed93/golangbootcamp/models"
)

type cartController struct {
	cartIDPattern *regexp.Regexp
	cartService   CartService
}

type CartService interface {
	GetAllCarts() ([]models.Cart, error)
}

func (ca cartController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if request.URL.Path == "/carts" {
		switch request.Method {
		case http.MethodGet:
			ca.getAll(response, request)
		case http.MethodPost:
			ca.post(response, request)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (ca *cartController) getAll(response http.ResponseWriter, request *http.Request) {
	cars, err := ca.cartService.GetAllCarts()
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not get the carts"))
		return
	}
	encodeResponseAsJSON(cars, response)
}

func (ca *cartController) post(w http.ResponseWriter, r *http.Request) {
	cart, err := ca.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Cart object"))
		return
	}
	fmt.Println(cart)

}

func (ca *cartController) parseRequest(r *http.Request) (models.Cart, error) {
	dec := json.NewDecoder(r.Body)
	var c models.Cart
	err := dec.Decode(&c)
	if err != nil {
		return models.Cart{}, err
	}
	return c, nil
}

func NewCartController(cartService CartService) *cartController {
	return &cartController{
		cartIDPattern: regexp.MustCompile(`^/carts/(\d+)/?`),
		cartService:   cartService,
	}
}
