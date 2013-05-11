package edn

import "container/list"
import "fmt"
import "bytes"

type Vector []Value
type List list.List
type Map map[Value]Value
type Set map[Value]bool
type Int int
type String string

type Value interface {
	String() string
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

func (i Int) String() string {
	return fmt.Sprintf("%d", int(i))
}

func (s String) String() string {
	return fmt.Sprintf("%#v", string(s))
}
