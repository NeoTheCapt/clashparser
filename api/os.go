package api

func mac(rules []string) []string {
	return append(
		Parser.Rules4Mac,
		rules...,
	)
}

func win(rules []string) []string {
	return append(
		Parser.Rules4Win,
		rules...,
	)
}

func ios(rules []string) []string {
	return append(
		Parser.Rules4Ios,
		rules...,
	)
}
