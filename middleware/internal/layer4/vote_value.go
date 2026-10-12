package layer4

import (
	"strconv"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

func ReadVote(header string) (string, bool) {
	value, err := strconv.Atoi(header)
	switch {
	case err != nil:
		return "", false
	case value > 0:
		return concepts.Favour, true
	case value < 0:
		return concepts.Against, true
	}
	return concepts.Withhold, true
}
