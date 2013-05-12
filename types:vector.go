package edn

import (
	"bytes"
	"reflect"
)

type Vector []Value

func (vec Vector) Equals(val Value) bool {
	if v, ok := val.(Vector); ok {
		if len(vec) != len(v) {
			return false
		}

		return reflect.DeepEqual(vec, v)
	}

	return false
}

func (vec Vector) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	for i, val := range vec {
		buffer.WriteString(val.String())
		if i < len(vec)-1 {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString("]")
	return buffer.String()
}
