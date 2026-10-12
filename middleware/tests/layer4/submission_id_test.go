package layer4_test

import (
	"regexp"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// a submission id is qap_ followed by 32 random hexadecimal characters, and never repeats
func TestNew(t *testing.T) {
	format := regexp.MustCompile("^qap_[0-9a-f]{32}$")
	seen := map[string]bool{}
	for range 100 {
		id := layer4.NewSubmissionId()
		if !format.MatchString(id) || seen[id] {
			t.Fatalf("unexpected id %q", id)
		}
		seen[id] = true
	}
}
