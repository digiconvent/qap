package concepts

type TypedEndpoint[Payload any] interface {
	Check(payload Payload) map[string][]error
	Execute(request TypedRequest[Payload]) Result
}

type Reader interface {
	Read(id string) any
}

type TypedRequest[Payload any] struct {
	Payload    Payload
	Method     string
	Url        string
	Route      string
	Params     map[string]string
	Headers    map[string][]string
	User       string
	Submission *Submission
}

type Handler struct {
	Check   func(body string) map[string][]error
	Execute func(request Request) Result
	Read    func(id string) any
}
