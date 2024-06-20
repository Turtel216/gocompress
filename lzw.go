package gocompress

// Function to encode string using LWZ encoding
func LZW_encoding(str string) []int {
	// Map to store code-to-string mappings
	table := make(map[string]int)

	// Initialize table
	for i := 0; i <= 255; i++ {
		table[string(rune(i))] = i
	}

	prev := "" //prev input character
	curr := "" // current input character
	prev += string(str[0])
	code := 256           // keeping track of current code
	var output_code []int // The slice holding the encoded output

	// Iterate over string
	for i := 0; i < len(str); i++ {
		// get next input character
		if i != len(str)-1 {
			curr += string(str[i+1])
		}
		// if the last two characters are in the table,
		// add them to the prev
		if _, ok := table[prev+curr]; ok {
			prev = prev + curr
		} else {
			// append the corresponding code the output
			output_code = append(output_code, table[prev])
			// add the code to the map
			table[prev+curr] = code

			// Update to next avaible code
			code++
			// Update previous character for next iteration
			prev = curr
		}

		// Reset current
		curr = ""
	}

	// append the missing last entry
	output_code = append(output_code, table[prev])

	return output_code
}

// Function to decompress a string that has been compressed using LZW encoding
func LZW_decompression(encoded_input []int) string {
	// Map to store code-to-string mappings
	table := make(map[int]string)

	// Initialize table
	for i := 0; i <= 255; i++ {
		table[i] = string(rune(i))
	}

	var old int = encoded_input[0]  // code for the previous string
	var _new int                    // code for the current element
	var tmp_str string = table[old] // Decoded string base on the previous string
	curr := ""                      // String to be built for the current code
	curr += string(tmp_str[0])
	var count int = 256         // Next availbe code for new strings
	var decoded_str string = "" // The decoded string that will be returned
	decoded_str += tmp_str      // add the first decoded character

	// Iterate through the encoded sequence until the second to last element
	for i := 0; i < len(encoded_input)-1; i++ {
		// update current code
		_new = encoded_input[i+1]

		// handle new code not in the table yet
		if _, ok := table[_new]; !ok {
			// combine previous string with first character of current node
			tmp_str = table[old] + string(curr[0])
		} else {
			// Existing code found in the table, use its corresponding  string
			tmp_str = table[_new]
		}

		// Update output string
		decoded_str += tmp_str

		// Reset the string to be built
		curr = ""
		curr += string(tmp_str[0])

		// Add a new entry to the table with next code and combined string
		table[count] = table[old] + curr

		// Update to next avaible code
		count++
		// Update the previous code to the current code
		old = _new
	}

	return decoded_str
}
