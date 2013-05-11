package edn

import . "testing"

func TestVectorString(t *T) {
	vec := make(Vector, 0)
	str := vec.String()

	if str != "[]" {
		t.Fail()
	}

	vec = append(vec, Int(1))
	str = vec.String()

	if str != "[1]" {
		t.Errorf("Expecting [1], received %+v", str)
	}

	vec = append(vec, make(Vector, 0), String("abc"))
	str = vec.String()

	if vec.String() != `[1 [] "abc"]` {
		t.Errorf(`Expecting [1 [] "abc"], received %+v`, str)
	}
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
