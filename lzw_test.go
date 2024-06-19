package gocompress

import (
	"fmt"
	"testing"
)

func TestLZW_encoding(t *testing.T) {
	test_input := "Hello there"
	var test_expected_output []int
	test_expected_output = []int{72, 101, 108, 108, 111, 32, 116, 104, 101, 114, 101}

	test_output := LZW_encoding(test_input)

	for i, value := range test_output {
		if test_expected_output[i] != test_output[i] {
			t.Error("Failted rle compression on: " + fmt.Sprint(test_input[i]) +
				". Output was: " + fmt.Sprint(value) +
				". Expected output was: " + fmt.Sprint(test_expected_output[i]))
		}
	}
}
