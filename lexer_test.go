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
	assertLexerYieldsCorrectTokens(t,
		"(1, 24223 ; i'm a comment\n 0 -4 {[5] #{2.34} \"ab\\\"c\"} '())",
		tokens(
			tOpenList, tNumber, tNumber, tComment, tNumber, tNumber,
			tOpenMap, tOpenVector, tNumber, tCloseVector, tOpenSet, tNumber,
			tCloseMapOrSet, tString, tCloseMapOrSet, tQuoteNextForm,
			tOpenList, tCloseList, tCloseList, tEOF),
		values("(", "1", "24223", "; i'm a comment", "0", "-4", "{",
			"[", "5", "]", "#{", "2.34", "}", "\"ab\\\"c\"", "}", "'",
			"(", ")", ")", ""))
}
