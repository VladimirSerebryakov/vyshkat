package app_test

import (
	"testing"

	"vyshkat/internal/app"
	"vyshkat/internal/di"
	"vyshkat/internal/domain"
	"vyshkat/internal/inspection"
	"vyshkat/internal/park"
)

// fakeInspector — тестовая подмена inspection.Inspector для проверки, что Options
// действительно переопределяет зависимость на уровне DI-контейнера.
type fakeInspector struct {
	verdict inspection.Verdict
}

func (f fakeInspector) Inspect(v domain.Vehicle) inspection.Verdict {
	return f.verdict
}

func TestBuildContainer_DefaultInspector(t *testing.T) {
	c := app.BuildContainer(app.Options{})
	reception := di.MustResolve[*park.Reception](c)

	scooter, err := domain.NewScooter("SC-1", "Самокат", 8, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	verdict, err := reception.Register(scooter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !verdict.Accepted {
		t.Errorf("expected scooter within energy limit to be accepted, reason: %s", verdict.Reason)
	}
}

func TestBuildContainer_OverridesInspectorViaOptions(t *testing.T) {
	fake := fakeInspector{verdict: inspection.Verdict{Accepted: false, Reason: "fake: always reject"}}
	c := app.BuildContainer(app.Options{Inspector: fake})
	reception := di.MustResolve[*park.Reception](c)

	scooter, err := domain.NewScooter("SC-1", "Самокат", 8, 1.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	verdict, err := reception.Register(scooter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if verdict.Accepted {
		t.Errorf("expected fake inspector to reject, but vehicle was accepted")
	}
}
