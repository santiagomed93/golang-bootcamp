package services

import (
	"encoding/json"
	"strconv"

	"github.com/santiagomed93/golangbootcamp/models"
)

type ItemService struct {
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (is *ItemService) GetItem(idItem int) (models.Item, error) {
	var item models.Item
	itemResponse, err := getJSON("https://challenge.getsandbox.com/articles/"+strconv.Itoa(idItem), item)
	if err != nil {
		return models.Item{}, err
	}
	return transformToItem(itemResponse), nil
}

func (is *ItemService) GetItems() ([]models.Item, error) {
	var items []models.Item
	itemsResponse, err := getJSON("https://challenge.getsandbox.com/articles", items)
	if err != nil {
		return []models.Item{}, err
	}
	return transformToItems(itemsResponse), nil
}

func transformToItem(target interface{}) models.Item {
	var item models.Item
	jsonString, err := json.Marshal(target)
	if err != nil {
		return models.Item{}
	}
	json.Unmarshal(jsonString, &item)
	return item
}

func transformToItems(target interface{}) []models.Item {
	var items []models.Item
	jsonString, err := json.Marshal(target)
	if err != nil {
		return []models.Item{}
	}
	json.Unmarshal(jsonString, &items)
	return items
}
