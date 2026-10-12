package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func runExecutor(executor func(s concepts.Submission) (concepts.Result, map[string]concepts.Change), s concepts.Submission) (result concepts.Result, changes map[string]concepts.Change) {
	defer func() {
		if recover() != nil {
			result, changes = concepts.Result{Status: 500, Headers: map[string]string{}}, nil
		}
	}()
	return executor(s)
}
