// Package template is the layout oracle for gomatic library repositories: a
// minimal, fully-conformant pure-Go library exemplifying the standards every
// library repo follows (named parameter types, constant sentinel errors, value
// semantics, 100% asserted coverage). Replace this package with the real
// library when instantiating the template.
package template

// Error is the package's constant error type; every failure the package can
// emit is a const of it, testable with errors.Is.
type Error string

// Error returns the sentinel's message.
func (e Error) Error() string { return string(e) }

// ErrEmptyName rejects a greeting for a blank name.
const ErrEmptyName Error = "name is empty"

// Name is the addressee of a greeting.
type Name string

// Greet renders the canonical greeting for name, rejecting a blank name with
// ErrEmptyName.
func Greet(name Name) (string, error) {
	if name == "" {
		return "", ErrEmptyName
	}
	return "Hello, " + string(name) + "!", nil
}
