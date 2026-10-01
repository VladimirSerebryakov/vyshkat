package domain

import (
	"errors"
)

type ChargingCabinet struct {
	InventoryItem
	dailyEnergyKWh float64
}

func NewChargingCabinet(number, name string, dailyEnergy float64) (*ChargingCabinet, error) {
	item, err := NewInventoryItem(number, name)
	if err != nil {
		return nil, err
	}
	if dailyEnergy < 0 {
		return nil, errors.New("daily energy must not be negative")
	}
	return &ChargingCabinet{InventoryItem: item, dailyEnergyKWh: dailyEnergy}, nil
}

func (cc *ChargingCabinet) DailyEnergyKWh() float64 {
	return cc.dailyEnergyKWh
}
