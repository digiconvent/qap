package concepts

type AuditSummary struct {
	Anonymous    bool              `json:"anonymous"`
	SubmissionId string            `json:"submission_id"`
	Permission   string            `json:"permission"`
	Policy       PolicyReference   `json:"policy"`
	Author       string            `json:"author"`
	Created      int64             `json:"created"`
	Decided      int64             `json:"decided"`
	Url          string            `json:"url"`
	Body         string            `json:"body"`
	Resolution   string            `json:"resolution"`
	Result       *Result           `json:"result"`
	Favour       []string          `json:"favour"`
	Against      []string          `json:"against"`
	Withhold     []string          `json:"withhold"`
	Comments     []Comment         `json:"comments,omitempty"`
	Changes      map[string]Change `json:"changes,omitempty"`
}
