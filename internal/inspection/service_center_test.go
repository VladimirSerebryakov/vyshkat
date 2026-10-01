package inspection_test

import (
	"testing"

	"vyshkat/internal/domain"
	"vyshkat/internal/inspection"
)

func TestServiceCenter_Inspect(t *testing.T) {
	sc := inspection.NewServiceCenter(2.0)

	scooterOK, err := domain.NewScooter("SC-1", "Самокат", 8, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	scooterAtLimit, err := domain.NewScooter("SC-2", "Самокат", 8, 2.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	scooterOverLimit, err := domain.NewScooter("SC-3", "Самокат", 8, 2.1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bicycle, err := domain.NewBicycle("BK-1", "Велосипед", 9, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cases := []struct {
		testName     string
		vehicle      domain.Vehicle
		wantAccepted bool
	}{
		{"within limit", scooterOK, true},
		{"exactly at limit", scooterAtLimit, true},
		{"above limit", scooterOverLimit, false},
		{"does not consume energy", bicycle, true},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			verdict := sc.Inspect(tc.vehicle)
			if verdict.Accepted != tc.wantAccepted {
				t.Errorf("Accepted = %v, want %v (reason: %s)", verdict.Accepted, tc.wantAccepted, verdict.Reason)
			}
		})
	}
}
