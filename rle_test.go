package gocompress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRlE_compression(t *testing.T) {
	test_input := "aaabcddafff"
	test_expected_output := "3abc2da3f"
	test_output := RLE_compress(test_input)

	assert.Equal(t, test_output, test_expected_output)

	test_input = "bcssssdda"
	test_expected_output = "bc4s2da"
	test_output = RLE_compress(test_input)

	assert.Equal(t, test_output, test_expected_output)

	test_input = "bcsda"
	test_output = RLE_compress(test_input)

	assert.Equal(t, test_output, test_input)
}

func TestRlE_decompression(t *testing.T) {
	test_input := "3afg2ad6c"
	test_expected_output := "aaafgaadcccccc"
	test_output := RLE_decompress(test_input)

	assert.Equal(t, test_output, test_expected_output)

	test_input = "aadgsggggeq"
	test_output = RLE_decompress(test_input)

	assert.Equal(t, test_output, test_input)
}
