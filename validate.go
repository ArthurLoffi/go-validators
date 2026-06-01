package govalidators

import "regexp"

var nameRegex = regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s]+$`)

// ValidateName reports whether the given string contains only valid characters
// (letters and spaces, including accented characters like À-ÿ).
func ValidateName(s string) bool {
	return nameRegex.MatchString(s)
}

var idRegex = regexp.MustCompile(`^\d+$`)

// ValidateID reports whether the given string is a valid numeric ID.
// Rejects any non-numeric input, including SQL injection attempts.
func ValidateID(s string) bool {
    return idRegex.MatchString(s)
}