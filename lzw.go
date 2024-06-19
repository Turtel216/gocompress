package gocompress

import (
	"fmt"
	"unicode/utf8"

	util "github.com/Turtel216/gocompress/internal/util"
)

// Function to encode string using LWZ encoding
func LZW_encoding(str string) []int {
	table := make(map[string]int)

	for i := 0; i <= 255; i++ {
		ch := ""
		ch += fmt.Sprint(i)
		table[ch] = i
	}

	p := ""
	c := ""
	p += string(str[0])
	code := 256
	var output_code []int

	for i := 0; i < utf8.RuneCountInString(str); i++ {
		if i != utf8.RuneCountInString(str)-1 {
			c += string(str[i+1])
		}
		if table[p+c] != util.GetLastValueOfMap(table) {
			p = p + c
		} else {
			output_code = append(output_code, table[p], len(output_code))
			table[p+c] = code
			code++
			p = c
		}

		c = ""
	}

	output_code = append(output_code, table[p], len(output_code))

	return output_code
}
