package utils

func ReverseString(s string) string {
	s1 := ""
	for i := len(s) - 1; i >= 0; i-- {
		s1 += string(s[i])
	}
	return s1
}
