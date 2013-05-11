package edn

import . "testing"

func assertEqual(expect, actual interface{}, t *T) {
	if expect != actual {
		t.Errorf("Expecting %+v, received %+v", expect, actual)
	}
}

func TestVectorString(t *T) {
	vec := make(Vector, 0)
	str := vec.String()

	if str != "[]" {
		t.Fail()
	}

	vec = append(vec, Int(1))
	assertEqual("[1]", vec.String(), t)

	vec = append(vec, make(Vector, 0), String("abc"))
	assertEqual(`[1 [] "abc"]`, vec.String(), t)
}

func TestEmptyListString(t *T) {
	list := new(List)
	str := list.String()

	if str != "()" {
		t.Fail()
	}
}

func TestEmptyMapString(t *T) {
	_map := new(Map)
	str := _map.String()

	if str != "{}" {
		t.Fail()
	}
}

func TestEmptySetString(t *T) {
	set := new(Set)
	str := set.String()

	if str != "#{}" {
		t.Fail()
	}
}
