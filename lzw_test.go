package gocompress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLZW_encoding(t *testing.T) {
	test_input := "Hello there"
	var test_expected_output []int
	test_expected_output = []int{72, 101, 108, 108, 111, 32, 116, 104, 101, 114, 101}

	test_output := LZW_encoding(test_input)

	assert.Equal(t, test_output, test_expected_output)
}

func TestLZW_decompression(t *testing.T) {
	test_expected_output := "Hello there"
	var test_input []int
	test_input = []int{72, 101, 108, 108, 111, 32, 116, 104, 101, 114, 101}

	test_output := LZW_decompression(test_input)

	assert.Equal(t, test_output, test_expected_output)
}

func TestLZW_lossless_compression(t *testing.T) {
	test_input := "In a hole in the ground there lived a hobbit"

	test_output := LZW_decompression(LZW_encoding(test_input))

	assert.Equal(t, test_output, test_input)
}
