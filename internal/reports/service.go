package reports

import (
	"errors"
	"vyshkat/internal/domain"
)

type Service struct {
	reader ParkReader
}

func NewService(reader ParkReader) (*Service, error) {
	if reader == nil {
		return nil, errors.New("reader must not be nil")
	}
	return &Service{reader: reader}, nil
}

func (s *Service) ParkCount() int {
	items := s.FullInventory()
	return len(items)
}

func (s *Service) SumEnergy() float64 {
	items := s.FullInventory()
	var res float64 = 0
	for _, item := range items {
		consumer, ok := item.(domain.EnergyConsumer)
		if ok {
			res += consumer.DailyEnergyKWh()
		}
	}
	return res
}

func (s *Service) NoviceFriendlyVehicles() []domain.Vehicle {
	var vehicles []domain.Vehicle = s.reader.Vehicles()
	res := make([]domain.Vehicle, 0)
	for _, v := range vehicles {
		friendly := domain.IsNoviceFriendly(v)
		if friendly {
			res = append(res, v)
		}
	}
	return res
}
