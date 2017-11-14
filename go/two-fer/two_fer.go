package twofer

func ShareWith(s string) string {
	var out string = "you"
	if s != "" {
		out = s
	}
	return "One for " + out + ", one for me."
}
