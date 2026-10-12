package layer4

import "github.com/digiconvent/qap/middleware/internal/concepts"

const idleSeconds = 60

func Expired(s concepts.Submission, now int64) bool {
	return s.Resolution != concepts.Pending && now-s.LastPolled >= idleSeconds
}
