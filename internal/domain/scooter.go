package domain

import (
	"errors"
)

type Scooter struct {
	InventoryItem
	simplicity     int
	dailyEnergyKWh float64
}

func NewScooter(number, name string, simplicity int, dailyEnergy float64) (*Scooter, error) {
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
	return &Scooter{InventoryItem: item, simplicity: simplicity, dailyEnergyKWh: dailyEnergy}, nil
}

func (s *Scooter) Simplicity() int {
	return s.simplicity
}

func (s *Scooter) DailyEnergyKWh() float64 {
	return s.dailyEnergyKWh
}

func (s *Scooter) Kind() VehicleKind {
	return KindScooter
}
