package strongword_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/bounoable/strongword"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	expr := regexp.MustCompile(`secret[0-9]{4}\*\*[^a-z]{2}`)

	cases := []struct {
		Test           string
		Password       string
		Rules          []strongword.Rule
		ExpectedErrors []error
	}{
		{
			Test:     "fail MinLength",
			Password: "short",
			Rules:    []strongword.Rule{strongword.MinLength(6)},
			ExpectedErrors: []error{strongword.MinLengthError{
				MinLength:      6,
				ProvidedLength: 5,
			}},
		},
		{
			Test:     "pass MinLength",
			Password: "definitelynotshort",
			Rules:    []strongword.Rule{strongword.MinLength(10)},
		},
		{
			Test:     "fail Digits",
			Password: "secret12",
			Rules:    []strongword.Rule{strongword.Digits(3)},
			ExpectedErrors: []error{strongword.DigitsError{
				Minimum:  3,
				Provided: 2,
			}},
		},
		{
			Test:     "pass Digits",
			Password: "secret1234",
			Rules:    []strongword.Rule{strongword.Digits(3)},
		},
		{
			Test:     "fail SpecialChars",
			Password: "secret#*",
			Rules:    []strongword.Rule{strongword.SpecialChars(3)},
			ExpectedErrors: []error{strongword.SpecialCharsError{
				Minimum:  3,
				Provided: 2,
			}},
		},
		{
			Test:     "pass SpecialChars",
			Password: "secret#*.",
			Rules:    []strongword.Rule{strongword.SpecialChars(3)},
		},
		{
			Test:     "fail Characters",
			Password: "secret",
			Rules:    []strongword.Rule{strongword.Characters("+-?", 1)},
			ExpectedErrors: []error{strongword.RunesError{
				Runes:    []rune{'+', '-', '?'},
				Minimum:  1,
				Provided: 0,
			}},
		},
		{
			Test:     "fail Characters (minimum)",
			Password: "secret+-",
			Rules:    []strongword.Rule{strongword.Characters("+-?", 3)},
			ExpectedErrors: []error{strongword.RunesError{
				Runes:    []rune{'+', '-', '?'},
				Minimum:  3,
				Provided: 2,
			}},
		},
		{
			Test:     "pass Characters",
			Password: "secret+-?",
			Rules:    []strongword.Rule{strongword.Characters("+-?", 3)},
		},
		{
			Test:     "pass Characters (repeated)",
			Password: "secret+++",
			Rules:    []strongword.Rule{strongword.Characters("+-?", 3)},
		},
		{
			Test:           "fail regular expression",
			Password:       "secret",
			Rules:          []strongword.Rule{strongword.Regexp(expr)},
			ExpectedErrors: []error{strongword.RegexpError{Regexp: expr}},
		},
		{
			Test:     "pass regular expression",
			Password: "secret3728**+8",
			Rules:    []strongword.Rule{strongword.Regexp(expr)},
		},
		{
			Test:           "fail CharsAndDigits (only digits)",
			Password:       "1345678",
			Rules:          []strongword.Rule{strongword.CharsAndDigits()},
			ExpectedErrors: []error{strongword.CharsAndDigitsError{}},
		},
		{
			Test:           "fail CharsAndDigits (only chars)",
			Password:       "secret",
			Rules:          []strongword.Rule{strongword.CharsAndDigits()},
			ExpectedErrors: []error{strongword.CharsAndDigitsError{}},
		},
		{
			Test:     "pass CharsAndDigits",
			Password: "secret123",
			Rules:    []strongword.Rule{strongword.CharsAndDigits()},
		},
		{
			Test:     "fail default rules (MinLength(8), CharsAndDigits(), SpecialChars(1))",
			Password: "secret12345678",
			ExpectedErrors: []error{strongword.SpecialCharsError{
				Minimum:  1,
				Provided: 0,
			}},
		},
		{
			Test:     "pass default rules (MinLength(8), CharsAndDigits(), SpecialChars(1))",
			Password: "secret123*",
		},
	}

	for _, c := range cases {
		t.Run(c.Test, func(t *testing.T) {
			errs := strongword.Validate(c.Password, c.Rules...)

			assert.Len(t, errs, len(c.ExpectedErrors))

			for _, exerr := range c.ExpectedErrors {
				var founderr error
				for _, err := range errs {
					if errors.As(err, &exerr) {
						founderr = err
					}
				}

				assert.NotNil(t, founderr)
			}
		})
	}
}

func ExampleValidate() {
	// Default rule set
	errs := strongword.Validate("weakpassword")

	// Custom rule set
	errs = strongword.Validate(
		"weakpassword",
		strongword.MinLength(12),
		strongword.SpecialChars(3),
		strongword.Regexp(regexp.MustCompile("(?)[0-9]{4}")),
	)

	for _, err := range errs {
		fmt.Println(err)
	}
}
