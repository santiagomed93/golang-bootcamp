package repositories

import (
	"github.com/santiagomed93/golangbootcamp/database"
	"github.com/santiagomed93/golangbootcamp/models"
)

type CartRepository struct {
	db *database.Database
}

func (cr *CartRepository) GetAllCarts() ([]models.Cart, error) {
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

func (cr *CartRepository) GetCartByID(idCart int) (models.Cart, error) {
	cart := models.Cart{}
	err := cr.db.Conn.QueryRow(`SELECT * FROM shopping_cart
	WHERE cart_id=?`, idCart).Scan(&cart.ID, &cart.Owner)
	if err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}

func (cr *CartRepository) CreateCart(cart models.Cart) (int, error) {
	result, err := cr.db.Conn.Exec(`INSERT INTO shopping_cart
	(cart_owner) VALUES(?)`, cart.Owner)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}

func (cr *CartRepository) UpdateCartByID(idCart int, cart models.Cart) error {
	_, err := cr.db.Conn.Exec(`UPDATE shopping_cart 
	SET cart_owner = ? WHERE cart_id=?`, cart.Owner, idCart)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CartRepository) DeleteCartByID(idCart int) error {
	_, err := cr.db.Conn.Exec(`DELETE FROM shopping_cart
	WHERE cart_id=?`, idCart)
	if err != nil {
		return err
	}
	return nil
}

func NewCartRepository(db *database.Database) *CartRepository {
	return &CartRepository{
		db: db,
	}
}
