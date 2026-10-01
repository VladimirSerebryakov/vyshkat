package domain_test

import (
	"testing"
	"vyshkat/internal/domain"
)

type fakeRated struct {
	simplicity int
}

func (f fakeRated) Simplicity() int {
	return f.simplicity
}

func TestIsNoviceFriendly(t *testing.T) {
	cases := []struct {
		name       string
		simplicity int
		want       bool
	}{
		{"below threshold", 5, false},
		{"exactly threshold", 6, true},
		{"above threshold", 10, true},
		{"minimum value", 1, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := domain.IsNoviceFriendly(fakeRated{simplicity: tc.simplicity})
			if got != tc.want {
				t.Errorf("IsNoviceFriendly(simplicity=%d) = %v, want %v", tc.simplicity, got, tc.want)
			}
		})
	}
}
