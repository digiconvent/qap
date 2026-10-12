package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// a deadline passes from the exact moment of creation plus its seconds, and never without a deadline
func TestDeadlinePassed(t *testing.T) {
	sixty := 60
	if layer4.DeadlinePassed(1000, &sixty, 1059) || !layer4.DeadlinePassed(1000, &sixty, 1060) || !layer4.DeadlinePassed(1000, &sixty, 9999) || layer4.DeadlinePassed(1000, nil, 9999) {
		t.Fatal("deadline passes at the wrong moment")
	}
}

// an extension ends its seconds after the deadline, and never without both a deadline and an extension
func TestExtensionPassed(t *testing.T) {
	sixty, thirty := 60, 30
	if layer4.ExtensionPassed(1000, &sixty, &thirty, 1089) || !layer4.ExtensionPassed(1000, &sixty, &thirty, 1090) ||
		layer4.ExtensionPassed(1000, nil, &thirty, 9999) || layer4.ExtensionPassed(1000, &sixty, nil, 9999) {
		t.Fatal("extension passes at the wrong moment")
	}
}
