package util

func EqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func MapStringSlice(f func(string) string, a []string) []string {
	for i := range a {
		a[i] = f(a[i])
	}
	return a
}
