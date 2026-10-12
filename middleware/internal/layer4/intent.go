package layer4

type IntentOutcome int

const (
	KnownIntent IntentOutcome = iota
	MissingIntent
	UnknownIntent
)

func Intent[Handler any](header string, intents map[string]Handler) (Handler, IntentOutcome) {
	handler, known := intents[header]
	switch {
	case header == "":
		return handler, MissingIntent
	case !known:
		return handler, UnknownIntent
	}
	return handler, KnownIntent
}
