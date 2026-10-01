package reports_test

import (
	"testing"

	"vyshkat/internal/domain"
	"vyshkat/internal/reports"
)

// fakeParkReader — тестовая подмена reports.ParkReader с заранее заданным содержимым,
// без реального park.Park.
type fakeParkReader struct {
	vehicles []domain.Vehicle
	items    []domain.Item
}

func (f fakeParkReader) Vehicles() []domain.Vehicle {
	return f.vehicles
}

func (f fakeParkReader) Items() []domain.Item {
	return f.items
}

func newTestService(t *testing.T) *reports.Service {
	t.Helper()

	scooter, err := domain.NewScooter("SC-1", "Самокат", 8, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bicycle, err := domain.NewBicycle("BK-1", "Велосипед", 3, true) // simplicity < 6 — не для новичков
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	cabinet, err := domain.NewChargingCabinet("CC-1", "Зарядный шкаф", 0.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	reader := fakeParkReader{
		vehicles: []domain.Vehicle{scooter, bicycle},
		items:    []domain.Item{cabinet},
	}

	svc, err := reports.NewService(reader)
	if err != nil {
		t.Fatalf("NewService failed: %v", err)
	}
	return svc
}

func TestService_ParkCount(t *testing.T) {
	svc := newTestService(t)
	if got := svc.ParkCount(); got != 3 {
		t.Errorf("ParkCount() = %d, want 3", got)
	}
}

func TestService_SumEnergy(t *testing.T) {
	svc := newTestService(t)
	want := 1.5 // 1.0 (самокат) + 0.5 (шкаф), велосипед не потребляет энергию
	if got := svc.SumEnergy(); got != want {
		t.Errorf("SumEnergy() = %v, want %v", got, want)
	}
}

func TestService_NoviceFriendlyVehicles(t *testing.T) {
	svc := newTestService(t)
	got := svc.NoviceFriendlyVehicles()
	if len(got) != 1 {
		t.Fatalf("expected 1 novice-friendly vehicle, got %d", len(got))
	}
	if got[0].InventoryNumber() != "SC-1" {
		t.Errorf("unexpected vehicle in result: %s", got[0].InventoryNumber())
	}
}

func TestService_FullInventory(t *testing.T) {
	svc := newTestService(t)
	got := svc.FullInventory()
	if len(got) != 3 {
		t.Fatalf("expected 3 inventory lines, got %d", len(got))
	}
}

func TestNewService_NilReader(t *testing.T) {
	_, err := reports.NewService(nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
