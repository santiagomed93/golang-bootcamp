package services

import (
	"github.com/santiagomed93/golangbootcamp/models"
)

type CartService struct {
	cartRepository CartRepository
}

type CartRepository interface {
	GetAllCarts() ([]models.Cart, error)
	GetCartByID(int) (models.Cart, error)
	CreateCart(models.Cart) (int, error)
	UpdateCartByID(int, models.Cart) error
	DeleteCartByID(int) error
}

func (cs *CartService) GetAllCarts() ([]models.Cart, error) {
	carts, err := cs.cartRepository.GetAllCarts()
	if err != nil {
		return []models.Cart{}, err
	}
	return carts, nil
}

func (cs *CartService) GetCartByID(idCart int) (models.Cart, error) {
	cart, err := cs.cartRepository.GetCartByID(idCart)
	if err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}

func (cs *CartService) CreateCart(cart models.Cart) (int, error) {
	result, err := cs.cartRepository.CreateCart(cart)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (cs *CartService) UpdateCartByID(idCart int, cart models.Cart) error {
	err := cs.cartRepository.UpdateCartByID(idCart, cart)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CartService) DeleteCartByID(idCart int) error {
	err := cs.cartRepository.DeleteCartByID(idCart)
	if err != nil {
		return err
	}
	return nil
}

func NewCartService(cartRepository CartRepository) *CartService {
	return &CartService{
		cartRepository: cartRepository,
	}
}
