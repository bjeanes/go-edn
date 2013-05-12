package edn

import (
	"bytes"
	"container/list"
	"fmt"
	"reflect"
)

type Vector []Value
type List list.List
type Map map[Value]Value
type Set map[Value]bool
type Int int
type String string

type Value interface {
	String() string
	Equals(Value) bool
}

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

func (m Map) Equals(v Value) bool {
	panic("unimplemented")
	return false
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

func (vec Vector) Equals(val Value) bool {
	if v, ok := val.(Vector); ok {
		if len(vec) != len(v) {
			return false
		}

		return reflect.DeepEqual(vec, v)
	}

	return false
}

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
	list := (*list.List)(l)
	var buffer bytes.Buffer
	buffer.WriteString("(")

	for el := list.Front(); el != nil; el = el.Next() {
		item := el.Value.(Value)
		buffer.WriteString(item.String())
		if el.Next() != nil {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString(")")
	return buffer.String()
}

func (i Int) Equals(v Value) bool {
	panic("unimplemented")
	return false
}

func (i Int) String() string {
	return fmt.Sprintf("%d", int(i))
}

func (s String) Equals(v Value) bool {
	panic("unimplemented")
	return false
}

func (s String) String() string {
	return fmt.Sprintf("%#v", string(s))
}
