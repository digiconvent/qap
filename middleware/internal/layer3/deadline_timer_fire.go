package layer3

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func (t *DeadlineTimer) fire(id string) {
	t.mutex.Lock()
	delete(t.timers, id)
	t.mutex.Unlock()

	s, ok := t.dependencies.Get(id)
	if !ok || s.Resolution != concepts.Pending {
		return
	}
	now := t.dependencies.Now().Unix()
	outcome := t.dependencies.Evaluate(s, s.VoterCount, now)
	switch {
	case outcome == concepts.Executed:
		t.dependencies.Execute(s)
	case outcome == concepts.Pending && s.Policy.ExtensionSeconds != nil && !layer4.ExtensionPassed(s.When, s.Policy.DeadlineSeconds, s.Policy.ExtensionSeconds, now):
		t.Start(s)
	default:
		t.dependencies.Reject(s)
	}
}
