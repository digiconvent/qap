package realm

import (
	"slices"

	qap "github.com/digiconvent/qap/middleware"
)

type Recording struct {
	Executed []execution        `json:"executed"`
	Audited  []qap.AuditSummary `json:"audited"`
}

func (r *Realm) Record(summary qap.AuditSummary) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.audited = append(r.audited, summary)
}

func (r *Realm) Recorded() Recording {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return Recording{Executed: slices.Clone(r.executed), Audited: slices.Clone(r.audited)}
}
