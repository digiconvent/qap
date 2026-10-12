package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (m *Store) SetResult(id string, result concepts.Result) (concepts.Submission, error) {
	return m.change(id, false, func(s concepts.Submission) concepts.Submission {
		s.Result = &result
		return s
	})
}
