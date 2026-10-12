package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func params(values map[string]string) func(key string) string {
	return func(key string) string { return values[key] }
}

// every placeholder is filled with the route parameter of the same name
func TestResolvePermission(t *testing.T) {
	cases := []struct {
		template string
		values   map[string]string
		want     string
		ok       bool
	}{
		{"library.book.delete", nil, "library.book.delete", true},
		{"library.{id}.add-book", map[string]string{"id": "123"}, "library.123.add-book", true},
		{"{org}.book.{id}", map[string]string{"org": "acme", "id": "7"}, "acme.book.7", true},
		{"library.{id}.add-book", nil, "library..add-book", false},
	}
	for _, c := range cases {
		if got, ok := layer4.ResolvePermission(c.template, params(c.values)); got != c.want || ok != c.ok {
			t.Errorf("%s: got %q %v, want %q %v", c.template, got, ok, c.want, c.ok)
		}
	}
}
