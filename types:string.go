package edn

import (
	"fmt"
)

// String represents an EDN string: "abc"
type String string

// Equals compares the String to another Value for equality.
func (s String) Equals(v Value) bool {
	return s == v
}

// String returns the EDN string representation of the String.
func (s String) String() string {
	return fmt.Sprintf("%#v", string(s))
}

// Keyword represents an EDN keyword: :foo/bar or :baz
type Keyword string

func (this Keyword) Equals(other Value) bool {
	// NOTE: this will allow a Keyword to == a String.
	// TODO: Make Keyword a struct (namespace + word) that
	//       is a distinct type.

	return this == other
}

// String returns the EDN string representation of the Keyword.
func (this Keyword) String() string {
	return fmt.Sprintf(":%v", string(this))
}

// Character represents an EDN character: \a, \b, \c, \newline, etc...
type Character rune // TODO: possibly move out of string file
