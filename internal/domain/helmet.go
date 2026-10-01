package domain

import (
	"errors"
)

type Helmet struct {
	InventoryItem
	size string
}

func NewHelmet(number, name, size string) (*Helmet, error) {
	item, err := NewInventoryItem(number, name)
	if err != nil {
		return nil, err
	}
	if size == "" {
		return nil, errors.New("helmet size must not be empty")
	}
	return &Helmet{InventoryItem: item, size: size}, nil
}
