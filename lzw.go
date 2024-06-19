package gocompress

// Function to encode string using LWZ encoding
func LZW_encoding(str string) []int {
	// dictionery mapping longest encountered words and a list of codes
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

			code++
			prev = curr
		}

		curr = ""
	}

	// append the missing last entry
	output_code = append(output_code, table[prev])

	return output_code
}
