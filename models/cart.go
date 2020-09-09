package models

type Cart struct {
	ID    int      `json:"id,string"`
	Owner string   `json:"owner"`
	Items []ItemDB `json:"items"`
}
