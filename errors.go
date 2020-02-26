package strongword

import (
	"fmt"
	"regexp"
)

// MinLengthError means the provided word is too short.
type MinLengthError struct {
	MinLength      int
	ProvidedLength int
}

func (err MinLengthError) Error() string {
	return fmt.Sprintf("minimum length is %d characters, only %d provided", err.MinLength, err.ProvidedLength)
}

// DigitsError means the provided word does not contain enough digits.
type DigitsError struct {
	Minimum  int
	Provided int
}

func (err DigitsError) Error() string {
	return fmt.Sprintf("must contain at least %d digits, but does only contain %d", err.Minimum, err.Provided)
}

// SpecialCharsError means the provided word does not contain enough special characters.
type SpecialCharsError struct {
	Minimum  int
	Provided int
}

func (err SpecialCharsError) Error() string {
	return fmt.Sprintf("must contain at least %d special characters, but does only contain %d", err.Minimum, err.Provided)
}

// RunesError means the provided word does not contain enough runes of a given set.
type RunesError struct {
	Runes    []rune
	Minimum  int
	Provided int
}

func (err RunesError) Error() string {
	return fmt.Sprintf("must contain at least %d of any of %v, but does only contain %d", err.Minimum, err.Runes, err.Provided)
}

// RegexpError means the provided word does match against a regular expression.
type RegexpError struct {
	Regexp *regexp.Regexp
}

func (err RegexpError) Error() string {
	return fmt.Sprintf("does not match against %s", err.Regexp)
}

// CharsAndDigitsError means the provided word does not contain either any character or any digit.
type CharsAndDigitsError struct {
	Detail string
}

func (err CharsAndDigitsError) Error() string {
	return err.Detail
}
