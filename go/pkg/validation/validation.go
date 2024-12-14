package validation

import (
	"fmt"
	"math"
)

// IsEmptyString checks if a string is empty.
func IsEmptyString(value, fieldName string) error {
	if value == "" {
		return fmt.Errorf("%s cannot be empty", fieldName)
	}
	return nil
}

// IsInRange checks if a int out of range.
func IsInRange(value, min, max int, fieldName string) error {

	if min == math.MinInt && max == math.MaxInt {
		return nil
	}
	if min != math.MinInt && value < min {
		return fmt.Errorf("%s must be greater than or equal to %d", fieldName, min)
	}
	if max != math.MaxInt && value > max {
		return fmt.Errorf("%s must be less than or equal to %d", fieldName, max)
	}
	return nil
}
