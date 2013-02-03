package edn

import . "testing"

func TestEmptyGivesOnlyEOF(t *T) {
	lexer := Lex("")
	token, _ := lexer.Next()

	if token.kind != tEOF {
		t.Error("expecting EOF")
	}
}

// I suspect there's a potential race condition here first since
// the lexer is in a different thread. If `Next()` is called while the lexer
// is still in its main `for{}` loop, `done` could still be `false`
func TestEmptyIsDoneAfterFirstToken(t *T) {
	lexer := Lex("")
	_, done := lexer.Next()

	if !done {
		t.Error("expecting no more tokens")
	}
}
