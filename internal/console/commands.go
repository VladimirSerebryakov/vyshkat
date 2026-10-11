package console

import (
	"fmt"
	"io"
	"strconv"

	"vyshkat/internal/domain"
	"vyshkat/internal/park"
	"vyshkat/internal/reports"
)

const helpText = `Команды:
  add-scooter <номер> <название> <простота 1-10> <кВт·ч в сутки>
  add-bicycle <номер> <название> <простота 1-10> <сумка true|false>
  add-ebike   <номер> <название> <простота 1-10> <кВт·ч в сутки>
  add-helmet  <номер> <название> <размер>
  add-cabinet <номер> <название> <кВт·ч в сутки>
  report-count       количество единиц в парке
  report-energy      суммарное суточное энергопотребление
  report-novices     техника для новичков
  report-inventory   наименования и инвентарные номера всей техники и вещей
  help
  exit
Названия из нескольких слов берите в кавычки: "Самокат Дедлайн".`

type session struct {
	reception *park.Reception
	p         *park.Park
	reports   *reports.Service
	out       io.Writer
}

func (s *session) execute(args []string) (exit bool, err error) {
	name, params := args[0], args[1:]

	switch name {
	case "add-scooter":
		return false, s.addScooter(params)
	case "add-bicycle":
		return false, s.addBicycle(params)
	case "add-ebike":
		return false, s.addEBike(params)
	case "add-helmet":
		return false, s.addHelmet(params)
	case "add-cabinet":
		return false, s.addCabinet(params)
	case "report-count":
		fmt.Fprintf(s.out, "Всего единиц в парке: %d\n", s.reports.ParkCount())
		return false, nil
	case "report-energy":
		fmt.Fprintf(s.out, "Суммарное суточное энергопотребление: %.2f кВт·ч\n", s.reports.SumEnergy())
		return false, nil
	case "report-novices":
		s.reportNovices()
		return false, nil
	case "report-inventory":
		s.reportInventory()
		return false, nil
	case "help":
		fmt.Fprintln(s.out, helpText)
		return false, nil
	case "exit", "quit":
		return true, nil
	default:
		return false, fmt.Errorf("неизвестная команда %q, введите help", name)
	}
}

func (s *session) addScooter(params []string) error {
	if err := expectArgs(params, 4, "add-scooter <номер> <название> <простота> <кВт·ч>"); err != nil {
		return err
	}
	simplicity, err := parseInt(params[2], "простота")
	if err != nil {
		return err
	}
	energy, err := parseFloat(params[3], "энергопотребление")
	if err != nil {
		return err
	}
	v, err := domain.NewScooter(params[0], params[1], simplicity, energy)
	if err != nil {
		return err
	}
	return s.registerVehicle(v)
}

func (s *session) addBicycle(params []string) error {
	if err := expectArgs(params, 4, "add-bicycle <номер> <название> <простота> <сумка true|false>"); err != nil {
		return err
	}
	simplicity, err := parseInt(params[2], "простота")
	if err != nil {
		return err
	}
	withBag, err := strconv.ParseBool(params[3])
	if err != nil {
		return fmt.Errorf("сумка: ожидается true или false, получено %q", params[3])
	}
	v, err := domain.NewBicycle(params[0], params[1], simplicity, withBag)
	if err != nil {
		return err
	}
	return s.registerVehicle(v)
}

func (s *session) addEBike(params []string) error {
	if err := expectArgs(params, 4, "add-ebike <номер> <название> <простота> <кВт·ч>"); err != nil {
		return err
	}
	simplicity, err := parseInt(params[2], "простота")
	if err != nil {
		return err
	}
	energy, err := parseFloat(params[3], "энергопотребление")
	if err != nil {
		return err
	}
	v, err := domain.NewEBike(params[0], params[1], simplicity, energy)
	if err != nil {
		return err
	}
	return s.registerVehicle(v)
}

func (s *session) addHelmet(params []string) error {
	if err := expectArgs(params, 3, "add-helmet <номер> <название> <размер>"); err != nil {
		return err
	}
	it, err := domain.NewHelmet(params[0], params[1], params[2])
	if err != nil {
		return err
	}
	return s.addItem(it)
}

func (s *session) addCabinet(params []string) error {
	if err := expectArgs(params, 3, "add-cabinet <номер> <название> <кВт·ч>"); err != nil {
		return err
	}
	energy, err := parseFloat(params[2], "энергопотребление")
	if err != nil {
		return err
	}
	it, err := domain.NewChargingCabinet(params[0], params[1], energy)
	if err != nil {
		return err
	}
	return s.addItem(it)
}

func (s *session) registerVehicle(v domain.Vehicle) error {
	verdict, err := s.reception.Register(v)
	if err != nil {
		return err
	}
	if verdict.Accepted {
		fmt.Fprintf(s.out, "принят: %s (%s) — %s\n", v.Name(), v.InventoryNumber(), verdict.Reason)
	} else {
		fmt.Fprintf(s.out, "отклонён техосмотром: %s (%s) — %s\n", v.Name(), v.InventoryNumber(), verdict.Reason)
	}
	return nil
}

func (s *session) addItem(it domain.Item) error {
	if err := s.p.AddItem(it); err != nil {
		return err
	}
	fmt.Fprintf(s.out, "поставлен на баланс: %s (%s)\n", it.Name(), it.InventoryNumber())
	return nil
}

func (s *session) reportNovices() {
	vehicles := s.reports.NoviceFriendlyVehicles()
	if len(vehicles) == 0 {
		fmt.Fprintln(s.out, "Техники для новичков нет")
		return
	}
	fmt.Fprintln(s.out, "Техника для новичков:")
	for _, v := range vehicles {
		fmt.Fprintf(s.out, "  - %s: %s\n", v.InventoryNumber(), v.Name())
	}
}

func (s *session) reportInventory() {
	inventory := s.reports.FullInventory()
	if len(inventory) == 0 {
		fmt.Fprintln(s.out, "Парк пуст")
		return
	}
	fmt.Fprintln(s.out, "Полный инвентарный список:")
	for _, inv := range inventory {
		fmt.Fprintf(s.out, "  - %s: %s\n", inv.InventoryNumber(), inv.Name())
	}
}

func expectArgs(params []string, n int, usage string) error {
	if len(params) != n {
		return fmt.Errorf("ожидается %d аргумента(ов), получено %d; формат: %s", n, len(params), usage)
	}
	return nil
}

func parseInt(raw, field string) (int, error) {
	v, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("%s: ожидается целое число, получено %q", field, raw)
	}
	return v, nil
}

func parseFloat(raw, field string) (float64, error) {
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: ожидается число, получено %q", field, raw)
	}
	return v, nil
}
