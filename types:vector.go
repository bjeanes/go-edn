package edn

import (
	"bytes"
	"reflect"
)

// Vector represents an EDN vector: [1 2 3]
type Vector []Value

//////// Value interface:

// Equals compares the Vector to another Value for equality.
func (vec Vector) Equals(val Value) bool {
	return reflect.DeepEqual(vec, val)
}

// String returns the EDN string representation of the Vector.
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

// Into returns a new Sequence (backed by a Vector) with the items from both 
// the current Vector and the other Sequence.
func (this Vector) Into(other Sequence) Sequence {
	// TODO: if `other` is a Vector, perform a copy instead of using .Each()

	seq := make(Vector, this.Length(), this.Length()+other.Length())
	copy(seq, this)

	other.Each(func(val Value, _ int) {
		seq = append(seq, val)
	})

	return seq
}

// Each calls the provided function once for each item in its collection.
//
// The provided function is passed the item and its index: f(item, index).
func (this Vector) Each(f func(Value, int)) {
	for i, v := range this {
		f(v, i)
	}
}

// Length provides the current item count from the Vector.
func (this Vector) Length() int {
	return len(this)
}
