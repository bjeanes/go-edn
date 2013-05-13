package edn

import (
	"bytes"
	"reflect"
)

type Set map[Value]bool

func (s Set) Insert(v Value) Sequence {
	s[v] = true
	return s
}

//////// Value interface:

func (s Set) Equals(v Value) bool {
	return reflect.DeepEqual(s, v)
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

//////// Sequence interface:

func (this Set) Length() int {
	return len(this)
}

func (this Set) Into(other Sequence) Sequence {
	set := Set{}

	f := func(val Value, _ int) {
		set.Insert(val)
	}

	this.Each(f)
	other.Each(f)

	return set
}

func (this Set) Each(f func(Value, int)) {
	i := 0
	for item := range this {
		f(item, i)
		i++
	}
}
