package layer4

func DeadlinePassed(created int64, seconds *int, now int64) bool {
	return seconds != nil && now >= created+int64(*seconds)
}

func ExtensionPassed(created int64, deadline *int, extension *int, now int64) bool {
	return deadline != nil && extension != nil && now >= created+int64(*deadline)+int64(*extension)
}
