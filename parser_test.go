package edn

import (
	. "testing"
)

func parse(s string, t *T) (val Value) {
	val, err := ParseString(s)

	if err != nil {
		t.Errorf("Expected parsing %+v to succeed. %v", s, err)
	}

	return
}

func assertValueEqual(actual, expected Value, t *T) {
	if !expected.Equals(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDoesNotParseEmptyInput(t *T) {
	_, err := ParseString("")

	if err == nil {
		t.Error("Expected parsing empty input to fail")
	}
}

func TestParseList(t *T) {
	val := parse("()", t)

	if !new(List).Equals(val) {
		t.Errorf("Expected parsing \"()\" to return an empty list, got %+v", val)
	}

	str := `(() "abc" [] "def")`
	val = parse(str, t)

	l := new(List)
	ll := l.raw()
	ll.PushBack(new(List))
	ll.PushBack(String("abc"))
	ll.PushBack(Vector{})
	ll.PushBack(String("def"))

	if !l.Equals(val) {
		t.Errorf("Expected %v, got %v", l, val)
	}
}

func TestParseVector(t *T) {
	assertValueEqual(parse(`[]`, t), Vector{}, t)
	assertValueEqual(parse(`[[]]`, t), Vector{Vector{}}, t)
	l := new(List)
	l.raw().PushBack(String("abc"))
	assertValueEqual(parse(`[[("abc")] "def"]`, t), Vector{Vector{l}, String("def")}, t)
}

func TestParseString(t *T) {
	val := parse(`""`, t)
	if val == nil || !val.Equals(String("")) {
		t.Errorf("Expected \"\", got %v", val)
	}

	val = parse(`"abc"`, t)
	if val == nil || !val.Equals(String("abc")) {
		t.Errorf("Expected \"abc\", got %v", val)
	}
}
