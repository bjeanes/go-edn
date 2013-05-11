%{
package edn 
import "fmt" 
%}

%union { 
	n int
} 

%token openBracket closeBracket openParen closeParen openBrace closeBrace octothorpe
%% 
input: value
;

value: list
| vector
| string
;

string: '"' '"'
;

list: '(' ')'
;

vector: '[' ']'
;
%% 
