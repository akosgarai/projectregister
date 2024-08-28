package transformers

import (
	"strconv"
)

// StringToInt64 converts the input string to int64.
// If the conversion is not successful, it returns 0.
// This function is useful when we need to convert a request parameter to int64.
// The 0 value is not a valid identifier in the database, so it is a good indicator for the invalid conversion.
func StringToInt64(s string) int64 {
	integerValue, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return integerValue
}

// StringSliceToInt64Slice converts the input string slice to int64 slice.
// If the conversion for any element is not successful, it skips that element.
func StringSliceToInt64Slice(s []string) []int64 {
	result := make([]int64, 0)
	for _, v := range s {
		integerValue, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			result = append(result, integerValue)
		}
	}
	return result
}
