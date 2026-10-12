package layer4_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
	"github.com/gofiber/fiber/v3"
)

func written(t *testing.T, result concepts.Result) (*http.Response, string) {
	t.Helper()
	app := fiber.New()
	app.Get("/", func(ctx fiber.Ctx) error { return layer4.WriteResult(ctx, result) })
	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	return resp, string(body)
}

// a result is written with its status and headers, and its body as text, bytes, json or nothing
func TestWriteResult(t *testing.T) {
	cases := []struct {
		body any
		want string
	}{
		{nil, ""},
		{"plain text", "plain text"},
		{[]byte("raw bytes"), "raw bytes"},
		{map[string]any{"title": "There and Back Again"}, `{"title":"There and Back Again"}`},
	}
	for _, c := range cases {
		resp, body := written(t, concepts.Result{Status: 201, Headers: map[string]string{"X-Book": "red"}, Body: c.body})
		if resp.StatusCode != 201 || resp.Header.Get("X-Book") != "red" || body != c.want {
			t.Errorf("%v: got %d %q %q", c.body, resp.StatusCode, resp.Header.Get("X-Book"), body)
		}
	}
}
