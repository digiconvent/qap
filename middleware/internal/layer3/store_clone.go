package layer3

import (
	"maps"
	"slices"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

func clone(s concepts.Submission) concepts.Submission {
	s.Policy = clonePolicy(s.Policy)
	s.Headers = cloneHeaders(s.Headers)
	s.Params = maps.Clone(s.Params)
	s.Favour = slices.Clone(s.Favour)
	s.Against = slices.Clone(s.Against)
	s.Withhold = slices.Clone(s.Withhold)
	s.Comments = slices.Clone(s.Comments)
	if s.Result != nil {
		result := *s.Result
		result.Headers = maps.Clone(result.Headers)
		s.Result = &result
	}
	return s
}

func cloneHeaders(headers map[string][]string) map[string][]string {
	if headers == nil {
		return nil
	}
	copied := make(map[string][]string, len(headers))
	for name, values := range headers {
		copied[name] = slices.Clone(values)
	}
	return copied
}

func clonePolicy(p concepts.Policy) concepts.Policy {
	p.Groups = slices.Clone(p.Groups)
	p.AgainstThreshold = cloneNumber(p.AgainstThreshold)
	p.Quorum = cloneNumber(p.Quorum)
	p.DeadlineSeconds = cloneNumber(p.DeadlineSeconds)
	p.ExtensionSeconds = cloneNumber(p.ExtensionSeconds)
	return p
}

func cloneNumber(n *int) *int {
	if n == nil {
		return nil
	}
	value := *n
	return &value
}
