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
	v types.Value
} 

%token tOpenBracket tCloseBracket
%token tOpenParen tCloseParen
%token tOpenBrace tCloseBrace
%token tOctothorpe
%token tString tKeyword
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

map
	: tOpenBrace tCloseBrace
	  { $$.v = types.Map{} }
	;

list
	: tOpenParen values tCloseParen { $$.v = $2.v }
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
