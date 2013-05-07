package edn

import . "testing"

func TestEmptyVector(t *T) {
	edn, err := ParseString("[]")

	if err != nil {
		t.Fail()
	}

	arr := edn.([]EDN)

	if len(arr) != 0 {
		t.Fail()
	}
}
