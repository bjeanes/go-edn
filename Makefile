.PHONY: all setup test clean parser lexer deps

all: test
	go build

setup: deps parser lexer
	@:

lexer:
	rm -f lexer.nn.go
	nex lexer.nn
	go fmt lexer.nn.go

parser:
	rm -f parser.y.go
	go tool yacc -o parser.y.go parser.y

deps:
	go get -u ...

clean:
	rm -rf *.output *.nn.go *.y.go

test: clean setup
	go test -v
