package park

import (
	"fmt"
	"vyshkat/internal/domain"
)

// Хранит весь транспорт и вещи
type Park struct {
	vehicles []domain.Vehicle
	items    []domain.Item
}

func NewPark() *Park {
	return &Park{}
}

func (p *Park) Vehicles() []domain.Vehicle {
	result := make([]domain.Vehicle, len(p.vehicles))
	copy(result, p.vehicles)
	return result
	// Реализовал через copy для инкапсуляции, из вне можно читать корректно
	// Но попытка изменить из вне не изменит наш Park
}

func (p *Park) Items() []domain.Item {
	result := make([]domain.Item, len(p.items))
	copy(result, p.items)
	return result
}

// Инкапсулируем, вызвать можно будет только из пакета park, вызывать будет Reception после валидации
func (p *Park) addVehicle(v domain.Vehicle) error {
	err := p.checkInventoryNumberUnique(v)
	if err != nil {
		return err
	}
	p.vehicles = append(p.vehicles, v)
	return nil
}

// Вещи не проходят техосмотр, добавляем сразу
func (p *Park) AddItem(it domain.Item) error {
	err := p.checkInventoryNumberUnique(it)
	if err != nil {
		return err
	}
	p.items = append(p.items, it)
	return nil
}

func (p *Park) checkInventoryNumberUnique(it domain.Inventoried) error {
	num := it.InventoryNumber()
	for _, i := range p.items {
		if i.InventoryNumber() == num {
			return fmt.Errorf("inventory number %q is already used", num)
		}
	}
	for _, v := range p.vehicles {
		if v.InventoryNumber() == num {
			return fmt.Errorf("inventory number %q is already used", num)
		}
	}
	return nil
}
