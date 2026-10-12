package layer4

import (
	"crypto/rand"
	"encoding/hex"
)

func NewSubmissionId() string {
	random := make([]byte, 16)
	rand.Read(random)
	return "qap_" + hex.EncodeToString(random)
}
