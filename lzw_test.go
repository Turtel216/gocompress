package gocompress

import (
	"fmt"
	"testing"
)

func TestLZW_encoding(t *testing.T) {
	test_input := "geekific-geekific"
	// test_expected_output := "3abc2da3f"
	test_output := LZW_encoding(test_input)

	fmt.Println(test_output)

	// if test_output != test_expected_output {
	// 	t.Error("Failted rle compression on string: " + test_input + ". Output was: " + test_output)
	// }
}
