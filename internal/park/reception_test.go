package park_test

import (
	"testing"

	"vyshkat/internal/domain"
	"vyshkat/internal/inspection"
	"vyshkat/internal/park"
)

// fakeInspector — тестовая подмена inspection.Inspector с заранее заданным вердиктом.
// Используется во всех тестах пакета park_test (и в reception_test.go, и в park_test.go).
type fakeInspector struct {
	verdict inspection.Verdict
}

func (f fakeInspector) Inspect(v domain.Vehicle) inspection.Verdict {
	return f.verdict
}

func TestReception_Register_Accepted_AddsVehicleToPark(t *testing.T) {
	p := park.NewPark()
	fake := fakeInspector{verdict: inspection.Verdict{Accepted: true, Reason: "ok"}}
	r, err := park.NewReception(fake, p)
	if err != nil {
		t.Fatalf("NewReception failed: %v", err)
	}

	scooter, err := domain.NewScooter("SC-1", "Test Scooter", 8, 1.0)
	if err != nil {
		t.Fatalf("NewScooter failed: %v", err)
	}

	verdict, err := r.Register(scooter)
	if err != nil {
		t.Fatalf("Register returned unexpected error: %v", err)
	}
	if !verdict.Accepted {
		t.Errorf("expected Accepted=true, got false")
	}
	if len(p.Vehicles()) != 1 {
		t.Errorf("expected 1 vehicle in park, got %d", len(p.Vehicles()))
	}
}

func TestReception_Register_Rejected_DoesNotAddVehicle(t *testing.T) {
	p := park.NewPark()
	fake := fakeInspector{verdict: inspection.Verdict{Accepted: false, Reason: "rejected by fake"}}
	r, err := park.NewReception(fake, p)
	if err != nil {
		t.Fatalf("NewReception failed: %v", err)
	}

	scooter, err := domain.NewScooter("SC-1", "Test Scooter", 8, 1.0)
	if err != nil {
		t.Fatalf("NewScooter failed: %v", err)
	}

	verdict, err := r.Register(scooter)
	if err != nil {
		t.Fatalf("Register returned unexpected error: %v", err)
	}
	if verdict.Accepted {
		t.Errorf("expected Accepted=false")
	}
	if len(p.Vehicles()) != 0 {
		t.Errorf("expected 0 vehicles in park, got %d", len(p.Vehicles()))
	}
}

func TestNewReception_NilDependencies(t *testing.T) {
	validPark := park.NewPark()
	validInspector := fakeInspector{}

	cases := []struct {
		testName  string
		inspector inspection.Inspector
		p         *park.Park
	}{
		{"nil inspector", nil, validPark},
		{"nil park", validInspector, nil},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := park.NewReception(tc.inspector, tc.p)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}
