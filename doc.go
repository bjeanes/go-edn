/*
An EDN (http://edn-format.org) parser for Go.

The main entry points for this library are ParseString(string) and 
ParseReader(io.Reader), which both return a representation of the data 
according to this library's Value interface.

	edn.ParseString(`(:a [123 #{5 7 3}] {"string" 123})`)
*/
package edn
