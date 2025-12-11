package utils

import "strings"

func NormalizeSort(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	s = strings.Join(strings.Fields(s), " ") // remove duplicate spaces
	return s
}
