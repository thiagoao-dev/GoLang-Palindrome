package algorithm

// Palindrome function
func Palindrome(text []string) bool {
	input := text[0]
	var output string
	for k := range input {
		output = input[k:k+1] + output
	}
	return input == output
}
