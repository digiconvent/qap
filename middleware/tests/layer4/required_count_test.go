package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// a positive threshold is a number of votes, a negative threshold a percentage of eligible voters rounded up
func TestRequiredCount(t *testing.T) {
	cases := []struct {
		threshold, eligible, want int
		satisfiable               bool
	}{
		{3, 5, 3, true},
		{3, 3, 3, true},
		{4, 3, 4, false},
		{-50, 5, 3, true},
		{-50, 4, 2, true},
		{-100, 4, 4, true},
		{-1, 7, 1, true},
		{0, 5, 0, true},
	}
	for _, c := range cases {
		if got, ok := layer4.RequiredCount(c.threshold, c.eligible); got != c.want || ok != c.satisfiable {
			t.Errorf("threshold %d of %d: got %d %v", c.threshold, c.eligible, got, ok)
		}
	}
}

// a favour threshold above the number of members needs every member instead
func TestRequiredFavourCount(t *testing.T) {
	cases := []struct{ threshold, eligible, want int }{
		{3, 5, 3},
		{4, 3, 3},
		{-50, 4, 2},
		{-100, 4, 4},
	}
	for _, c := range cases {
		if got := layer4.RequiredFavourCount(c.threshold, c.eligible); got != c.want {
			t.Errorf("threshold %d of %d: got %d, want %d", c.threshold, c.eligible, got, c.want)
		}
	}
}
