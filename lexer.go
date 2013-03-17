package edn

import "fmt"
import "strings"

// A lot of this is based on: http://cuddle.googlecode.com/hg/talk/lex.html

type tokenType int

const (
	tEOF tokenType = 1 << iota
	tOpenParen
	tCloseParen
	tOpenBrace
	tCloseBrace
	tOpenBracket
	tCloseBracket
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
	reader   *strings.Reader
	start    int
	position int
	tokens   chan token
}

func (l *lexer) peek() (ch rune, size int, err error) {
	ch, size, err = l.reader.ReadRune()
	l.reader.UnreadRune()
	return
}

type stateFn func(*lexer) stateFn

var stateFns map[string]stateFn = map[string]stateFn{}

func init() {
	stateFns["start"] = func(l *lexer) stateFn {
		defaultFn := stateFns["start"]

		if (l.reader.Len() > 0) {
			ch, _, err := l.reader.ReadRune()

			if(err != nil) {
				return nil
			}

			switch ch {
				default: return defaultFn
				case '(':
					l.emit(tOpenParen)
					return defaultFn
				case ')':
					l.emit(tCloseParen)
					return defaultFn
			}
		}

		l.emit(tEOF)
		return nil
	}
}

func (l *lexer) run() {
	for t := stateFns["start"]; t != nil; {
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
		reader: strings.NewReader(input),
		tokens: make(chan token, 100),
	}

	go l.run()
	return l
}
