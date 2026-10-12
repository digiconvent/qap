package layer2

import (
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

func (i Intents) Comment(ctx fiber.Ctx, call Call) error {
	s, failure := i.find(ctx)
	if failure != nil {
		return failure.Send(ctx)
	}
	if !i.dependencies.Eligible(call.User, s.Policy) {
		return concepts.NotEligible.Send(ctx)
	}
	encoded := ctx.Get(commentHeader)
	if encoded == "" {
		return concepts.MissingComment.Send(ctx)
	}
	text, err := url.PathUnescape(encoded)
	if err != nil {
		return concepts.InvalidComment.Send(ctx)
	}
	if utf8.RuneCountInString(text) > maxCommentLength {
		return concepts.CommentTooLong.Send(ctx)
	}
	comment := concepts.Comment{Author: call.User, At: i.dependencies.Now().Unix(), Text: strings.Clone(text)}
	s, err = i.dependencies.Store.Comment(s.Id, comment)
	if err != nil {
		return concepts.UnknownSubmission.Send(ctx)
	}
	return ctx.JSON(i.dependencies.View(s, call.User))
}
