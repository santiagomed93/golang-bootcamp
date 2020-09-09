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

func (is *ItemService) GetItem(idItem int) (models.ItemAPI, error) {
	var item models.ItemAPI
	itemResponse, err := getJSON("https://challenge.getsandbox.com/articles/"+strconv.Itoa(idItem), item)
	if err != nil {
		return models.ItemAPI{}, err
	}
	return transformToItem(itemResponse), nil
}

func (is *ItemService) GetItems() ([]models.ItemAPI, error) {
	var items []models.ItemAPI
	itemsResponse, err := getJSON("https://challenge.getsandbox.com/articles", items)
	if err != nil {
		return []models.ItemAPI{}, err
	}
	return transformToItems(itemsResponse), nil
}

func transformToItem(target interface{}) models.ItemAPI {
	var item models.ItemAPI
	jsonString, err := json.Marshal(target)
	if err != nil {
		return models.ItemAPI{}
	}
	json.Unmarshal(jsonString, &item)
	return item
}

func transformToItems(target interface{}) []models.ItemAPI {
	var items []models.ItemAPI
	jsonString, err := json.Marshal(target)
	if err != nil {
		return []models.ItemAPI{}
	}
	json.Unmarshal(jsonString, &items)
	return items
}
