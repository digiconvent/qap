package layer3

import (
	"sync"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

type Store struct {
	mutex       sync.Mutex
	order       []string
	submissions map[string]concepts.Submission
}

func NewStore() *Store {
	return &Store{submissions: map[string]concepts.Submission{}}
}
