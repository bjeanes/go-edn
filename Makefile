.PHONY: all prepare test clean parser lexer get-deps

all: prepare test
	go build -v ./...

prepare: parser lexer
	@:

lexer:
	nex < lexer.nn > lexer.nn.go
	go fmt lexer.nn.go

	# Make "Lexer" private by renaming it... :S
	sed  -i '' 's:Lexer:lexer:g' lexer.nn.go
	sed  -i '' 's:Newlexer:newLexer:g' lexer.nn.go

parser:
	rm -f parser.y.go
	go tool yacc -o parser.y.go parser.y

	# Make "yyParse" accept a reference to a result
	sed -i '' 's:yyParse(yylex yyLexer) int:yyParse(yylex yyLexer, result *yySymType) int:' parser.y.go

get-deps:
	go get -d -v -u github.com/blynn/nex
	go install github.com/blynn/nex
	go install github.com/bjeanes/go-edn/types

clean:
	rm -rf *.output *.nn.go *.y.go

test:
	go test -v ./...
