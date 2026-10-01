package reports

import (
	"vyshkat/internal/domain"
)

// Интерфейс для чтения parkа, сам Park удовлетворяет его, но интерфейс удобен для подмены в тестах (Mock)
type ParkReader interface {
	Vehicles() []domain.Vehicle
	Items() []domain.Item
}
