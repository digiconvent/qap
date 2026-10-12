package layer4

import "github.com/digiconvent/qap/middleware/internal/concepts"

type Counts struct {
	Favour        int
	Against       int
	Withhold      int
	Participation int
}

func CountVotes(s concepts.Submission) Counts {
	return Counts{
		Favour:        len(s.Favour),
		Against:       len(s.Against),
		Withhold:      len(s.Withhold),
		Participation: len(s.Favour) + len(s.Against),
	}
}
