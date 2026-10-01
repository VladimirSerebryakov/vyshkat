package domain

// Тип велосипедов, требуется добавить если создаём новый тип транспорта
type VehicleKind string

const (
	KindScooter VehicleKind = "scooter"
	KindBicycle VehicleKind = "bicycle"
	KindEBike   VehicleKind = "e-bike"
)

// Интерфейс описывающий транспорт
// Инвентаризуем, имеет описанную сложность для новичков, и имеет тип
type Vehicle interface {
	Inventoried
	NoviceRated
	Kind() VehicleKind
}
