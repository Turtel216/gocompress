package gocompress

// Function to encode string using LWZ encoding
func LZW_encoding(str string) []int {
	table := make(map[string]int)

	for i := 0; i <= 255; i++ {
		table[string(rune(i))] = i
	}

	prev := ""
	curr := ""
	prev += string(str[0])
	code := 256
	var output_code []int

	for i := 0; i < len(str); i++ {
		if i != len(str)-1 {
			curr += string(str[i+1])
		}
		if _, ok := table[prev+curr]; ok {
			prev = prev + curr
		} else {
			output_code = append(output_code, table[prev])
			table[prev+curr] = code
			code++
			prev = curr
		}

		curr = ""
	}

	output_code = append(output_code, table[prev])

	return output_code
}
