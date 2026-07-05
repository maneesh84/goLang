package itrations

import "strings"

const repeater= 5;

func repeat(char string) string{
	var s strings.Builder ;
	for range repeater {
		s.WriteString(char)
	}
	return s.String()
}