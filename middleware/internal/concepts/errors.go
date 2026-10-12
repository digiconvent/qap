package concepts

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

var (
	ErrUnknownSubmission = errors.New("unknown submission")
	ErrAlreadyResolved   = errors.New("submission already resolved")
	ErrOutdated          = errors.New("submission changed since it was evaluated")
)

type Error struct {
	Status int
	Body   map[string]any
}

func (e Error) Send(ctx fiber.Ctx) error {
	return ctx.Status(e.Status).JSON(e.Body)
}

func failure(status int, message string) Error {
	return Error{Status: status, Body: map[string]any{"error": message}}
}

var (
	Unauthorized            = failure(401, "unauthorized")
	MissingIntent           = failure(406, "QAP intent header required")
	PolicyDoesNotIncludeYou = failure(403, "Policy does not include you")
	MisconfiguredRoute      = failure(500, "misconfigured permission template")
	UnknownSubmission       = failure(404, "Submission state does not exist")
	AlreadyResolved         = failure(409, "Submission is already resolved")
	NotEligible             = failure(403, "You cannot vote on this")
	InvalidVote             = failure(400, "X-QAP-Vote must be a whole number")
	WithholdNotPermitted    = failure(400, "withhold is not permitted for this poll")
	NotTheAuthor            = failure(403, "Only the author can delete this submission")
	InvalidComment          = failure(400, "X-QAP-Comment must be percent-encoded")
	CommentTooLong          = failure(400, "X-QAP-Comment must be at most 500 characters")
	InvalidId               = failure(400, "the id must be a UUID")
	MissingSubmissionId     = Error{Status: 406, Body: map[string]any{"error": "QAP id header required", "hint": "Specify X-QAP-ID"}}
	MissingComment          = Error{Status: 406, Body: map[string]any{"error": "QAP comment header required", "hint": "Specify X-QAP-Comment"}}
)

func UnknownIntent(value string) Error {
	return failure(400, "Unknown intent: '"+value+"'")
}

func CannotBeCarriedOut(problems map[string][]string) Error {
	return Error{Status: 422, Body: map[string]any{"error": "the request cannot be carried out", "problems": problems}}
}

func MissingPolicyId(acceptable []PolicyReference) Error {
	return Error{Status: 406, Body: map[string]any{
		"error":               "QAP policy id header required",
		"hint":                "Specify X-QAP-Policy-ID",
		"acceptable_policies": acceptable,
	}}
}
