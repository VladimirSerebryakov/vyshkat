package domain

import (
	"errors"
)

type EBike struct {
	InventoryItem
	simplicity     int
	dailyEnergyKWh float64
}

func NewEBike(number, name string, simplicity int, dailyEnergy float64) (*EBike, error) {
	item, err := NewInventoryItem(number, name)
	if err != nil {
		return nil, err
	}
	if simplicity < 1 || simplicity > 10 {
		return nil, errors.New("simplicity score must be between 1 and 10")
	}
	if dailyEnergy <= 0 {
		return nil, errors.New("daily energy must be greater than 0")
	}
	return &EBike{InventoryItem: item, simplicity: simplicity, dailyEnergyKWh: dailyEnergy}, nil
}

func (e *EBike) Simplicity() int {
	return e.simplicity
}

func (e *EBike) DailyEnergyKWh() float64 {
	return e.dailyEnergyKWh
}

func (e *EBike) Kind() VehicleKind {
	return KindEBike
}
