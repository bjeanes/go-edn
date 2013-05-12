package edn

import (
	"fmt"
)

type Int int

func (i Int) Equals(v Value) bool {
	return i == v
}

func (i Int) String() string {
	return fmt.Sprintf("%d", int(i))
}
