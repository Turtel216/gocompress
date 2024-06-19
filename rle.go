package gocompress

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Lossless compression using the Run-length encoding algorithm.
Best suited for compressing strings with a lot of repeating
characters
*/
func RLE_compress(str string) string { //TODO fixe issue when duplicates are on the end of the list

	var prev_char rune           // Stores the previous character
	var first_duplicate int      // tracks the index of the first duplicate
	var duplicate_number int = 0 // counts the number of duplicates
	var offset int = 0           // adjusts index to changes made

	// Iterate over the string
	for i, _char := range str {
		if prev_char != 0 { // Skip the first iteration

			// Check if a duplicate is encountered
			if prev_char == _char {

				// Check if this is the first duplicate
				if duplicate_number == 0 {
					first_duplicate = i - 1
					duplicate_number++

					// set previous character
					prev_char = _char
					continue
				}

				// set previous character
				prev_char = _char
				duplicate_number++
				continue
			}

			// Compress duplicates
			if prev_char != _char && duplicate_number != 0 {
				// Added previous duplicate
				duplicate_number++

				// Replace first duplicate and remove the rest of the duplicates
				str = str[:first_duplicate+offset] + fmt.Sprint(duplicate_number) + str[i-1+offset:]

				// Calculate offset
				offset = offset - duplicate_number + 2
				// Reset duplicate counter
				duplicate_number = 0

				// set previous character
				prev_char = _char
				continue
			}
		}

		prev_char = _char
	}

	if duplicate_number != 0 {
		// Added previous duplicate
		duplicate_number++

		// Replace first duplicate and remove the rest of the duplicates
		str = str[:first_duplicate+offset] + fmt.Sprint(duplicate_number) + string(prev_char)
	}

	// Return compressed string
	return str
}

/*
Decompress a string that has been compressed
with the Run-length encoding algorithm
*/
func RLE_decompress(str string) string {
	var offset int = 0 // adjusts index to changes made
	var prev_char rune

	// Iterate over the string
	for i, _char := range str {
		if prev_char != 0 { // skip first iteration
			// if the char is a number,
			// add this number duplicates of the following char
			if unicode.IsNumber(prev_char) {
				// convert char to int
				digit, _ := strconv.Atoi(string(prev_char))

				// add the duplcicates to the string
				str = str[:i-1+offset] + getDuplicatesString(digit, _char) + str[i+1+offset:]

				// Calculate offset
				offset = offset + digit - 2
				prev_char = _char
				continue
			}

			prev_char = _char
			continue
		}

		prev_char = _char
	}

	return str
}

// Helper function that returns n number of repeats of the given rune
func getDuplicatesString(n int, _char rune) string {
	var new_str string        // the string to be returned
	str_char := string(_char) // the char converted to string

	// Create string with duplicates
	new_str = strings.Repeat(str_char, n)

	return new_str
}
