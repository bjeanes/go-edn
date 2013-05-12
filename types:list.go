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
	if v, ok := val.(*List); ok {
		if (*list.List)(v).Len() != (*list.List)(l).Len() {
			return false
		}

		return reflect.DeepEqual(l, v)
	}

	return false
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
