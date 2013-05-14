.PHONY: all prepare test clean parser lexer get-deps

all: prepare test
	go build -v ...

prepare: parser lexer
	@:

lexer:
	nex < lexer.nn > lexer.nn.go
	go fmt lexer.nn.go

parser:
	rm -f parser.y.go
	go tool yacc -o parser.y.go parser.y

get-deps:
	go get -d -v -u ... github.com/blynn/nex
	go install github.com/blynn/nex

clean:
	rm -rf *.output *.nn.go *.y.go

test:
	go test -v
