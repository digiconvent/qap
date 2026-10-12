package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (m *Store) change(id string, pendingOnly bool, apply func(concepts.Submission) concepts.Submission) (concepts.Submission, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	s, ok := m.submissions[id]
	if !ok {
		return concepts.Submission{}, concepts.ErrUnknownSubmission
	}
	if pendingOnly && s.Resolution != concepts.Pending {
		return concepts.Submission{}, concepts.ErrAlreadyResolved
	}
	s = apply(clone(s))
	m.submissions[id] = s
	return clone(s), nil
}
