# EDN

A parser for EDN in Go. See https://github.com/edn-format/edn

## Building

The lexer is generated with [nex](http://www-cs-students.stanford.edu/~blynn/nex/) 
and the parser is generated with [Go's bundled `yacc` command](http://golang.org/cmd/yacc/).

In order to be compatible with `go get`<sup>\[[1](https://botbot.me/freenode/go-nuts/msg/3137158/)\]</sup>,
which does not support a precompilation step<sup>\[[2](http://golang.org/doc/articles/go_command.html#tmp_4)\]</sup>
the generated files are committed alongside the source files.

The lexer and parser are re-created every time the project is made, which you can simply do with:

```
make
```

This will also run the tests.

## License

Not yet licensed, mostly out of laziness. If you want to make use of this project today, just
[contact me](mailto:me@bjeanes.com).