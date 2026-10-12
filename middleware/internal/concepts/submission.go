package concepts

type Submission struct {
	Id         string
	Endpoint   string
	Method     string
	Headers    map[string][]string
	Body       string
	Route      string
	Params     map[string]string
	Handler    Handler
	Permission string
	Policy     Policy
	Author     string
	When       int64
	Resolution string
	Result     *Result
	Favour     []string
	Against    []string
	Withhold   []string
	Comments   []Comment
	LastPolled int64
	Revision   int
	VoterCount int
}

func (s Submission) Request() Request {
	submission := s
	return Request{Method: s.Method, Url: s.Endpoint, Route: s.Route, Params: s.Params, Headers: s.Headers, Body: s.Body, User: s.Author, Submission: &submission}
}
