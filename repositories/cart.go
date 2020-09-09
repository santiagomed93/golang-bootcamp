package repositories

import (
	"database/sql"

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
	findCartByID := `SELECT * FROM shopping_cart WHERE cart_id=?`
	err := cr.db.Conn.QueryRow(findCartByID, idCart).Scan(&cart.ID, &cart.Owner)
	if err != nil {
		return models.Cart{}, err
	}
	findCartChildren := `SELECT item_id,quantity FROM item_cart WHERE cart_id=?`
	results, err := cr.db.Conn.Query(findCartChildren, idCart)
	if err != nil && err != sql.ErrNoRows {
		return models.Cart{}, err
	}
	if err == sql.ErrNoRows {
		return cart, nil
	}
	items := make([]models.ItemDB, 0)
	for results.Next() {
		var item models.ItemDB
		results.Scan(&item.ID, &item.Quantity)
		items = append(items, item)
	}
	cart.Items = items
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
	findCartChildren := `SELECT item_id FROM item_cart WHERE cart_id=?`
	results, err := cr.db.Conn.Query(findCartChildren, idCart)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		return cr.deleteCart(idCart)
	}
	itemsID := make([]int, 0)
	for results.Next() {
		var itemID int
		results.Scan(&itemID)
		itemsID = append(itemsID, itemID)
	}
	deleteCartItem := `DELETE FROM item_cart WHERE item_id=?`
	for _, itemID := range itemsID {
		_, err := cr.db.Conn.Exec(deleteCartItem, itemID)
		if err != nil {
			return err
		}
	}
	return cr.deleteCart(idCart)
}

func (cr *CartRepository) deleteCart(idCart int) error {
	deleteCart := `DELETE FROM shopping_cart WHERE cart_id=?`
	_, err := cr.db.Conn.Exec(deleteCart, idCart)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CartRepository) CreateCartItem(idCart int, item models.ItemDB) error {
	var quantity int
	checkCartItem := `SELECT quantity FROM item_cart WHERE item_id=? AND cart_id=?`
	err := cr.db.Conn.QueryRow(checkCartItem, item.ID, idCart).Scan(&quantity)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		createCartItem := `INSERT INTO item_cart (item_id,quantity,cart_id) VALUES(?,?,?)`
		_, err := cr.db.Conn.Exec(createCartItem, item.ID, 1, idCart)
		if err != nil {
			return err
		}
		return nil
	}

	quantity = quantity + 1
	updateCartItem := `UPDATE item_cart SET quantity=? WHERE item_id=? AND cart_id=?`
	_, err = cr.db.Conn.Exec(updateCartItem, quantity, item.ID, idCart)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CartRepository) UpdateCartItemQuantity(cartID int, itemID int, quantity int) error {
	updateCartItemQuantity := `UPDATE item_cart SET quantity=? WHERE cart_id=? AND item_id=?`
	_, err := cr.db.Conn.Exec(updateCartItemQuantity, quantity, cartID, itemID)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CartRepository) DeleteCartItemByID(cartID int, itemID int) error {
	deleteCartItem := `DELETE FROM item_cart WHERE cart_id=? AND item_id=?`
	_, err := cr.db.Conn.Exec(deleteCartItem, cartID, itemID)
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
