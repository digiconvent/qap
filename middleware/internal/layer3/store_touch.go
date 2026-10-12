package layer3

func (m *Store) Touch(id string, at int64) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	s, ok := m.submissions[id]
	if !ok {
		return
	}
	s.LastPolled = at
	m.submissions[id] = s
}
