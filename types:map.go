package edn

import (
	"bytes"
	"reflect"
)

// Map represents an EDN map: {:a "b", 123 true}
type Map map[Value]Value

// Equals compares the Map to another Value for equality.
func (m Map) Equals(v Value) bool {
	return reflect.DeepEqual(m, v)
}

// String returns the EDN string representation of the Map.
func (m Map) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{")

	i := 0
	for k, v := range m {
		i++
		buffer.WriteString(k.String())
		buffer.WriteString(" ")
		buffer.WriteString(v.String())
		if i < len(m) {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString("}")
	return buffer.String()
}
