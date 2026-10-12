package layer3

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func CarryOut(s concepts.Submission) (concepts.Result, map[string]concepts.Change) {
	request := s.Request()
	if s.Handler.Check != nil {
		if problems := layer4.Problems(s.Handler.Check(request.Body)); problems != nil {
			return concepts.Result{Status: 422, Headers: map[string]string{}, Body: concepts.CannotBeCarriedOut(problems).Body}, nil
		}
	}
	id := request.Params["id"]
	if s.Handler.Read == nil || id == "" {
		return s.Handler.Execute(request), nil
	}
	before := s.Handler.Read(id)
	result := s.Handler.Execute(request)
	return result, layer4.Changes(before, s.Handler.Read(id))
}
