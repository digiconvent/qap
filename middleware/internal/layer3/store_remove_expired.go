package layer3

import (
	"slices"

	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func (m *Store) RemoveExpired(now int64) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.order = slices.DeleteFunc(m.order, func(id string) bool {
		if !layer4.Expired(m.submissions[id], now) {
			return false
		}
		delete(m.submissions, id)
		return true
	})
}
