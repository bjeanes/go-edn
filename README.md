# EDN

A parser for EDN in Go. See https://github.com/edn-format/edn

## Building

The lexer is generated with [nex](http://www-cs-students.stanford.edu/~blynn/nex/) 
and the parser is generated with [Go's bundled `yacc` command](http://golang.org/cmd/yacc/).

Since the Go source code for each needs to be generated before this will build, a
`Makefile` is provided. Simply run `make` to handle everything.

**TODO:** Consider committing the generated files (despite being artifacts) so that
consumers of this library can just do `import edn "github.com/bjeanes/go-edn"` without
manually pulling down and installing the package. This seems to be
[the](http://golang.org/doc/articles/go_command.html#tmp_4) 
[expectation](https://botbot.me/freenode/go-nuts/msg/3137158/).

### Running Tests

Just run `make test` to re-generate the lexer, parser, and run the tests.
