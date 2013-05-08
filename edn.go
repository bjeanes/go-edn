package edn

import "io"

type EDN interface {
}

func ParseEDN(reader io.Reader) EDN {
	yyParse(NewLexer(reader))
	return []EDN{}
}
