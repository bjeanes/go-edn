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

func (l *List) Equals(val Value) bool {
	return reflect.DeepEqual(l, val)
}

func (l *List) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")

	for el := l.raw().Front(); el != nil; el = el.Next() {
		item := el.Value.(Value)
		buffer.WriteString(item.String())
		if el.Next() != nil {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString(")")
	return buffer.String()
}
