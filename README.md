# EDN

An [EDN][edn] parser for Go.

## Building

**Note:** Development on this library is done against Go 1.1. In particular,
Go 1.0's YACC differs from Go 1.1's YACC with regards to how `fmt` is imported.

The lexer is generated with [nex][nex]
and the parser is generated with [Go's YACC tool][yacc].

In order to be compatible with `go get`<sup>\[[1][irc discussion]\]</sup>,
which does not support a precompilation step<sup>\[[2][go build advice]\]</sup>
the generated files are committed alongside the source files.

The lexer and parser are re-created every time the project is made, which you
can simply do with:

```
make
```

This will also run the tests.

## License

MPL 2.0

[edn]: https://github.com/edn-format/edn
[nex]: http://www-cs-students.stanford.edu/~blynn/nex/
[yacc]: http://golang.org/cmd/yacc/
[gohash]: https://code.google.com/p/gohash/
[gohash interfaces]: https://code.google.com/p/gohash/source/browse/hash/set.go#37
[irc discussion]: https://botbot.me/freenode/go-nuts/msg/3137158/
[go build advice]: http://golang.org/doc/articles/go_command.html#tmp_4
