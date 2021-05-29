package helpers

func UnPtrStr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func PtrString(s string) *string {
	return &s
}
