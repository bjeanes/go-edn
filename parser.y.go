
//line parser.y:2
package edn 
import __yyfmt__ "fmt"
//line parser.y:2
		
/*
If this file is not parser.y, it was generated from parser.y and
should not be edited directly.
*/

import "github.com/bjeanes/go-edn/types"

func init() {
	//yyDebug = 4
}

//line parser.y:16
type yySymType struct {
	yys int 
	k types.Value
	v types.Value
}

const tOpenBracket = 57346
const tCloseBracket = 57347
const tOpenParen = 57348
const tCloseParen = 57349
const tOpenBrace = 57350
const tCloseBrace = 57351
const tOctothorpe = 57352
const tString = 57353
const tKeyword = 57354
const tCharacter = 57355
const tWhitespace = 57356

var yyToknames = []string{
	"tOpenBracket",
	"tCloseBracket",
	"tOpenParen",
	"tCloseParen",
	"tOpenBrace",
	"tCloseBrace",
	"tOctothorpe",
	"tString",
	"tKeyword",
	"tCharacter",
	"tWhitespace",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:112


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 35,
	9, 21,
	-2, 12,
}

const yyNprod = 27
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 53

var yyAct = []int{

	24, 2, 6, 5, 23, 26, 15, 22, 14, 31,
	18, 34, 17, 16, 19, 20, 37, 27, 5, 28,
	25, 5, 30, 3, 29, 33, 29, 4, 21, 5,
	35, 32, 36, 29, 38, 13, 39, 40, 12, 15,
	11, 14, 41, 18, 10, 17, 16, 19, 20, 9,
	8, 7, 1,
}
var yyPact = []int{

	-11, -1000, 35, -1000, -11, -1000, -11, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -11, -11, -1000, -3, -11, -1000,
	-1000, -1000, -1000, 15, -1000, 4, -11, 2, -1000, 35,
	-1000, -1000, 7, -11, -1000, -11, -11, -1000, -1000, 35,
	-1000, -1000,
}
var yyPgo = []int{

	0, 52, 0, 2, 51, 50, 49, 44, 40, 38,
	35, 27, 23, 4, 25, 17,
}
var yyR1 = []int{

	0, 1, 3, 3, 3, 3, 3, 3, 3, 11,
	12, 12, 2, 2, 13, 13, 10, 9, 6, 7,
	14, 14, 15, 15, 8, 4, 5,
}
var yyR2 = []int{

	0, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 0, 1, 1, 4, 1, 1, 1, 4,
	3, 1, 1, 3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -12, -11, 14, -3, -4, -5, -6,
	-7, -8, -9, -10, 6, 4, 11, 10, 8, 12,
	13, -12, -2, -13, -2, -13, 8, -15, -2, -2,
	7, 5, -13, -14, 9, -3, -3, 9, -2, -2,
	-2, -3,
}
var yyDef = []int{

	12, -2, 0, 13, 10, 9, 12, 2, 3, 4,
	5, 6, 7, 8, 12, 12, 18, 0, 12, 17,
	16, 11, 1, 12, 14, 12, 12, 0, 22, 0,
	25, 26, 12, 12, 24, -2, 12, 19, 23, 0,
	15, 20,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer, result *yySymType) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.y:30
		{ result.v = yyS[yypt-1].v }
	case 14:
		//line parser.y:56
		{ yyVAL.v = types.Value(new(types.List))}
	case 15:
		//line parser.y:57
		{
			yyS[yypt-3].v.(*types.List).Insert(yyS[yypt-1].v)
		  }
	case 19:
		//line parser.y:76
		{ 
		  	set := types.Sequence(types.Set{})
		  	values := types.Sequence(yyS[yypt-1].v.(*types.List))
		  	yyVAL.v = set.Into(values)
		  }
	case 20:
		//line parser.y:84
		{ yyVAL.k = yyS[yypt-2].v; yyVAL.v = yyS[yypt-0].v }
	case 21:
		//line parser.y:85
		{ yylex.Error("Map literal must contain an even number of forms") }
	case 22:
		//line parser.y:89
		{ yyVAL.v = types.Map{} }
	case 23:
		//line parser.y:91
		{
		  	yyS[yypt-2].v.(types.Map)[yyS[yypt-1].k] = yyS[yypt-1].v
		  }
	case 24:
		//line parser.y:97
		{ yyVAL = yyS[yypt-1] }
	case 25:
		//line parser.y:101
		{ yyVAL = yyS[yypt-1] }
	case 26:
		//line parser.y:106
		{
		  	vec := types.Sequence(types.Vector{})
		  	values := types.Sequence(yyS[yypt-1].v.(*types.List))
			yyVAL.v = vec.Into(values)
		  }
	}
	goto yystack /* stack new state and value */
}
