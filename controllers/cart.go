package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/santiagomed93/golangbootcamp/models"
)

type cartController struct {
	cartIDPath    *regexp.Regexp
	cartItemsPath *regexp.Regexp
	cartItemID    *regexp.Regexp
	cartService   CartService
}

type CartService interface {
	GetAllCarts() ([]models.Cart, error)
	GetCartByID(int) (models.Cart, error)
	CreateCart(models.Cart) (int, error)
	UpdateCartByID(int, models.Cart) error
	DeleteCartByID(int) error
}

func (ca cartController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if request.URL.Path == "/carts" {
		switch request.Method {
		case http.MethodGet:
			ca.getAllCarts(response, request)
		case http.MethodPost:
			ca.createCart(response, request)
		default:
			response.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		match := false
		cartIDPath := ca.cartIDPath.FindStringSubmatch(request.URL.Path)
		if len(cartIDPath) != 0 {
			match = true
			cartID, err := strconv.Atoi(cartIDPath[1])
			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("Could not get cart id"))
				return
			}
			switch request.Method {
			case http.MethodGet:
				ca.getCartByID(cartID, response)
			case http.MethodPut:
				ca.modifyCartByID(cartID, response, request)
			case http.MethodDelete:
				ca.deleteCartByID(cartID, response)
			default:
				response.WriteHeader(http.StatusNotImplemented)
			}
		}

		cartItemsPath := ca.cartItemsPath.FindStringSubmatch(request.URL.Path)
		if len(cartItemsPath) != 0 {
			match = true
			cartID, err := strconv.Atoi(cartItemsPath[1])
			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("Could not get cart id"))
				return
			}
			switch request.Method {
			case http.MethodPost:
				ca.createCartItem(cartID, response, request)
			default:
				response.WriteHeader(http.StatusNotImplemented)
			}
		}

		cartItemID := ca.cartItemID.FindStringSubmatch(request.URL.Path)
		if len(cartItemID) != 0 {
			match = true
			cartID, err := strconv.Atoi(cartItemID[1])
			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("Could not get cart id"))
				return
			}
			itemID, err := strconv.Atoi(cartItemID[2])
			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("Could not get item id"))
				return
			}
			switch request.Method {
			case http.MethodPut:
				ca.modifyItemCartQuantity(cartID, itemID, response, request)
			case http.MethodDelete:
				ca.deleteItemCart(cartID, itemID, response)
			default:
				response.WriteHeader(http.StatusNotImplemented)
			}
		}
		if !match {
			response.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (ca *cartController) getAllCarts(response http.ResponseWriter, request *http.Request) {
	carts, err := ca.cartService.GetAllCarts()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not get carts"))
		return
	}
	encodeResponseAsJSON(carts, response)
}

func (ca *cartController) getCartByID(id int, response http.ResponseWriter) {
	cart, err := ca.cartService.GetCartByID(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not get cart by id"))
		return
	}
	encodeResponseAsJSON(cart, response)
}

func (ca *cartController) createCart(response http.ResponseWriter, request *http.Request) {
	cart, err := ca.parseRequest(request)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not parse cart object"))
		return
	}
	if cart.ID != 0 {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("It's not necessary id field"))
		return
	}
	createdCart, err := ca.cartService.CreateCart(cart)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not create cart"))
		return
	}
	encodeResponseAsJSON(createdCart, response)

}

func (ca *cartController) modifyCartByID(id int, response http.ResponseWriter, request *http.Request) {
	cart, err := ca.parseRequest(request)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not parse cart object"))
		return
	}
	err = ca.cartService.UpdateCartByID(id, cart)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not update cart"))
		return
	}
	response.WriteHeader(http.StatusOK)
}

func (ca *cartController) deleteCartByID(id int, response http.ResponseWriter) {
	err := ca.cartService.DeleteCartByID(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Could not delete cart"))
		return
	}
	response.WriteHeader(http.StatusOK)
}

func (ca *cartController) createCartItem(id int, response http.ResponseWriter, request *http.Request) {
	//TODO
}

func (ca *cartController) modifyItemCartQuantity(cartID int, itemID int, response http.ResponseWriter, request *http.Request) {
	//TODO
}

func (ca *cartController) deleteItemCart(cartID int, itemID int, response http.ResponseWriter) {
	//TODO
}

func (ca *cartController) parseRequest(r *http.Request) (models.Cart, error) {
	var c models.Cart
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return models.Cart{}, err
	}
	return c, nil
}

func NewCartController(cartService CartService) *cartController {
	return &cartController{
		cartIDPath:    regexp.MustCompile(`/carts/(\d+)$`),
		cartItemsPath: regexp.MustCompile(`/carts/(\d+)/items$`),
		cartItemID:    regexp.MustCompile(`/carts/(\d+)/items/(\d+)$`),
		cartService:   cartService,
	}
}
