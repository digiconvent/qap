package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (m *Store) ListPending(route string, method string) []concepts.Submission {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var pending []concepts.Submission
	for _, id := range m.order {
		s := m.submissions[id]
		if s.Resolution == concepts.Pending && s.Route == route && s.Method == method {
			pending = append(pending, clone(s))
		}
	}
	return pending
}
