package layer3_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

func number(n int) *int {
	return &n
}

func votes(n int) []string {
	return make([]string, n)
}

// the poll evaluation checks against, then quorum and favour, then the deadline and its extension
func TestEvaluate(t *testing.T) {
	cases := []struct {
		name       string
		policy     concepts.Policy
		favour     int
		against    int
		withhold   int
		now        int64
		voterCount int
		want       string
	}{
		{"favour below threshold", concepts.Policy{FavourThreshold: 2}, 1, 0, 0, 0, 5, concepts.Pending},
		{"favour meets threshold", concepts.Policy{FavourThreshold: 2}, 2, 0, 0, 0, 5, concepts.Executed},
		{"against meets threshold", concepts.Policy{FavourThreshold: 2, AgainstThreshold: number(1)}, 1, 1, 0, 0, 5, concepts.Rejected},
		{"against wins over favour", concepts.Policy{FavourThreshold: 1, AgainstThreshold: number(1)}, 1, 1, 0, 0, 5, concepts.Rejected},
		{"unsatisfiable against threshold", concepts.Policy{FavourThreshold: 9, AgainstThreshold: number(9)}, 0, 2, 0, 0, 5, concepts.Pending},
		{"quorum not met", concepts.Policy{FavourThreshold: 1, Quorum: number(3)}, 1, 0, 2, 0, 5, concepts.Pending},
		{"quorum met", concepts.Policy{FavourThreshold: 1, Quorum: number(2)}, 1, 1, 0, 0, 5, concepts.Executed},
		{"percentage favour threshold", concepts.Policy{FavourThreshold: -50}, 2, 0, 0, 0, 4, concepts.Executed},
		{"a favour threshold above the members is met by every member", concepts.Policy{FavourThreshold: 9}, 5, 0, 0, 0, 5, concepts.Executed},
		{"a favour threshold above the members still needs every member", concepts.Policy{FavourThreshold: 9}, 4, 0, 0, 0, 5, concepts.Pending},
		{"an against threshold of 0 disables the veto", concepts.Policy{FavourThreshold: 2, AgainstThreshold: number(0)}, 1, 3, 0, 0, 5, concepts.Pending},
		{"deadline passed while pending", concepts.Policy{FavourThreshold: 2, DeadlineSeconds: number(60)}, 1, 0, 0, 60, 5, concepts.Rejected},
		{"criteria met at the deadline", concepts.Policy{FavourThreshold: 2, DeadlineSeconds: number(60)}, 2, 0, 0, 60, 5, concepts.Executed},
		{"deadline not passed", concepts.Policy{FavourThreshold: 2, DeadlineSeconds: number(60)}, 1, 0, 0, 59, 5, concepts.Pending},
		{"silence at the deadline extends", concepts.Policy{FavourThreshold: 2, DeadlineSeconds: number(60), ExtensionSeconds: number(30)}, 1, 0, 0, 60, 5, concepts.Pending},
		{"silence until the end of the extension passes", concepts.Policy{FavourThreshold: 2, DeadlineSeconds: number(60), ExtensionSeconds: number(30)}, 1, 0, 0, 90, 5, concepts.Executed},
		{"silence passes despite a quorum", concepts.Policy{FavourThreshold: 2, Quorum: number(4), DeadlineSeconds: number(60), ExtensionSeconds: number(30)}, 1, 0, 0, 90, 5, concepts.Executed},
		{"a vote against ends the extension", concepts.Policy{FavourThreshold: 2, DeadlineSeconds: number(60), ExtensionSeconds: number(30)}, 1, 1, 0, 70, 5, concepts.Rejected},
		{"a vote against before the deadline means no extension", concepts.Policy{FavourThreshold: 2, DeadlineSeconds: number(60), ExtensionSeconds: number(30)}, 1, 1, 0, 60, 5, concepts.Rejected},
		{"the threshold met during the extension executes", concepts.Policy{FavourThreshold: 2, DeadlineSeconds: number(60), ExtensionSeconds: number(30)}, 2, 0, 0, 70, 5, concepts.Executed},
		{"an extension without a deadline changes nothing", concepts.Policy{FavourThreshold: 2, ExtensionSeconds: number(30)}, 1, 0, 0, 9999, 5, concepts.Pending},
	}
	for _, c := range cases {
		s := concepts.Submission{Policy: c.policy, Favour: votes(c.favour), Against: votes(c.against), Withhold: votes(c.withhold)}
		if got := layer3.EvaluatePoll(s, c.voterCount, c.now); got != c.want {
			t.Errorf("%s: got %s, want %s", c.name, got, c.want)
		}
	}
}
