package edn

import . "testing"

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
