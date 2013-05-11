package edn

import "io"

func ParseReader(reader io.Reader) Value {
	yyParse(NewLexer(reader))
	return Vector([]Value{})
}
