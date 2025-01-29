package CodeWars

import (
	"strings"
)

func Compose(s1, s2 string) string {
	lines1 := strings.Split(s1, "\n")
	lines2 := strings.Split(s2, "\n")
	n := len(lines1)

	var resultLines []string

	for i := 0; i < n; i++ {
		line1 := lines1[i]
		line2 := lines2[n-1-i]

		prefix := line1[:i+1]
		suffix := line2[:len(line2)-i]

		resultLines = append(resultLines, prefix+suffix)
	}

	return strings.Join(resultLines, "\n")
}
