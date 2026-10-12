package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (m *Store) Create(s concepts.Submission) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.order = append(m.order, s.Id)
	m.submissions[s.Id] = clone(s)
}
