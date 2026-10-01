package demo

import (
	"fmt"
	"io"

	"vyshkat/internal/domain"
	"vyshkat/internal/park"
	"vyshkat/internal/reports"
)

func Run(reception *park.Reception, p *park.Park, service *reports.Service, out io.Writer) {
	fmt.Fprintln(out, "ВышКат: учёт парка")

	fmt.Fprintln(out, "\nПриём транспорта (техосмотр)")
	scooter1, err := domain.NewScooter("SC-1", "Самокат Дедлайн", 8, 1.0)
	if err != nil {
		panic(err)
	}
	registerVehicle(reception, out, scooter1)

	scooter2, err := domain.NewScooter("SC-2", "Самокат Энергозатратный", 7, 3.5)
	if err != nil {
		panic(err)
	}
	registerVehicle(reception, out, scooter2)

	bicycle1, err := domain.NewBicycle("BC-1", "Велосипед городской", 9, true)
	if err != nil {
		panic(err)
	}
	registerVehicle(reception, out, bicycle1)

	ebike1, err := domain.NewEBike("EB-1", "Электровелосипед", 5, 1.8)
	if err != nil {
		panic(err)
	}
	registerVehicle(reception, out, ebike1)

	fmt.Fprintln(out, "\nПриём вещей (без техосмотра)")

	helmet, err := domain.NewHelmet("HM-1", "Шлем спортивный", "XL")
	if err != nil {
		panic(err)
	}
	registerItem(p, out, helmet)

	cabinet, err := domain.NewChargingCabinet("CC-1", "Зарядный шкаф", 0.5)
	if err != nil {
		panic(err)
	}
	registerItem(p, out, cabinet)

	fmt.Fprintln(out, "\nОтчёты")
	fmt.Fprintf(out, "Всего единиц в парке: %d\n", service.ParkCount())
	fmt.Fprintf(out, "Суммарное суточное энергопотребление: %.1f кВт*ч\n", service.SumEnergy())

	fmt.Fprintln(out, "\nТехника для новичков:")
	for _, v := range service.NoviceFriendlyVehicles() {
		fmt.Fprintf(out, "-- %s (%s)\n", v.Name(), v.InventoryNumber())
	}

	fmt.Fprintln(out, "\nПолный инвентарный список:")
	for _, inv := range service.FullInventory() {
		fmt.Fprintf(out, "  - %s: %s\n", inv.InventoryNumber(), inv.Name())
	}
}

func registerVehicle(reception *park.Reception, out io.Writer, v domain.Vehicle) {
	verdict, err := reception.Register(v)
	if err != nil {
		panic(err)
	}
	if verdict.Accepted {
		fmt.Fprintf(out, "%s (%s): ПРИНЯТ - %s\n", v.Name(), v.InventoryNumber(), verdict.Reason)
	} else {
		fmt.Fprintf(out, "%s (%s): ОТКЛОНЁН - %s\n", v.Name(), v.InventoryNumber(), verdict.Reason)
	}
}

func registerItem(p *park.Park, out io.Writer, it domain.Item) {
	err := p.AddItem(it)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(out, "%s (%s): добавлен\n", it.Name(), it.InventoryNumber())
}
