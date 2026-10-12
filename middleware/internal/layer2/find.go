package layer2

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

func (i Intents) find(ctx fiber.Ctx) (concepts.Submission, *concepts.Error) {
	id := ctx.Get(submissionIdHeader)
	if id == "" {
		return concepts.Submission{}, &concepts.MissingSubmissionId
	}
	s, ok := i.dependencies.Store.Get(id)
	if !ok {
		return concepts.Submission{}, &concepts.UnknownSubmission
	}
	return s, nil
}
