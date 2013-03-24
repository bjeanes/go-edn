package edn

import . "testing"

func assertLexerYieldsCorrectTokens(
	t *T,
	source string,
	types []tokenType,
	values []string) {

	if len(types) != len(values) {
		t.Error("Expecting equal number of types and values")
		return
	}

	tokens := make([]token, 0)

	for token := range Lex(source).tokens {
		tokens = append(tokens, token)
	}

	if len(tokens) != len(types) {
		t.Errorf("Got %d tokens, expecting %d:\n\n%#v\n%#v",
			len(tokens),
			len(values),
			tokens, values)
		return
	}

	for i, actual := range tokens {
		expected := token{
			kind:  types[i],
			value: values[i],
		}

		if actual != expected {
			t.Errorf("Expecting %#v; actual %#v", expected, actual)
		}
	}
}

func tokens(tokens ...tokenType) []tokenType {
	return tokens
}
func values(values ...string) []string {
	return values
}

func TestEmptyGivesOnlyEOF(t *T) {
	assertLexerYieldsCorrectTokens(t,
		"",
		tokens(tEOF),
		values(""))
}

func TestOpenCloseLists(t *T) {
	assertLexerYieldsCorrectTokens(t,
		"()",
		tokens(tOpenList, tCloseList, tEOF),
		values("(", ")", ""))
}

func TestComplex(t *T) {
	// edn := "(1, 24223 ; i'm a comment\n 0 -4 {[5] #{2.34} \"ab\\\"c\"} '() " +
	edn := `(
		1, 24223 ; i'm a comment
		0 -4 {[5] #{2.34} "ab\"c"}
		'()
		2N, 2M, 2.3M, 2e3, 2E3, 2e-3,
		2e+3, -3e+2, -3e+4M, +3.2E-4
	)`
	assertLexerYieldsCorrectTokens(t,
		edn,
		tokens(
			tOpenList, tNumber, tNumber, tComment, tNumber, tNumber,
			tOpenMap, tOpenVector, tNumber, tCloseVector, tOpenSet, tNumber,
			tCloseMapOrSet, tString, tCloseMapOrSet, tQuoteNextForm,
			tOpenList, tCloseList, tNumber, tNumber, tNumber, tNumber,
			tNumber, tNumber, tNumber, tNumber, tNumber, tNumber,
			tCloseList, tEOF),
		values("(", "1", "24223", "; i'm a comment", "0", "-4", "{",
			"[", "5", "]", "#{", "2.34", "}", `"ab\"c"`, "}", "'",
			"(", ")", "2N", "2M", "2.3M", "2e3", "2E3", "2e-3", "2e+3",
			"-3e+2", "-3e+4M", "+3.2E-4", ")", ""))
}
