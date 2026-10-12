package layer4_test

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
	"github.com/gofiber/fiber/v3"
)

// copies keep the route template and the headers without the protocol headers, and stay intact after the framework reuses its buffers for later requests
func TestCopy_SurvivesLaterRequests(t *testing.T) {
	var copies []concepts.Request
	app := fiber.New()
	app.All("/*", func(ctx fiber.Ctx) error {
		copies = append(copies, layer4.CopyRequest(ctx))
		return nil
	})
	first := httptest.NewRequest(http.MethodDelete, "/api/library/book/123/delete?reason=damaged", strings.NewReader("first body"))
	first.Header.Set("Content-Type", "application/json")
	first.Header.Set("X-QAP-Intent", "submit")
	app.Test(first)
	second := httptest.NewRequest(http.MethodPost, "/x", strings.NewReader("second body that is longer"))
	second.Header.Set("Content-Type", "text/plain; charset=utf-8")
	app.Test(second)

	got := copies[0]
	if got.Method != "DELETE" || got.Url != "/api/library/book/123/delete?reason=damaged" || got.Route != "/*" || got.Body != "first body" {
		t.Fatalf("got %+v", got)
	}
	if !slices.Equal(got.Headers["Content-Type"], []string{"application/json"}) {
		t.Fatalf("got headers %v", got.Headers)
	}
	if _, kept := got.Headers["X-Qap-Intent"]; kept {
		t.Fatalf("the protocol header was kept: %v", got.Headers)
	}
}

// copies keep the route parameters by name
func TestCopy_KeepsTheRouteParameters(t *testing.T) {
	var copied concepts.Request
	app := fiber.New()
	app.Post("/api/library/:library/book/:id", func(ctx fiber.Ctx) error {
		copied = layer4.CopyRequest(ctx)
		return nil
	})
	app.Test(httptest.NewRequest(http.MethodPost, "/api/library/shire/book/123", nil))

	if copied.Params["library"] != "shire" || copied.Params["id"] != "123" || len(copied.Params) != 2 {
		t.Fatalf("got %v", copied.Params)
	}
}
