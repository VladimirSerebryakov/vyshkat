package di_test

import (
	"testing"

	"vyshkat/internal/di"
)

type testService struct {
	id int
}

func TestContainer_MustResolve_ReturnsSameInstance(t *testing.T) {
	c := di.New()
	callCount := 0

	di.Register(c, func(c *di.Container) *testService {
		callCount++
		return &testService{id: callCount}
	})

	first := di.MustResolve[*testService](c)
	second := di.MustResolve[*testService](c)

	if callCount != 1 {
		t.Errorf("factory called %d times, want 1 (singleton)", callCount)
	}
	if first != second {
		t.Errorf("MustResolve returned different instances, want same pointer")
	}
}

func TestContainer_RegisterInstance_OverridesFactory(t *testing.T) {
	c := di.New()

	di.Register(c, func(c *di.Container) *testService {
		t.Fatalf("factory must not be called when an instance is pre-registered")
		return nil
	})

	override := &testService{id: 42}
	di.RegisterInstance[*testService](c, override)

	got := di.MustResolve[*testService](c)
	if got != override {
		t.Errorf("MustResolve did not return the pre-registered instance")
	}
}

func TestContainer_MustResolve_PanicsWhenNotRegistered(t *testing.T) {
	c := di.New()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, got none")
		}
	}()

	di.MustResolve[*testService](c)
}
