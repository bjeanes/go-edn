package edn

import . "testing"

func TestStackPushAndPopSingleItem(t *T) {
	stack := new(Stack)
	item := "my item"

	stack.Push(item)

	if stack.Pop() != item {
		t.Error("didn't get item back")
	}
}

func TestStackLength(t *T) {
	stack := new(Stack)
	item := "my item"

	stack.Push(item)
	stack.Push(item)
	stack.Push(item)

	if stack.Length() != 3 {
		t.Error("length incorrect")
	}
}

func TestStackPushAndPopMultipleItems(t *T) {
	stack := new(Stack)
	item1 := "my 1st item"
	item2 := "my 2nd item"

	stack.Push(item1)
	stack.Push(item2)

	test1 := stack.Pop()
	test2 := stack.Pop()

	if test1 != item2 || test2 != item1 {
		t.Error("not stacking")
	}
}
