package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// withhold votes are allowed exactly when the policy permits them
func TestWithholdAllowed(t *testing.T) {
	if !layer4.WithholdAllowed(concepts.Policy{WithholdPermitted: true}) || layer4.WithholdAllowed(concepts.Policy{}) {
		t.Fatal("withhold rule does not follow the policy")
	}
}
