package edn

import (
	"bytes"
	"reflect"
)

type Vector []Value

//////// Value interface

func (vec Vector) Equals(val Value) bool {
	return reflect.DeepEqual(vec, val)
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

//////// Sequence interface:

// TODO: if `other` is a Vector, perform a copy instead of using .Each()
func (this Vector) Into(other Sequence) Sequence {
	seq := make(Vector, this.Length(), this.Length()+other.Length())
	copy(seq, this)

	f := func(val Value, _ int) {
		seq = append(seq, val)
	}

	other.Each(f)

	return seq
}

func (this Vector) Each(f func(Value, int)) {
	for i, v := range this {
		f(v, i)
	}
}

func (this Vector) Length() int {
	return len(this)
}
