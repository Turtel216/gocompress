package gocompress

import "testing"

func TestRlE_compression(t *testing.T) {
	test_input := "aaabcddafff"
	test_expected_output := "3abc2da3f"
	test_output := RLE_compress(test_input)
	if test_output != test_expected_output {
		t.Error("Failted rle compression on string: " + test_input + ". Output was: " + test_output)
	}

	test_input = "bcssssdda"
	test_expected_output = "bc4s2da"
	test_output = RLE_compress(test_input)
	if test_output != test_expected_output {
		t.Error("Failted rle compression on string: " + test_input + ". Output was: " + test_output)
	}

	test_input = "bcsda"
	test_output = RLE_compress(test_input)
	if test_output != test_input {
		t.Error("Failted rle compression on string: " + test_input + ". Output was: " + test_input)
	}
}
