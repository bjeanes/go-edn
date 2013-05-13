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
