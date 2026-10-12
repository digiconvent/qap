package layer4

import (
	"slices"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

func RecordVote(s concepts.Submission, user string, vote string) concepts.Submission {
	voter := func(other string) bool { return other == user }
	s.Favour = slices.DeleteFunc(slices.Clone(s.Favour), voter)
	s.Against = slices.DeleteFunc(slices.Clone(s.Against), voter)
	s.Withhold = slices.DeleteFunc(slices.Clone(s.Withhold), voter)
	switch vote {
	case concepts.Favour:
		s.Favour = append(s.Favour, user)
	case concepts.Against:
		s.Against = append(s.Against, user)
	default:
		s.Withhold = append(s.Withhold, user)
	}
	return s
}
