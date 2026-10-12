package layer3_test

import (
	"testing"
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

// on every tick, cleanup removes the decided submissions nobody has polled for a minute
func TestCleanup(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))
	s.Resolve("qap_1", 0, concepts.Executed, 0)
	ticks := make(chan time.Time)
	layer3.StartCleanup(s, ticks, func() time.Time { return time.Unix(60, 0) })
	ticks <- time.Now()
	ticks <- time.Now()
	close(ticks)

	if _, ok := s.Get("qap_1"); ok {
		t.Fatal("expected the expired submission to be cleaned up")
	}
}
