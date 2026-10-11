package inspection

import (
	"vyshkat/internal/domain"
)

type ServiceCenter struct {
	MaxDailyEnergyKWh float64
}

func NewServiceCenter(maxDailyEnergyKWh float64) *ServiceCenter {
	return &ServiceCenter{MaxDailyEnergyKWh: maxDailyEnergyKWh}
}

func (sc *ServiceCenter) Inspect(v domain.Vehicle) Verdict {
	consumer, ok := v.(domain.EnergyConsumer)
	if !ok {
		return Verdict{Accepted: true, Reason: "Транспорт не потребляет электричества"}
	}
	if consumer.DailyEnergyKWh() > sc.MaxDailyEnergyKWh {
		return Verdict{Accepted: false, Reason: "Ежедневное электропотребление транспорта превышает лимит"}
	}
	return Verdict{Accepted: true, Reason: "Транспорт потребляющий электричество"}
}
