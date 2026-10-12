package layer4

import "regexp"

var placeholder = regexp.MustCompile(`\{(\w+)\}`)

func ResolvePermission(template string, param func(key string) string) (string, bool) {
	ok := true
	resolved := placeholder.ReplaceAllStringFunc(template, func(match string) string {
		value := param(placeholder.FindStringSubmatch(match)[1])
		if value == "" {
			ok = false
		}
		return value
	})
	return resolved, ok
}
