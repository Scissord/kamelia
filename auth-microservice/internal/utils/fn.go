package utils

func StringOrEmpty(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
