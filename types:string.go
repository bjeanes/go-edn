package edn

import (
	"fmt"
)

type String string
type Character rune // TODO: possibly move out of string file

func (s String) Equals(v Value) bool {
	return s == v
}

func (s String) String() string {
	return fmt.Sprintf("%#v", string(s))
}

type Keyword string

// NOTE: this will allow a Keyword to == a String.
// TODO: Make Keyword a struct (namespace + word) that
//       is a distinct type.
func (this Keyword) Equals(other Value) bool {
	return this == other
}

func (this Keyword) String() string {
	return fmt.Sprintf(":%v", string(this))
}
