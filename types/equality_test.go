package types

import . "testing"

func TestStringEquals(t *T) {
	if !String("abc").Equals(String("abc")) {
		t.Fail()
	}

	if String("abc").Equals(String("def")) {
		t.Fail()
	}
}

func TestSetEquals(t *T) {
	a, b := make(Set, 0), make(Set, 0)

	if !a.Equals(b) {
		t.Fail()
	}

	a.Insert(Int(123))

	if a.Equals(b) {
		t.Fail()
	}

	b.Insert(Int(123))

	if !a.Equals(b) {
		t.Fail()
	}
}

func TestMapEquals(t *T) {
	a, b := Map{String("a"): String("a")},
		Map{String("b"): String("b")}
	if a.Equals(b) {
		t.Fail()
	}

	b = Map{String("a"): String("a")}
	if !a.Equals(b) {
		t.Fail()
	}

	b = Map{}
	if a.Equals(b) {
		t.Fail()
	}
}

func TestIntEquals(t *T) {
	a, b := Int(1), Int(1)
	if !a.Equals(b) {
		t.Fail()
	}
}
