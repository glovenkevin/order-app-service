package variable

func GetString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func NewString(s string) (r *string) {
	r = &s
	return
}
