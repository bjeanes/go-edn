default: clean edn

edn: parser.y.go lexer.nn.go
	go build

clean:
	rm -rf *.output *.nn.go *.y.go

test: clean edn
	go test

lexer.nn.go:
	nex lexer.nn
	go fmt lexer.nn.go

parser.y.go:
	go tool yacc -o parser.y.go parser.y

