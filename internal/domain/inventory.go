package domain

import (
	"errors"
)

// Предмет из инвентаря (number - уникален, проверка вынесена в reception, name - может повторяться)
type InventoryItem struct {
	number string
	name   string
}

// Интерфейс описывающие инвентаризуемые предметы, ему удовлетворяет InventoryItem
type Inventoried interface {
	InventoryNumber() string
	Name() string
}

func NewInventoryItem(number, name string) (InventoryItem, error) {
	if number == "" {
		return InventoryItem{}, errors.New("inventory number must not be empty")
	}
	if name == "" {
		return InventoryItem{}, errors.New("inventory name must not be empty")
	}
	return InventoryItem{number: number, name: name}, nil
}

func (item InventoryItem) InventoryNumber() string {
	return item.number
}

func (item InventoryItem) Name() string {
	return item.name
}
