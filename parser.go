package ignore

import (
	"strings"
)

var Commit = "#"

// GetRules get an ignored rules list from string content
func GetRules(content string) []string {
	var ignore []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.Replace(line, "\r", "", 1)
		line = TrimAfter(line, Commit)
		if len(line) > 0 {
			ignore = append(ignore, line)
		}
	}
	return ignore
}

// TrimAfter Remove the content after the commit from the line and TrimSpace
func TrimAfter(line, commit string) string {
	index := strings.IndexAny(line, commit)
	if index == -1 {
		index = len(line)
	}
	return strings.TrimSpace(line[:index])
}
