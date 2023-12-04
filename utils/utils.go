package utils

import "strings"

type Line struct {
	Index int
	Line  string
}

func ParseFormat(input []string) []*Line {
	ret := make([]*Line, len(input))
	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], ":")
		index, _, _ := ExtractNum(len(parts[0])-1, parts[0])
		ret[i] = &Line{
			Index: index,
			Line:  strings.TrimSpace(parts[1]),
		}
	}
	return ret
}
