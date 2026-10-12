package layer4

import (
	"encoding/json"
	"strings"
)

func DecodePayload[Payload any](body string) (Payload, map[string][]error) {
	var payload Payload
	if strings.TrimSpace(body) == "" {
		return payload, nil
	}
	if err := json.Unmarshal([]byte(body), &payload); err != nil {
		return payload, map[string][]error{"body": {err}}
	}
	return payload, nil
}
