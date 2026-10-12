package layer4_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// the problems of a check become messages per field, and a check without errors has no problems
func TestProblems(t *testing.T) {
	got := layer4.Problems(map[string][]error{"title": {errors.New("is required"), nil}, "isbn": {}, "year": nil})
	if want := map[string][]string{"title": {"is required"}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	if got := layer4.Problems(map[string][]error{"title": {nil}}); got != nil {
		t.Fatalf("expected no problems, got %v", got)
	}
	if got := layer4.Problems(nil); got != nil {
		t.Fatalf("expected no problems, got %v", got)
	}
}
