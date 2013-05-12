package edn

import (
	. "testing"
)

func TestDoesNotParseEmptyInput(t *T) {
	_, err := ParseString("")

	if err == nil {
		t.Error("Expected parsing empty input to fail")
	}
}

func TestParseList(t *T) {
	val, err := ParseString("()")

	if err != nil {
		t.Errorf("Expected parsing \"()\" to succeed. %v", err)
	}

	if !new(List).Equals(val) {
		t.Errorf("Expected parsing \"()\" to return an empty list, got %+v", val)
	}

	str := `(() "abc" [] "def")`
	val, err = ParseString(str)

	if err != nil {
		t.Errorf("Expected parsing %#v to succeed. %v", str, err)
	}

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

func TestParseString(t *T) {
	val, err := ParseString(`""`)

	if err != nil {
		t.Errorf("Expected parsing '\"\"' to succeed. %v", err)
	}

	if val == nil || !val.Equals(String("")) {
		t.Errorf("Expected \"\", got %v", val)
	}

	val, err = ParseString(`"abc"`)

	if err != nil {
		t.Errorf("Expected parsing '\"abc\"' to succeed. %v", err)
	}

	if val == nil || !val.Equals(String("abc")) {
		t.Errorf("Expected \"abc\", got %v", val)
	}
}
