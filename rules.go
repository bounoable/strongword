package strongword

import (
	"regexp"
)

var (
	// DefaultRules are the rules that are used if none are provided.
	DefaultRules = []Rule{
		MinLength(8),
		CharsAndDigits(),
		SpecialChars(1),
	}
)

// MinLength validates the length of a word.
func MinLength(length int) Rule {
	return RuleFunc(func(word string) error {
		if len(word) < length {
			return MinLengthError{
				MinLength:      length,
				ProvidedLength: len(word),
			}
		}

		return nil
	})
}

// Digits validates the digit count in a word.
func Digits(min int) Rule {
	var expr = regexp.MustCompile("[0-9]")

	return RuleFunc(func(word string) error {
		matches := expr.FindAllStringIndex(word, -1)

		if len(matches) < min {
			return DigitsError{
				Minimum:  min,
				Provided: len(matches),
			}
		}

		return nil
	})
}

// SpecialChars validates the special character count in a word.
// Special characters are all charaters that are not latin-alphanumerical (a-z & 0-9).
func SpecialChars(min int) Rule {
	var expr = regexp.MustCompile("[^a-z0-9]")

	return RuleFunc(func(word string) error {
		matches := expr.FindAllStringIndex(word, -1)

		if len(matches) < min {
			return SpecialCharsError{
				Minimum:  min,
				Provided: len(matches),
			}
		}

		return nil
	})
}

// Runes validates that a word contains any of runes at least min times.
func Runes(runes []rune, min int) Rule {
	return RuleFunc(func(word string) error {
		var count int

		for _, char := range runes {
			for _, wordchar := range word {
				if char == wordchar {
					count++
				}
			}
		}

		if count < min {
			return RunesError{
				Runes:    runes,
				Minimum:  min,
				Provided: count,
			}
		}

		return nil
	})
}

// Characters validates that a word contains any unicode point if chars at least min times.
func Characters(chars string, min int) Rule {
	return Runes([]rune(chars), min)
}

// Regexp validates a word against a regular expression.
func Regexp(expr *regexp.Regexp) Rule {
	return RuleFunc(func(word string) error {
		if !expr.MatchString(word) {
			return RegexpError{
				Regexp: expr,
			}
		}

		return nil
	})
}

var (
	charExpr  = regexp.MustCompile("(?i)[a-z]")
	digitExpr = regexp.MustCompile("[0-9]")
)

// CharsAndDigits validates that a word does contain at least one character and one digit.
func CharsAndDigits() Rule {
	return RuleFunc(func(word string) error {
		if !charExpr.MatchString(word) {
			return CharsAndDigitsError{
				Detail: "no characters provided",
			}
		}

		if !digitExpr.MatchString(word) {
			return CharsAndDigitsError{
				Detail: "no digits provided",
			}
		}

		return nil
	})
}
