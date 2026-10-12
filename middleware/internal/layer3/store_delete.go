package layer3

import (
	"slices"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

func (m *Store) Delete(id string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	s, ok := m.submissions[id]
	if !ok {
		return concepts.ErrUnknownSubmission
	}
	if s.Resolution != concepts.Pending {
		return concepts.ErrAlreadyResolved
	}
	m.remove(id)
	return nil
}

func (m *Store) remove(id string) {
	delete(m.submissions, id)
	m.order = slices.DeleteFunc(m.order, func(other string) bool { return other == id })
}
