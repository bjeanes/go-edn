package edn

import . "testing"

func assertLexerYieldsCorrectTokens(
	t *T,
	source string,
	types []tokenType,
	values []string) {

	tokens := make([]token, 0)

	for token := range Lex(source).tokens {
		tokens = append(tokens, token)
	}

	if len(tokens) != len(types) {
		t.Errorf("Got %d tokens, expecting %d", len(tokens), len(types))
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

func TestOpenCloseParens(t *T) {
	assertLexerYieldsCorrectTokens(t,
		"()",
		tokens(tOpenParen, tCloseParen, tEOF),
		values("(", ")", ""))
}

func TestComplex(t *T) {
	assertLexerYieldsCorrectTokens(t,
		"(1, 24223 {[5] \"abc\"} '())",
		tokens(
			tOpenParen, tNumber, tNumber, tOpenBrace,
			tOpenBracket, tNumber, tCloseBracket, tString,
			tCloseBrace, tQuoteNextForm, tOpenParen,
			tCloseParen, tCloseParen, tEOF),
		values("(", "1", "24223", "{", "[", "5", "]", "\"abc\"",
			"}", "'", "(", ")", ")", ""))
}
