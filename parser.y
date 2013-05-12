%{
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
	: value { lastResult = Value($$.v) }
	;

value
	: list
	| vector
	| string
	| set
	| map
	;

string
	: tString { 
		s := string($$.v.(String))
		$$.v = String(s[1:len(s)-1]) 
	  }
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
	: tOpenParen tCloseParen { $$.v = new(List) }
	;

vector
	: tOpenBrace tOpenBrace { $$.v = Vector{} }
	;
%%