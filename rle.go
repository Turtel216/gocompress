package gocompress

import "fmt"

/*
Lossless compression using the Run-length encoding algorithm.
Best suited for compressing strings with a lot of repeating
characters
*/
func RLE_compress(str string) string {

	var prev_char rune // Store the previous character
	var first_duplicate int
	var duplicate_number int = 0

	// Iterate over the string
	for i, _char := range str {
		if prev_char != 0 { // Skip the first iteration

			// Check if a duplicate is encountered
			if prev_char == _char {

				// Check if this is the first duplicate
				if duplicate_number == 0 {
					first_duplicate = i
					duplicate_number++
					continue
				}

				duplicate_number++
			}

			// Compress duplicates
			if prev_char != _char && duplicate_number != 0 {
				// Replace first duplicate
				str = str[:first_duplicate] + string('0'+duplicate_number) + str[:first_duplicate+1]

				// Remove all duplicates
				str = str[:first_duplicate+2] + str[:i-1]

				// Reset duplicate counter
				duplicate_number = 0
			}
		}

		prev_char = _char
	}

	// Return compressed string
	return str
}

/*
Decompress a string that has been compressed
with the Run-length encoding algorithm
*/
func RLE_decompress(str string) string {
	// Iterate over the string
	for i, _char := range str {
		fmt.Println(_char, i)
	}

	return str
}
