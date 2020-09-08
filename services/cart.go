package services

import (
	"github.com/santiagomed93/golangbootcamp/models"
)

type CartService struct {
	cartRepository CartRepository
}

type CartRepository interface {
	GetAll() ([]models.Cart, error)
}

func (cs *CartService) GetAllCarts() ([]models.Cart, error) {
	carts, err := cs.cartRepository.GetAll()
	if err != nil {
		return []models.Cart{}, err
	}
	return carts, nil
}

func NewCartService(cartRepository CartRepository) *CartService {
	return &CartService{
		cartRepository: cartRepository,
	}
}
