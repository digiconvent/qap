package layer4

func Problems(errs map[string][]error) map[string][]string {
	var problems map[string][]string
	for field, list := range errs {
		for _, err := range list {
			if err == nil {
				continue
			}
			if problems == nil {
				problems = map[string][]string{}
			}
			problems[field] = append(problems[field], err.Error())
		}
	}
	return problems
}
