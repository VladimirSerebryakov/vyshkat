package domain

// Интерфейс для потребителей электроэнергии
type EnergyConsumer interface {
	DailyEnergyKWh() float64
}
