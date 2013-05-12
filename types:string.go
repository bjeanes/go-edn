package edn

import (
	//"bytes"
	"fmt"
	//"reflect"
)

type String string

func (s String) Equals(v Value) bool {
	panic("unimplemented")
	return false
}

func (s String) String() string {
	return fmt.Sprintf("%#v", string(s))
}
