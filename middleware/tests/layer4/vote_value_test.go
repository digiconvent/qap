package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// a positive number is favour, zero is withhold, a negative number is against, anything else is invalid
func TestReadVote(t *testing.T) {
	cases := map[string]string{"1": concepts.Favour, "3": concepts.Favour, "0": concepts.Withhold, "-1": concepts.Against, "-7": concepts.Against}
	for header, want := range cases {
		if got, ok := layer4.ReadVote(header); !ok || got != want {
			t.Errorf("%s: got %q %v", header, got, ok)
		}
	}
	for _, header := range []string{"", "x", "1.5", "yes"} {
		if _, ok := layer4.ReadVote(header); ok {
			t.Errorf("%q: expected invalid", header)
		}
	}
}
