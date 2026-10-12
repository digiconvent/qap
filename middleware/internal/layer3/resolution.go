package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (m *Store) Resolve(id string, revision int, resolution string, at int64) (concepts.Submission, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	s, ok := m.submissions[id]
	switch {
	case !ok:
		return concepts.Submission{}, concepts.ErrUnknownSubmission
	case s.Resolution != concepts.Pending:
		return concepts.Submission{}, concepts.ErrAlreadyResolved
	case s.Revision != revision:
		return concepts.Submission{}, concepts.ErrOutdated
	}
	s = clone(s)
	s.Resolution = resolution
	s.LastPolled = at
	m.submissions[id] = s
	return clone(s), nil
}
