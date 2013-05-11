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
	return "#{}"
}

func (m Map) String() string {
	return "{}"
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

func (l List) String() string {
	return "()"
}

func (i Int) String() string {
	return fmt.Sprintf("%d", int(i))
}

func (s String) String() string {
	return fmt.Sprintf("%#v", string(s))
}
