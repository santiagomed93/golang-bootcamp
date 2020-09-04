package repositories

import (
	"github.com/santiagomed93/golangbootcamp/database"
	"github.com/santiagomed93/golangbootcamp/models"
)

type CartRepository struct {
	db *database.Database
}

func (cr *CartRepository) GetAll() ([]models.Cart, error) {
	results, err := cr.db.Conn.Query(`SELECT * FROM shopping_cart`)
	if err != nil {
		return nil, err
	}
	carts := make([]models.Cart, 0)
	for results.Next() {
		var cart models.Cart
		results.Scan(&cart.ID, &cart.Owner)
		carts = append(carts, cart)
	}
	return carts, nil
}

func NewCartRepository(db *database.Database) *CartRepository {
	return &CartRepository{
		db: db,
	}
}
