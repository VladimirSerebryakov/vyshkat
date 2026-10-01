package domain

import (
	"errors"
)

type Bicycle struct {
	InventoryItem
	simplicity int
	withBag    bool
}

func NewBicycle(number string, name string, simplicity int, withBag bool) (*Bicycle, error) {
	item, err := NewInventoryItem(number, name)
	if err != nil {
		return nil, err
	}
	if simplicity < 1 || simplicity > 10 {
		return nil, errors.New("simplicity score must be between 1 and 10")
	}
	return &Bicycle{InventoryItem: item, simplicity: simplicity, withBag: withBag}, nil
}

func (b *Bicycle) Simplicity() int {
	return b.simplicity
}

func (b *Bicycle) Kind() VehicleKind {
	return KindBicycle
}

func (b *Bicycle) HaveBag() bool {
	return b.withBag
}
