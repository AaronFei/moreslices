package moreslices

// Remove the element from slice if true
func DeleteAll[E any](s []E, pFunc func(e E) bool) []E {
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
func Insert[E any](s []E, e E, index int) []E {
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
