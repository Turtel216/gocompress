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

// Function to decompress a string that has been compressed using LZW encoding
func LZW_decompression(encoded_input []int) string {
	// dictionery mapping longest encountered words and a list of codes
	table := make(map[int]string)

	// Initialize table
	for i := 0; i <= 255; i++ {
		table[i] = string(rune(i))
	}

	var old int = encoded_input[0]
	var _new int
	var s string = table[old]
	c := ""
	c += string(s[0])
	var count int = 256
	var decoded_str string = ""
	decoded_str += s

	for i := 0; i < len(encoded_input)-1; i++ {
		_new = encoded_input[i+1]

		if _, ok := table[_new]; !ok {
			s = table[old] + string(c[0])
		} else {
			s = table[_new]
		}

		decoded_str += s
		c = ""
		c += string(s[0])
		table[count] = table[old] + c
		count++
		old = _new
	}

	return decoded_str
}
