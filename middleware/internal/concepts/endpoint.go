package concepts

type Handler struct {
	Execute func(request Request) Result
	Check   func(request Request) map[string][]error
	Read    func(id string) any
}
