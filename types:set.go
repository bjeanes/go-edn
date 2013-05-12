package edn

import (
	"bytes"

//"reflect"
)

type Set map[Value]bool

func (s Set) Equals(v Value) bool {
	panic("unimplemented")
	return false
}

func (s Set) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("#{")

	i := 0
	for k, _ := range s {
		i++
		buffer.WriteString(k.String())
		if i < len(s) {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString("}")
	return buffer.String()
}

func (s Set) Insert(v Value) {
	s[v] = true
}
