package concepts

type Handler struct {
	Execute func(request Request) Result
	Check   func(body string) map[string][]error
	Read    func(id string) any
}
