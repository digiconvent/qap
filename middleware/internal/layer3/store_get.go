package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (m *Store) Get(id string) (concepts.Submission, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	s, ok := m.submissions[id]
	return clone(s), ok
}
