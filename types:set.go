package edn

import (
	"bytes"
	"reflect"
)

// Set represents an EDN set: #{1 2 3}
type Set map[Value]bool

// Insert adds a non-present Value into the Set and returns the Set. Adding an
// existing Value will have no effect.
func (this Set) Insert(value Value) Sequence {
	this[value] = true
	return this
}

//////// Value interface:

// Equals compares the Set to another Value for equality.
func (this Set) Equals(other Value) bool {
	return reflect.DeepEqual(this, other)
}

// String returns the EDN string representation of the Set.
func (this Set) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("#{")

	i := 0
	for k, _ := range this {
		i++
		buffer.WriteString(k.String())
		if i < len(this) {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString("}")
	return buffer.String()
}

//////// Sequence interface:

// Length provides the current item count from the Set.
func (this Set) Length() int {
	return len(this)
}

// Into returns a new Sequence (backed by a Set) with the items from both the
// current Set and the other Sequence.
func (this Set) Into(other Sequence) Sequence {
	set := Set{}

	f := func(val Value, _ int) {
		set.Insert(val)
	}

	this.Each(f)
	other.Each(f)

	return set
}

// Each calls the provided function once for each item in its collection.
//
// The provided function is passed the item and its index: f(item, index).
func (this Set) Each(f func(Value, int)) {
	i := 0
	for item := range this {
		f(item, i)
		i++
	}
}
