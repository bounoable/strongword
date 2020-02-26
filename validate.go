// Package strongword is a simple utility for validating password strengths against a provided set of rules.
package strongword

// Rule is a validation rule.
type Rule interface {
	Validate(string) error
}

// RuleFunc is a validation rule.
type RuleFunc func(string) error

// Validate validates the given word.
func (fn RuleFunc) Validate(word string) error {
	return fn(word)
}

// Validate validates the given word against the provided rules.
// If no rules are provided, the default rule set is used:
// MinLength(8), CharsAndDigits(), SpecialChars(1)
func Validate(word string, rules ...Rule) []error {
	if len(rules) == 0 {
		rules = DefaultRules
	}

	var errs []error
	for _, rule := range rules {
		if err := rule.Validate(word); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
