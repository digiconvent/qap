package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// a decided submission expires once nobody has polled it for one minute, and a pending one never expires
func TestExpired(t *testing.T) {
	decided := concepts.Submission{Resolution: concepts.Executed, LastPolled: 1000}
	pending := concepts.Submission{Resolution: concepts.Pending, LastPolled: 0}
	if layer4.Expired(decided, 1059) || !layer4.Expired(decided, 1060) || layer4.Expired(pending, 99999) {
		t.Fatal("expiry is wrong")
	}
}
