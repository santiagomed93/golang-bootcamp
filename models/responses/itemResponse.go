package responses

type ItemResponse struct {
	CartID   int `json:"cartID"`
	ItemID   int `json:"itemID"`
	Quantity int `json:"quantity"`
}
