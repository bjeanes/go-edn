# EDN

An [EDN][edn] parser for Go.

## Considerations

* Currently the `Map` type is an alias to `map[Value]Value` (`Value` being an 
  interface provided by this library). In Go, `map`s do not support all type 
  as keys. 

  Since `Set`s are currently backed by `map[Value]bool`, the same restriction
  applies to set elements.

  Notably, `Vector`, `Set`, `Map`, etc., can not be map keys or set elements,
  despite being a valid expression in EDN. 

   * Consider using the hash-set and hash-map implementations from 
     [gohash][gohash]. The trade-off here will be further departure from
     native types as well as requiring all EDN types to meet the 
     [`Hasher` and `Equalser` interfaces][gohash interfaces].

* [Go's YACC tool][yacc] does not allow customizing the
  signature of the generated parse function (such as I believe Bison does).
  This means that in order to extract the data out of the parser, global state
  is involved, making the parser non-parallelizable. I hope to discover an
  alternate way to achieve this goal that does not find global state. Help
  welcome!

## Building

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

Not yet licensed, mostly out of laziness. If you want to make use of this 
project today, just [contact me](mailto:me@bjeanes.com).

[edn]: https://github.com/edn-format/edn
[nex]: http://www-cs-students.stanford.edu/~blynn/nex/
[yacc]: http://golang.org/cmd/yacc/
[gohash]: https://code.google.com/p/gohash/
[gohash interfaces]: https://code.google.com/p/gohash/source/browse/hash/set.go#37
[irc discussion]: https://botbot.me/freenode/go-nuts/msg/3137158/
[go build advice]: http://golang.org/doc/articles/go_command.html#tmp_4