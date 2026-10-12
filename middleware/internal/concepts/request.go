package concepts

type Request struct {
	Method     string
	Url        string
	Route      string
	Params     map[string]string
	Headers    map[string][]string
	Body       string
	User       string
	Submission *Submission
}
