package util

// It does not cause any allocations.
func N(n int) []struct{} {
	return make([]struct{}, n)
}
