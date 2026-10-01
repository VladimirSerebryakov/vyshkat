package park_test

import (
	"testing"

	"vyshkat/internal/domain"
	"vyshkat/internal/inspection"
	"vyshkat/internal/park"
)

func TestPark_AddItem_DuplicateInventoryNumberRejected(t *testing.T) {
	p := park.NewPark()

	helmet1, err := domain.NewHelmet("X-1", "Шлем", "L")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := p.AddItem(helmet1); err != nil {
		t.Fatalf("unexpected error adding first item: %v", err)
	}

	helmet2, err := domain.NewHelmet("X-1", "Другой шлем", "M")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := p.AddItem(helmet2); err == nil {
		t.Errorf("expected error for duplicate inventory number, got nil")
	}
}

func TestPark_DuplicateInventoryNumber_AcrossVehicleAndItem(t *testing.T) {
	p := park.NewPark()
	fake := fakeInspector{verdict: inspection.Verdict{Accepted: true, Reason: "ok"}}
	r, err := park.NewReception(fake, p)
	if err != nil {
		t.Fatalf("NewReception failed: %v", err)
	}

	scooter, err := domain.NewScooter("Y-1", "Самокат", 8, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := r.Register(scooter); err != nil {
		t.Fatalf("unexpected error registering scooter: %v", err)
	}

	helmet, err := domain.NewHelmet("Y-1", "Шлем", "L")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := p.AddItem(helmet); err == nil {
		t.Errorf("expected error: item number conflicts with existing vehicle")
	}
}

func TestPark_Items_ReturnsDefensiveCopy(t *testing.T) {
	p := park.NewPark()
	helmet, err := domain.NewHelmet("HM-1", "Шлем", "L")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := p.AddItem(helmet); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := p.Items()
	evil, err := domain.NewHelmet("HM-666", "Подменённый шлем", "XXL")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	items[0] = evil

	if p.Items()[0].Name() == evil.Name() {
		t.Errorf("mutating the returned slice corrupted Park internal state")
	}
}

func TestPark_Vehicles_ReturnsDefensiveCopy(t *testing.T) {
	p := park.NewPark()
	fake := fakeInspector{verdict: inspection.Verdict{Accepted: true, Reason: "ok"}}
	r, err := park.NewReception(fake, p)
	if err != nil {
		t.Fatalf("NewReception failed: %v", err)
	}

	scooter, err := domain.NewScooter("SC-1", "Самокат", 8, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := r.Register(scooter); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	vehicles := p.Vehicles()
	evil, err := domain.NewScooter("SC-666", "Подменённый самокат", 1, 999)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	vehicles[0] = evil

	if p.Vehicles()[0].Name() == evil.Name() {
		t.Errorf("mutating the returned slice corrupted Park internal state")
	}
}
