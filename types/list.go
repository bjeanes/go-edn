package types

import (
	linked "container/list"
	"reflect"
)

// List is a type that represents an EDN list, such as: (1 2 3)
type List linked.List

func (this *List) raw() *linked.List {
	return (*linked.List)(this)
}

// Insert inserts N Value items to the List and returns a pointer to the List.
func (this *List) Insert(values ...Value) *List {
	for _, value := range values {
		this.raw().PushBack(value)
	}
	return this
}

//////// Value interface:

// Equals compares the List to another Value for equality.
func (this *List) Equals(other Value) bool {
	return reflect.DeepEqual(this, other)
}

// String returns the EDN string representation of the List.
func (this *List) String() string {
	return collectionString("(", this, ")")
}

//////// Sequence interface:

// Each calls the provided function once for each item in its collection.
//
// The provided function is passed the item and its index: f(item, index).
func (this *List) Each(f func(Value, int)) {
	for el, i := this.raw().Front(), 0; el != nil; el, i = el.Next(), i+1 {
		v, _ := el.Value.(Value)
		f(v, i)
	}
}

// Into returns a new Sequence (backed by a List) with the items from both the
// current List and the other Sequence.
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

// Length provides the current item count from the List.
func (this *List) Length() int {
	return this.raw().Len()
}
