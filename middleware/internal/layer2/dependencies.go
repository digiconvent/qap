package layer2

import (
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

type Store interface {
	Create(s concepts.Submission)
	Get(id string) (concepts.Submission, bool)
	Vote(id string, user string, vote string) (concepts.Submission, error)
	Comment(id string, comment concepts.Comment) (concepts.Submission, error)
	Delete(id string) error
	ListPending(route string, method string) []concepts.Submission
	Touch(id string, at int64)
}

type Deadlines interface {
	Start(s concepts.Submission)
	Cancel(id string)
}

type Dependencies struct {
	PolicyLookup     func(permission string) []concepts.Policy
	Eligible         func(user string, policy concepts.Policy) bool
	VoterCount       func(policy concepts.Policy) int
	Store            Store
	Deadlines        Deadlines
	Now              func() time.Time
	NewId            func() string
	CopyRequest      func(ctx fiber.Ctx) concepts.Request
	Problems         func(errs map[string][]error) map[string][]string
	View             func(s concepts.Submission, user string) any
	ReadVote         func(header string) (string, bool)
	WithholdAllowed  func(policy concepts.Policy) bool
	IsAuthor         func(s concepts.Submission, user string) bool
	Evaluate         func(s concepts.Submission, voterCount int, now int64) string
	Execute          func(s concepts.Submission) concepts.Submission
	Reject           func(s concepts.Submission) concepts.Submission
	RecordWithdrawal func(s concepts.Submission)
	Version          string
}
