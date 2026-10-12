package layer2

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (i Intents) settle(s concepts.Submission) concepts.Submission {
	switch i.dependencies.Evaluate(s, s.VoterCount, i.dependencies.Now().Unix()) {
	case concepts.Executed:
		s = i.dependencies.Execute(s)
	case concepts.Rejected:
		s = i.dependencies.Reject(s)
	}
	if s.Resolution != concepts.Pending {
		i.dependencies.Deadlines.Cancel(s.Id)
	}
	return s
}
