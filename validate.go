package govalidators

import "regexp"

var nameRegex = regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s]+$`)

// IsName reports whether the given string contains only valid characters
// (letters and spaces, including accented characters like À-ÿ).
func IsName(s string) bool {
	return nameRegex.MatchString(s)
}

var uintRegex = regexp.MustCompile(`^\d+$`)

// IsUint reports whether the given string is a valid numeric Uint.
// Rejects any non-numeric input, including SQL injection attempts.
func IsUint(s string) bool {
    return uintRegex.MatchString(s)
}