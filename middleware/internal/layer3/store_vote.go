package layer3

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func (m *Store) Vote(id string, user string, vote string) (concepts.Submission, error) {
	return m.change(id, true, func(s concepts.Submission) concepts.Submission {
		s = layer4.RecordVote(s, user, vote)
		s.Revision++
		return s
	})
}
