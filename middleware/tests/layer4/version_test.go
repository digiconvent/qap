package layer4_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// the version is the placeholder or a release version, and the client carries exactly the same one
func TestVersion(t *testing.T) {
	if layer4.Version != "-1.-1.-1" && !regexp.MustCompile(`^\d+\.\d+\.\d+$`).MatchString(layer4.Version) {
		t.Fatalf("got %q", layer4.Version)
	}
	client, err := os.ReadFile("../../../client/src/version.ts")
	if err != nil || string(client) != "export const version = '"+layer4.Version+"'\n" {
		t.Fatalf("the client has %q, the middleware %q (%v)", client, layer4.Version, err)
	}
}
