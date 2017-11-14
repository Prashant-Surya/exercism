package acronym

import (
	"strings"
)

//func Abbreviate(string) string

func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	var parts []string = strings.Split(s, " ")
	var out string = ""
	for i := 0; i < len(parts); i++ {
		sub_parts := strings.Split(parts[i], "-")
		for j := 0; j < len(sub_parts); j++ {
			out += string(sub_parts[j][0])
		}
	}

	return out
}
