package domain_test

import (
	"testing"

	"vyshkat/internal/domain"
)

func TestNewHelmet(t *testing.T) {
	cases := []struct {
		testName string
		number   string
		name     string
		size     string
		needErr  bool
	}{
		{"valid", "HM-1", "Шлем спортивный", "XL", false},
		{"empty num", "", "Шлем спортивный", "L", true},
		{"empty name", "HM-3", "", "M", true},
		{"empty size", "HM-4", "Шлем спортивный", "", true},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			h, err := domain.NewHelmet(tc.number, tc.name, tc.size)
			if tc.needErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if h.InventoryNumber() != tc.number {
				t.Errorf("InventoryNumber() = %q, want %q", h.InventoryNumber(), tc.number)
			}
			if h.Name() != tc.name {
				t.Errorf("Name() = %q, want %q", h.Name(), tc.name)
			}
		})
	}
}

func TestNewChargingCabinet(t *testing.T) {
	cases := []struct {
		testName    string
		number      string
		name        string
		dailyEnergy float64
		needErr     bool
	}{
		{"valid", "CC-1", "Зарядный шкаф", 0.5, false},
		{"zero energy allowed", "CC-2", "Зарядный шкаф", 0, false},
		{"negative energy", "CC-3", "Зарядный шкаф", -1, true},
		{"empty num", "", "Зарядный шкаф", 0.5, true},
		{"empty name", "CC-5", "", 0.5, true},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			cc, err := domain.NewChargingCabinet(tc.number, tc.name, tc.dailyEnergy)
			if tc.needErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if cc.DailyEnergyKWh() != tc.dailyEnergy {
				t.Errorf("DailyEnergyKWh() = %v, want %v", cc.DailyEnergyKWh(), tc.dailyEnergy)
			}
		})
	}
}

func TestHelmet_DoesNotImplementEnergyConsumer(t *testing.T) {
	h, err := domain.NewHelmet("HM-1", "Шлем спортивный", "XL")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := any(h).(domain.EnergyConsumer); ok {
		t.Errorf("Helmet must not implement EnergyConsumer")
	}
}

func TestChargingCabinet_ImplementsEnergyConsumer(t *testing.T) {
	cc, err := domain.NewChargingCabinet("CC-1", "Зарядный шкаф", 0.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := any(cc).(domain.EnergyConsumer); !ok {
		t.Errorf("ChargingCabinet must implement EnergyConsumer")
	}
}
