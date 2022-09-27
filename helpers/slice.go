package helpers

// Returns the maximum length of the strings from the slice.
// For each string in the slice, check if the length of the current string is greater than the current maximum length,
// then set current maximum length to the length of the current string.
func MaxLenOnSlice(slice *[]string) int {
	maxLen := 0
	for _, text := range *slice {
		if currentLen := len(text); currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

// Returns true if the text is in the slice
// Returns false otherwise.
func StringInSlice(text string, slice []string) bool {
	for _, b := range slice {
		if b == text {
			return true
		}
	}
	return false
}
