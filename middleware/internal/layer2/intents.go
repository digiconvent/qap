package layer2

const (
	submissionIdHeader = "X-QAP-ID"
	policyIdHeader     = "X-QAP-Policy-ID"
	voteHeader         = "X-QAP-Vote"
	commentHeader      = "X-QAP-Comment"
	versionHeader      = "X-QAP-Version"
	maxCommentLength   = 500
)

type Intents struct {
	dependencies Dependencies
}

func New(dependencies Dependencies) Intents {
	return Intents{dependencies: dependencies}
}

func (i Intents) Handlers() map[string]Handler {
	return map[string]Handler{
		"discover":          i.Discover,
		"submit":            i.Submit,
		"list":              i.List,
		"poll":              i.Poll,
		"vote":              i.Vote,
		"comment":           i.Comment,
		"delete-submission": i.Withdraw,
	}
}
