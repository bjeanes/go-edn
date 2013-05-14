%{
package edn 

/*
If this file is not parser.y, it was generated from parser.y and
should not be edited directly.
*/

import "github.com/bjeanes/go-edn/types"

func init() {
	//yyDebug = 4
}
%}

%union { 
	k types.Value
	v types.Value
} 

%token tOpenBracket tCloseBracket
%token tOpenParen tCloseParen
%token tOpenBrace tCloseBrace
%token tOctothorpe
%token tString tKeyword tCharacter
%token tWhitespace

%% 
input /* somebody help me not have to do this: */
	: ws✳ value ws✳ { result.v = $2.v }
	;

value
	: list
	| vector
	| string
	| set
	| map
	| keyword
	| character
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
	: ws✳ { $$.v = types.Value(new(types.List))}
	| values ws✳ value ws✳ {
		$1.v.(*types.List).Insert($3.v)
	  }
	;

character
	: tCharacter
	;

keyword
	: tKeyword
	;

string
	: tString
	;

set
	: tOctothorpe tOpenBrace values tCloseBrace
	  { 
	  	set := types.Sequence(types.Set{})
	  	values := types.Sequence($3.v.(*types.List))
	  	$$.v = set.Into(values)
	  }
	;

key_value
	: value ws✳ value { $$.k = $1.v; $$.v = $3.v }
	| value { yylex.Error("Map literal must contain an even number of forms") }
	;

key_values
	: ws✳ { $$.v = types.Map{} }
	| key_values key_value ws✳
	  {
	  	$1.v.(types.Map)[$2.k] = $2.v
	  }
	;

map
	: tOpenBrace key_values tCloseBrace { $$ = $2 }
	;

list
	: tOpenParen values tCloseParen { $$ = $2 }
	;

vector
	: tOpenBracket values tCloseBracket 
	  {
	  	vec := types.Sequence(types.Vector{})
	  	values := types.Sequence($2.v.(*types.List))
		$$.v = vec.Into(values)
	  }
	;
%%
