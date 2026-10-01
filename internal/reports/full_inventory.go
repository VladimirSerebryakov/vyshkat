package reports

import (
	"vyshkat/internal/domain"
)

func (s *Service) FullInventory() []domain.Inventoried {
	res := make([]domain.Inventoried, 0)
	for _, i := range s.reader.Items() {
		res = append(res, i)
	}
	for _, v := range s.reader.Vehicles() {
		res = append(res, v)
	}
	return res
}
