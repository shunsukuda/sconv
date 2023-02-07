package sconv

// Str19DecToU64 convert string(base 10, last 19 characters) to uint64.
// If parse error or length > 19, return false.
func Str19DecToU64(s string) (uint64, bool) {
	var v uint64
	l := len(s)
	if l <= 0 {
		return 0, false
	}
	for i := 0; i < 19; i++ {
		if l+i >= 19 {
			d := s[l-19+i]
			if d < '0' || d > '9' {
				return 0, false
			}
			v = v*10 + uint64(d-'0')
		} else {
			i = 18 - l
		}
	}
	return v, l <= 19
}
