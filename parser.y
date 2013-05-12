%{
/*
 * If this file is not parser.y, it was generated from parser.y and
 * should not be edited directly.
 */
package edn 
import "fmt"

// Eww... global state. TODO: how else to get actual data out of from yyParse?
var lastResult Value

func init() {
	//yyDebug = 4
}
%}

%union { 
	v Value
} 

%token tOpenBracket tCloseBracket
%token tOpenParen tCloseParen
%token tOpenBrace tCloseBrace
%token tOctothorpe
%token tString

%% 
input
	: value { lastResult = Value($1.v) }
	;

value
	: list
	| vector
	| string
	| set
	| map
	;

ws
	: " "
	| "\t"
	| ","
	| "\r"
	| "\n"
	;

ws＋
	: ws
	| ws ws＋

ws✳
	: /* empty */
	| ws＋

values
	: ws✳ { $$.v = Value(new(List))}
	| values ws✳ value {
		$$.v.(*List).raw().PushBack($3.v)
	  }
	;

string
	: tString
	;

set
	: tOctothorpe tOpenBrace tCloseBrace
	  { $$.v = Set{} }
	;

map
	: tOpenBrace tCloseBrace
	  { $$.v = Map{} }
	;

list
	: tOpenParen values tCloseParen { $$.v = $2.v }
	;

vector
	: tOpenBracket values tCloseBracket { $$.v = Value(Sequence(Vector{}).Into(Sequence($2.v.(*List)))) }
	;
%%