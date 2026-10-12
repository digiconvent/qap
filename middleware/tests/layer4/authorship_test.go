package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// only the person who created a submission is its author
func TestIsAuthor(t *testing.T) {
	s := concepts.Submission{Author: "alice"}
	if !layer4.IsAuthor(s, "alice") || layer4.IsAuthor(s, "bob") {
		t.Fatal("authorship is wrong")
	}
}
