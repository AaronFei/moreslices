package moreslices

// Remove the element from slice if true
func DeleteAll[S ~[]E, E any](s S, pFunc func(e E) bool) S {
	for i := 0; i < len(s); i++ {
		if pFunc(s[i]) {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

// Insert the element to specific index
// return nil if index out of range
func Insert[S ~[]E, E any](s S, e E, index int) S {
	if index > len(s) {
		return nil
	} else if index == len(s) {
		s = append(s, e)
	} else {
		s = append(s[:index+1], s[index:]...)
		s[index] = e
	}
	return s
}
