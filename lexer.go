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

func lexEDN(l *lexer) {
	ch, _, err := l.peek()

	if err != nil {
		return
	}

	switch ch {
	case '(':
		lexList(l)
	case '[':
		lexVector(l)
	case '{':
		lexMap(l)
	default:
		// TODO: proper error handling
		panic("Unexpected character " + fmt.Sprintf("'%c'", ch))
	}
}

func lexVector(l *lexer) {
	l.read()
	l.emit(tOpenBracket)

	for {
		ch, _, _ := l.peek()
		if ch == ']' {
			break
		}
		lexEDN(l)
	}

	l.read()
	l.emit(tCloseBracket)
}

func lexMap(l *lexer) {
	l.read()
	l.emit(tOpenBrace)

	for {
		ch, _, _ := l.peek()
		if ch == '}' {
			break
		}
		lexEDN(l) // key
		lexEDN(l) // value
	}

	l.read()
	l.emit(tCloseBrace)
}

func lexList(l *lexer) {
	l.read()
	l.emit(tOpenParen)

	for {
		ch, _, _ := l.peek()
		if ch == ')' {
			break
		}
		lexEDN(l)
	}

	l.read()
	l.emit(tCloseParen)
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

func (l *lexer) read() (ch rune, size int, err error) {
	ch, size, err = l.reader.ReadRune()
	l.position += size
	return
}

func (l *lexer) run() {
	lexEDN(l)
	l.emit(tEOF)
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
