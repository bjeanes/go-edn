package edn

import (
	"bytes"
	"container/list"
	"reflect"
)

type List list.List

func (l *List) raw() *list.List {
	return (*list.List)(l)
}

//////// Value interface

func (l *List) Equals(val Value) bool {
	return reflect.DeepEqual(l, val)
}

func (l *List) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")

	Sequence(l).Each(func(item Value, i int) {
		buffer.WriteString(item.String())
		if i+1 < Sequence(l).Length() {
			buffer.WriteString(" ")
		}
	})

	buffer.WriteString(")")
	return buffer.String()
}

//////// Sequence interface:

func (this *List) Each(f func(Value, int)) {
	for el, i := this.raw().Front(), 0; el != nil; el, i = el.Next(), i+1 {
		f(el.Value.(Value), i)
	}
}

func (this *List) Into(other Sequence) Sequence {
	return new(List)
}
func (this *List) Length() int {
	return this.raw().Len()
}
