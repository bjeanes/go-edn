package edn

import (
	"bytes"
	linked "container/list"
	"reflect"
)

type List linked.List

func (this *List) raw() *linked.List {
	return (*linked.List)(this)
}

func (this *List) Insert(values ...Value) *List {
	for _, value := range values {
		this.raw().PushBack(value)
	}
	return this
}

//////// Value interface:

func (this *List) Equals(other Value) bool {
	return reflect.DeepEqual(this, other)
}

func (this *List) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")

	this.Each(func(item Value, i int) {
		buffer.WriteString(item.String())
		if i+1 < this.Length() {
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
	list := new(List)
	list.raw().Init()

	add := func(value Value, _ int) {
		list.raw().PushBack(value)
	}

	this.Each(add)
	other.Each(add)

	return new(List)
}
func (this *List) Length() int {
	return this.raw().Len()
}
