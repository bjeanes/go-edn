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

func TestParseEmptyList(t *T) {
	val, err := ParseString("()")

	if err != nil {
		t.Errorf("Expected parsing \"()\" to succeed. %v", err)
	}

	if val == nil || !val.Equals(new(List)) {
		t.Errorf("Expected parsing \"()\" to return an empty list, got %v", val)
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
