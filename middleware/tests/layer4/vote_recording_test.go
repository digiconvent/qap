package layer4_test

import (
	"slices"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// a new vote replaces the earlier vote of the same person, so they appear in exactly one list
func TestRecordVote(t *testing.T) {
	s := concepts.Submission{Favour: []string{"alice", "bob"}}
	s = layer4.RecordVote(s, "bob", concepts.Against)
	s = layer4.RecordVote(s, "carol", concepts.Withhold)
	s = layer4.RecordVote(s, "carol", concepts.Favour)

	if !slices.Equal(s.Favour, []string{"alice", "carol"}) || !slices.Equal(s.Against, []string{"bob"}) || len(s.Withhold) != 0 {
		t.Fatalf("got %+v", s)
	}
}
