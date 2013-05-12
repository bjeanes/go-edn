package edn

import (
	"bytes"
	"reflect"
)

type Map map[Value]Value

func (m Map) Equals(v Value) bool {
	return reflect.DeepEqual(m, v)
}

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
