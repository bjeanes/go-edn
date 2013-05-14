package types

import (
	"bytes"
	"reflect"
)

// Map represents an EDN map: {:a "b", 123 true}
type Map map[Value]Value

// Equals compares the Map to another Value for equality.
func (this Map) Equals(other Value) bool {
	return reflect.DeepEqual(this, other)
}

// String returns the EDN string representation of the Map.
func (this Map) String() string {
	// TODO use collectionString("{", this, "}")
	//      Just need to figure make it meet Sequence interface

	var buffer bytes.Buffer
	buffer.WriteString("{")

	i := 0
	for k, v := range this {
		i++
		buffer.WriteString(k.String())
		buffer.WriteString(" ")
		buffer.WriteString(v.String())
		if i < len(this) {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString("}")
	return buffer.String()
}
