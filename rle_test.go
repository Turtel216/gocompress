package gocompress

import "testing"

func TestRlE_compression(t *testing.T) {
	test_input := "aaabcddaf"
	test_expected_output := "3abc2daf"
	test_output := RLE_compress(test_input)
	if test_output != test_expected_output {
		t.Error("Failted rle compression on string: " + test_input + ". Output was: " + test_output)
	}
}
