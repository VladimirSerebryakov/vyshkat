package domain_test

import (
	"testing"

	"vyshkat/internal/domain"
)

func TestNewScooter(t *testing.T) {
	cases := []struct {
		testName    string
		number      string
		name        string
		simplicity  int
		dailyEnergy float64
		needErr     bool
	}{
		{"valid", "SC-1", "Самокат", 8, 1.0, false},
		{"empty num", "", "Самокат", 8, 1.0, true},
		{"empty name", "SC-2", "", 8, 1.0, true},
		{"simplicity too low", "SC-3", "Самокат", 0, 1.0, true},
		{"simplicity too high", "SC-4", "Самокат", 11, 1.0, true},
		{"zero energy", "SC-5", "Самокат", 8, 0, true},
		{"negative energy", "SC-6", "Самокат", 8, -1, true},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			s, err := domain.NewScooter(tc.number, tc.name, tc.simplicity, tc.dailyEnergy)
			if tc.needErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if s.Simplicity() != tc.simplicity {
				t.Errorf("Simplicity() = %d, want %d", s.Simplicity(), tc.simplicity)
			}
			if s.DailyEnergyKWh() != tc.dailyEnergy {
				t.Errorf("DailyEnergyKWh() = %v, want %v", s.DailyEnergyKWh(), tc.dailyEnergy)
			}
			if s.Kind() != domain.KindScooter {
				t.Errorf("Kind() = %v, want %v", s.Kind(), domain.KindScooter)
			}
		})
	}
}

func TestNewBicycle(t *testing.T) {
	cases := []struct {
		testName   string
		number     string
		name       string
		simplicity int
		withBag    bool
		needErr    bool
	}{
		{"valid", "BK-1", "Велосипед", 9, true, false},
		{"empty num", "", "Велосипед", 9, true, true},
		{"empty name", "BK-2", "", 9, true, true},
		{"simplicity too low", "BK-3", "Велосипед", 0, false, true},
		{"simplicity too high", "BK-4", "Велосипед", 11, false, true},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			b, err := domain.NewBicycle(tc.number, tc.name, tc.simplicity, tc.withBag)
			if tc.needErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if b.Simplicity() != tc.simplicity {
				t.Errorf("Simplicity() = %d, want %d", b.Simplicity(), tc.simplicity)
			}
			if b.HaveBag() != tc.withBag {
				t.Errorf("HaveBag() = %v, want %v", b.HaveBag(), tc.withBag)
			}
			if b.Kind() != domain.KindBicycle {
				t.Errorf("Kind() = %v, want %v", b.Kind(), domain.KindBicycle)
			}
		})
	}
}

func TestNewEBike(t *testing.T) {
	cases := []struct {
		testName    string
		number      string
		name        string
		simplicity  int
		dailyEnergy float64
		needErr     bool
	}{
		{"valid", "EB-1", "Электровелосипед", 6, 1.8, false},
		{"empty num", "", "Электровелосипед", 6, 1.8, true},
		{"empty name", "EB-2", "", 6, 1.8, true},
		{"simplicity too low", "EB-3", "Электровелосипед", 0, 1.8, true},
		{"simplicity too high", "EB-4", "Электровелосипед", 11, 1.8, true},
		{"zero energy", "EB-5", "Электровелосипед", 6, 0, true},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			e, err := domain.NewEBike(tc.number, tc.name, tc.simplicity, tc.dailyEnergy)
			if tc.needErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if e.Kind() != domain.KindEBike {
				t.Errorf("Kind() = %v, want %v", e.Kind(), domain.KindEBike)
			}
		})
	}
}

func TestBicycle_DoesNotImplementEnergyConsumer(t *testing.T) {
	b, err := domain.NewBicycle("BK-1", "Велосипед", 9, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := any(b).(domain.EnergyConsumer); ok {
		t.Errorf("Bicycle must not implement EnergyConsumer")
	}
}

func TestScooterAndEBike_ImplementEnergyConsumer(t *testing.T) {
	s, err := domain.NewScooter("SC-1", "Самокат", 8, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := any(s).(domain.EnergyConsumer); !ok {
		t.Errorf("Scooter must implement EnergyConsumer")
	}

	e, err := domain.NewEBike("EB-1", "Электровелосипед", 6, 1.8)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := any(e).(domain.EnergyConsumer); !ok {
		t.Errorf("EBike must implement EnergyConsumer")
	}
}
