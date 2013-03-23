package edn

import "fmt"
import "strings"
import "unicode"
import re "regexp"

// A lot of this is based on: http://cuddle.googlecode.com/hg/talk/lex.html

type tokenType int

const (
	tEOF tokenType = 1 << iota
	tOpenList
	tCloseList
	tOpenMap
	tCloseMapOrSet
	tOpenVector
	tCloseVector
	tQuoteNextForm
	tOpenSet
	tString
	tCharacter
	tSymbol
	tKeyword
	tNumber
	tMetadata
	tIgnoreNextForm
	tWhitespace // needed?
	tComment

	tError
)

const (
	// From clojure.git@4671019d7:src/jvm/clojure/lang/EdnReader.java
	intPattern   = "([-+]?)((0)|([1-9][0-9]*)|0[xX]([0-9A-Fa-f]+)|0([0-7]+)|([1-9][0-9]?)[rR]([0-9A-Za-z]+)|0[0-9]+)(N)?"
	floatPattern = "([-+]?[0-9]+(\\.[0-9]*)?([eE][-+]?[0-9]+)?)(M)?"

	numberPattern = "(" + floatPattern + "|" + intPattern + ")"
	stringPattern = "\"(\\\\\"|[^\"\\\\])*?\"" // TODO: do multiline strings work?
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

func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch) || ch == ','
}

func lexEDN(l *lexer) {
	for {
		ch, _, err := l.read()

		if err != nil {
			return
		}

		switch ch {
		case '(':
			l.emit(tOpenList)
		case '[':
			l.emit(tOpenVector)
		case '{':
			l.emit(tOpenMap)
		case ')':
			l.emit(tCloseList)
		case ']':
			l.emit(tCloseVector)
		case '}':
			l.emit(tCloseMapOrSet)
		case ':': // keyword
			// TODO: Read keyword value
			l.emit(tKeyword)
		case '\\':
			// TODO: read character
			l.emit(tCharacter)
		case '\'':
			l.emit(tQuoteNextForm)
		case '^':
			l.emit(tMetadata)
		case '#': // ...
			ch, _, _ := l.read()

			switch ch {
			case '_':
				l.emit(tIgnoreNextForm)
			case '{':
				l.emit(tOpenSet)
			case '^':
				l.emit(tMetadata)
			}
		case '"':
			l.readTokenMatchingRegexp(stringPattern)
			l.emit(tString)
		case ';':
			for {
				ch, _, _ := l.read()
				if ch == '\n' {
					l.unread()
					break
				}
			}
			l.emit(tComment)
		case '\t', '\n', ',', ' ': // whitespace
			for {
				ch, _, _ := l.read()
				if !isWhitespace(ch) {
					l.unread()
					break
				}
			}
			// l.emit(tWhitespace)
			l.start = l.position // Temporarily ignore whitespace
		case '+', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.readTokenMatchingRegexp(numberPattern)
			l.emit(tNumber)
		default:
			// TODO: proper error handling
			pad := ""
			for i := 0; i < l.position; i++ {
				pad = pad + " "
			}

			err := "Unexpected character " + fmt.Sprintf("'%c'", ch) + ":\n\n" +
			    // TODO only show the current line of input
				l.input + "\n" + pad + "^"
			panic(err)
		}
	}
}

type lexer struct {
	input        string
	reader       *strings.Reader
	start        int
	position     int
	lastRuneSize int
	tokens       chan token
}

func (l *lexer) peek() (ch rune, size int, err error) {
	ch, size, err = l.read()
	l.unread()
	return
}

// FIXME: can only unread one rune
func (l *lexer) unread() {
	l.reader.UnreadRune()
	l.position -= l.lastRuneSize
	l.lastRuneSize = 0
}

func (l *lexer) read() (ch rune, size int, err error) {
	ch, size, err = l.reader.ReadRune()
	l.lastRuneSize = size
	l.position += size
	return
}

func (l *lexer) readTokenMatchingRegexp(pattern string) {
	pattern = "^" + pattern
	reg := re.MustCompilePOSIX(pattern)

	indexes := reg.FindStringIndex(l.input[l.position-1:])

	if indexes == nil {
		panic("No match")
	}

	if indexes[0] != 0 {
		// shouldn't happen due to anchoring
		panic("Unexpected regex match")
	}

	for i := 0; i < indexes[1]-1; i++ {
		l.read()
	}
}

func (l *lexer) run() {
	lexEDN(l)
	l.emit(tEOF)
	close(l.tokens)
}

func (l *lexer) nextValue() string {
	return l.input[l.start:l.position]
}

func (l *lexer) emit(tt tokenType) {
	value := l.nextValue()
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
