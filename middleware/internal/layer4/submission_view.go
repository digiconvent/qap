package layer4

import (
	"net/http"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

type submissionPart struct {
	Id         string                   `json:"id"`
	Endpoint   string                   `json:"endpoint"`
	Method     string                   `json:"method"`
	Body       string                   `json:"body"`
	Permission string                   `json:"permission"`
	Policy     concepts.PolicyReference `json:"policy"`
	Author     string                   `json:"author"`
	When       int64                    `json:"when"`
	Resolution string                   `json:"resolution"`
	Result     *concepts.Result         `json:"result"`
}

type pollPart struct {
	Favour     int                `json:"favour"`
	Against    int                `json:"against"`
	Withhold   int                `json:"withhold"`
	Comments   []concepts.Comment `json:"comments"`
	Resolution string             `json:"resolution"`
}

type view struct {
	Submission submissionPart `json:"submission"`
	Poll       pollPart       `json:"poll"`
}

func SubmissionView(s concepts.Submission, user string) any {
	result := s.Result
	if result != nil && s.Method == http.MethodGet && user != s.Author {
		hidden := *result
		hidden.Body = nil
		result = &hidden
	}
	return view{
		Submission: submissionPart{
			Id: s.Id, Endpoint: s.Endpoint, Method: s.Method, Body: s.Body,
			Permission: s.Permission, Policy: concepts.PolicyReference{Id: s.Policy.Id, Name: s.Policy.Name},
			Author: s.Author, When: s.When, Resolution: s.Resolution, Result: result,
		},
		Poll: pollPart{
			Favour: len(s.Favour), Against: len(s.Against), Withhold: len(s.Withhold),
			Comments: s.Comments, Resolution: s.Resolution,
		},
	}
}
