# EDN

A parser for EDN in Go. See https://github.com/edn-format/edn

## Building

The parser is being written in [Ragel](http://www.complang.org/ragel/‎). The
`build.sh` in the root directory will compile the Ragel parser into Go, and
run the package tests.

## Contributing

Since Go does not have a way to run custom build steps (that I know of), in 
order to allow people to use this library directly from the Git source, ensure
that any changes to the Ragel source file(s) are committed along with the 
**regenerated** corresponding Go source files.

Personally, I'd prefer to never commit artifacts (which `ragel_parser.go` is,
in this case), but I don't currently know of a way to avoid doing so without
breaking the standard Go ecosystem.

## License

This code is currently unlicensed, but I expect to consider licenses soon. If
you are interested in using and/or contributing to the library in the 
meantime, please just [contact me](me@bjeanes.com).