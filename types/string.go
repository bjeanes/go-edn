package types

import (
	"fmt"
)

// String represents an EDN string: "abc"
type String string

// Equals compares the String to another Value for equality.
func (this String) Equals(other Value) bool {
	return this == other
}

// String returns the EDN string representation of the String.
func (this String) String() string {
	return fmt.Sprintf("%#v", string(this))
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

// Equals compares the Character to another Value for equality.
func (this Character) Equals(other Value) bool {
	return this == other
}

// String returns the EDN string representation of the Character.
func (this Character) String() string {
	var str string
	switch rune(this) {
	case '\t':
		str = "tab"
	case '\r':
		str = "return"
	case '\n':
		str = "newline"
	default:
		str = string(this)
	}
	return fmt.Sprintf("\\%s", str)
}
