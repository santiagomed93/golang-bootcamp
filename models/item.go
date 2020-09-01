package models

import "fmt"

type Item struct {
	ID    int
	Title string
	Price float32
}

var (
	items  []*Item
	nextID = 1
)

func GetItems() []*Item {
	/*
		Esto es data de prueba, la idea es contruir un restClient para que me traiga
		estos datos de los endpoints que estan en la guia
	*/
	item1 := Item{ID: 1, Title: "Bannana", Price: 2.50}
	item2 := Item{ID: 2, Title: "Apple", Price: 3.20}
	item3 := Item{ID: 3, Title: "Cookies", Price: 10.40}
	item4 := Item{ID: 4, Title: "Noodles", Price: 23.50}
	item5 := Item{ID: 5, Title: "Olive Oil", Price: 13.00}
	item6 := Item{ID: 6, Title: "Water", Price: 0.50}
	item7 := Item{ID: 7, Title: "Beer", Price: 1.50}
	item8 := Item{ID: 8, Title: "Vodka", Price: 10.50}
	item9 := Item{ID: 9, Title: "Bread", Price: 0.20}
	item10 := Item{ID: 10, Title: "Grapes", Price: 0.50}
	item11 := Item{ID: 11, Title: "Rice", Price: 3.50}
	item12 := Item{ID: 12, Title: "Pizza", Price: 13.10}
	items = append(items, &item1, &item2, &item3, &item4, &item5, &item6, &item7, &item8, &item9, &item10, &item11, &item12)
	return items
}

func GetItemById(idItem int) (Item, error) {
	for _, item := range items {
		if item.ID == idItem {
			return *item, nil
		}
	}
	return Item{}, fmt.Errorf("Item with ID '%v' not found", idItem)
}
