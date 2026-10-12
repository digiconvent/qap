package layer4

import "github.com/digiconvent/qap/middleware/internal/concepts"

func IsAuthor(s concepts.Submission, user string) bool {
	return s.Author == user
}
