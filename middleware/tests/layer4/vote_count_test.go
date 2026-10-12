package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// votes are counted per list, and participation is favour plus against
func TestCountVotes(t *testing.T) {
	s := concepts.Submission{Favour: []string{"a", "b"}, Against: []string{"c"}, Withhold: []string{"d", "e", "f"}}
	want := layer4.Counts{Favour: 2, Against: 1, Withhold: 3, Participation: 3}
	if got := layer4.CountVotes(s); got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
