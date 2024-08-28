package transformers

import (
	"testing"
)

// TestStringToInt64 tests the StringToInt64 function.
func TestStringToInt64(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1", 1},
		{"0", 0},
		{"a", 0},
		{"", 0},
	}
	for _, test := range tests {
		result := StringToInt64(test.input)
		if result != test.expected {
			t.Errorf("The conversion is not successful. Input: %s, Expected: %d, Got: %d", test.input, test.expected, result)
		}
	}
}

// TestStringSliceToInt64Slice tests the StringSliceToInt64Slice function.
func TestStringSliceToInt64Slice(t *testing.T) {
	tests := []struct {
		input    []string
		expected []int64
	}{
		{[]string{"1", "2", "3"}, []int64{1, 2, 3}},
		{[]string{"1", "a", "3"}, []int64{1, 3}},
		{[]string{"1", "", "3"}, []int64{1, 3}},
	}
	for _, test := range tests {
		result := StringSliceToInt64Slice(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("The length of the result is not the same as the expected. Input: %v, Expected: %v, Got: %v", test.input, test.expected, result)
		}
		for i, v := range test.expected {
			if v != result[i] {
				t.Errorf("The conversion is not successful. Input: %v, Expected: %v, Got: %v", test.input, test.expected, result)
			}
		}
	}
}
