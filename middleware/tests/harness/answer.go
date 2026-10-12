package harness

type Answer struct {
	Status int
	Raw    string
	Body   map[string]any
	List   []any
}

func (a Answer) Submission() map[string]any {
	submission, _ := a.Body["submission"].(map[string]any)
	return submission
}

func (a Answer) Poll() map[string]any {
	poll, _ := a.Body["poll"].(map[string]any)
	return poll
}

func (a Answer) Id() string {
	id, _ := a.Submission()["id"].(string)
	return id
}

func (a Answer) Resolution() string {
	resolution, _ := a.Submission()["resolution"].(string)
	return resolution
}

func (a Answer) Votes(kind string) int {
	count, _ := a.Poll()[kind].(float64)
	return int(count)
}

func (a Answer) Message() string {
	message, _ := a.Body["error"].(string)
	return message
}

func (a Answer) Result() map[string]any {
	result, _ := a.Submission()["result"].(map[string]any)
	return result
}

func (a Answer) PolicyNames() []string {
	var names []string
	policies, _ := a.Body["policies"].([]any)
	for _, policy := range policies {
		p, _ := policy.(map[string]any)
		name, _ := p["name"].(string)
		names = append(names, name)
	}
	return names
}
