package models

type ItemAPI struct {
	ID    int     `json:"id,string"`
	Title string  `json:"title"`
	Price float32 `json:"price,string"`
}
