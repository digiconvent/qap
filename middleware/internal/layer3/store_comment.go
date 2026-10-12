package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (m *Store) Comment(id string, comment concepts.Comment) (concepts.Submission, error) {
	return m.change(id, false, func(s concepts.Submission) concepts.Submission {
		s.Comments = append(s.Comments, comment)
		return s
	})
}
