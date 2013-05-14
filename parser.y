%{
package edn 

/*
If this file is not parser.y, it was generated from parser.y and
should not be edited directly.
*/

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
%token tString tKeyword
%token tWhitespace

%% 
input /* somebody help me not have to do this: */
	: ws✳ value ws✳ { lastResult = Value($2.v) }
	;

value
	: list
	| vector
	| string
	| set
	| map
	| keyword
	;

ws
	: tWhitespace
	;

ws＋
	: ws
	| ws ws＋

ws✳
	: /* empty */
	| ws＋

values
	: ws✳ { $$.v = Value(new(List))}
	| values ws✳ value ws✳ {
		$1.v.(*List).raw().PushBack($3.v)
	  }
	;

keyword
	: tKeyword
	;

string
	: tString
	;

set
	: tOctothorpe tOpenBrace values tCloseBrace
	  { $$.v = Sequence(Set{}).Into(Sequence($3.v.(*List))) }
	;

map
	: tOpenBrace tCloseBrace
	  { $$.v = Map{} }
	;

list
	: tOpenParen values tCloseParen { $$.v = $2.v }
	;

vector
	: tOpenBracket values tCloseBracket { $$.v = Sequence(Vector{}).Into(Sequence($2.v.(*List))) }
	;
%%