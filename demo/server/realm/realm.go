package realm

import (
	"sync"

	qap "github.com/digiconvent/qap/middleware"
)

type Realm struct {
	mutex       sync.Mutex
	Guard       qap.Guard
	executed    []execution
	audited     []qap.AuditSummary
	memberships map[string][]string
}

type execution struct {
	Method   string  `json:"method"`
	Endpoint string  `json:"endpoint"`
	Payload  payload `json:"payload"`
	Author   string  `json:"author"`
}
