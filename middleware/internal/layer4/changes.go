package layer4

import (
	"encoding/json"
	"reflect"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

func Changes(before any, after any) map[string]concepts.Change {
	old, current := fieldsOf(before), fieldsOf(after)
	changes := map[string]concepts.Change{}
	for name, value := range old {
		if !reflect.DeepEqual(value, current[name]) {
			changes[name] = concepts.Change{Before: value, After: current[name]}
		}
	}
	for name, value := range current {
		if _, known := old[name]; !known {
			changes[name] = concepts.Change{Before: nil, After: value}
		}
	}
	if len(changes) == 0 {
		return nil
	}
	return changes
}

func fieldsOf(entity any) map[string]any {
	fields := map[string]any{}
	raw, err := json.Marshal(entity)
	if err != nil {
		return fields
	}
	if json.Unmarshal(raw, &fields) != nil {
		var value any
		json.Unmarshal(raw, &value)
		return map[string]any{"": value}
	}
	return fields
}
