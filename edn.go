package edn

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func ParseString(string string) (val Value, err error) {
	val, err = ParseReader(strings.NewReader(string))
	return
}

func ParseReader(reader io.Reader) (val Value, err error) {
	defer func() {
		// Nex's parser calls panic() on a lexing error
		if r := recover(); r != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("Error: %v", r))
			}
		}
	}()

	if yyParse(NewLexer(reader)) == 1 {
		val = Vector([]Value{})
	} else {
		err = errors.New("Error: could not parse provided EDN")
	}

	return
}
