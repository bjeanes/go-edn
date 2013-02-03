package edn

import "fmt"

// A lot of this is based on: http://cuddle.googlecode.com/hg/talk/lex.html

type tokenType int

const (
	tEOF tokenType = 1 << iota
	tRightParen
	tLeftParen
	tRightBrace
	tLeftBrace
	tRightBracket
	tLeftBracket
	tSingleQuote
	tDoubleQuote
	tSymbol
	tKeyword
	tString
	tNumber
	tWhitespace // needed?
	tComment    // needed?

	tError
)

type stateFn func(*lexer) stateFn

type token struct {
	kind  tokenType
	value string
}

func (t *token) String() string {
	switch t.kind {
	case tEOF:
		return "EOF"
	case tError:
		return t.value
	default:
		return fmt.Sprintf("%q", t.value)
	}

	panic("unreachable")
}

type lexer struct {
	input    string
	start    int
	position int
	tokens   chan token
}

// Lexer states:
func lex(l *lexer) stateFn {
	l.emit(tEOF)
	return nil
}

func (l *lexer) run() {
	for t := lex; t != nil; {
		t = t(l)
	}

	close(l.tokens)
}

func (l *lexer) emit(tt tokenType) {
	value := l.input[l.start:l.position]
	l.start = l.position
	l.tokens <- token{kind: tt, value: value}
}

func (l *lexer) Next() (t token, more bool) {
	t, more = <-l.tokens
	return
}

func Lex(input string) *lexer {
	l := &lexer{
		input:  input,
		tokens: make(chan token, 100),
	}

	go l.run()
	return l
}
