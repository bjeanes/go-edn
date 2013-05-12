package edn

import (
	//"bytes"
	"fmt"
	//"reflect"
)

type Int int

func (i Int) Equals(v Value) bool {
	panic("unimplemented")
	return false
}

func (i Int) String() string {
	return fmt.Sprintf("%d", int(i))
}
