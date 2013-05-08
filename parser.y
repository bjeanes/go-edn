%{
package edn 
import "fmt" 
%}

%union { 
	n int
} 

%token openBracket closeBracket openParen closeParen openBrace closeBrace octothorpe
%% 
input: /* empty */
;
%% 
