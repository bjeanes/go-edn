package edn

import (
	. "github.com/bjeanes/go-edn/types"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ParseString is like ParseReader except it takes a string.
// See ParseReader for more details.
func ParseString(string string) (val Value, err error) {
	val, err = ParseReader(strings.NewReader(string))
	return
}

// ParseReader parses EDN from an io.Reader.
//
// Data is returned as a Value in the first return value. 
// The second return value is nil on successful parses, and
// an error on unsuccessful parses (e.g. syntax error).
func ParseReader(reader io.Reader) (val Value, err error) {
	defer func() {
		// Nex's parser calls panic() on a lexing error
		if r := recover(); r != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("Error: %v", r))
			}
		}
	}()

	lexer := NewLexer(reader)
	if yyParse(lexer) == 0 {
		val = lastResult
	} else {
		err = errors.New("Error: could not parse provided EDN")
	}

	return
}
