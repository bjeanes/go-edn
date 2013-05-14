
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
const tNil = 57357

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
	"tNil",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:118


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 37,
	9, 23,
	-2, 13,
}

const yyNprod = 29
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 56

var yyAct = []int{

	26, 2, 6, 5, 25, 28, 16, 24, 15, 29,
	19, 36, 18, 17, 20, 21, 33, 22, 39, 3,
	30, 27, 32, 5, 23, 5, 31, 35, 31, 5,
	4, 14, 37, 34, 38, 31, 40, 13, 41, 42,
	12, 16, 11, 15, 43, 19, 10, 18, 17, 20,
	21, 9, 22, 8, 7, 1,
}
var yyPact = []int{

	-11, -1000, 37, -1000, -11, -1000, -11, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -11, -11, -1000, -3, -11,
	-1000, -1000, -1000, -1000, -1000, 15, -1000, 11, -11, 2,
	-1000, 37, -1000, -1000, 9, -11, -1000, -11, -11, -1000,
	-1000, 37, -1000, -1000,
}
var yyPgo = []int{

	0, 55, 0, 2, 54, 53, 51, 46, 42, 40,
	37, 31, 30, 19, 4, 27, 9,
}
var yyR1 = []int{

	0, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	12, 13, 13, 2, 2, 14, 14, 11, 10, 9,
	6, 7, 15, 15, 16, 16, 8, 4, 5,
}
var yyR2 = []int{

	0, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 0, 1, 1, 4, 1, 1, 1,
	1, 4, 3, 1, 1, 3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -13, -12, 14, -3, -4, -5, -6,
	-7, -8, -9, -10, -11, 6, 4, 11, 10, 8,
	12, 13, 15, -13, -2, -14, -2, -14, 8, -16,
	-2, -2, 7, 5, -14, -15, 9, -3, -3, 9,
	-2, -2, -2, -3,
}
var yyDef = []int{

	13, -2, 0, 14, 11, 10, 13, 2, 3, 4,
	5, 6, 7, 8, 9, 13, 13, 20, 0, 13,
	19, 18, 17, 12, 1, 13, 15, 13, 13, 0,
	24, 0, 27, 28, 13, 13, 26, -2, 13, 21,
	25, 0, 16, 22,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
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
		//line parser.y:31
		{ result.v = yyS[yypt-1].v }
	case 15:
		//line parser.y:58
		{ yyVAL.v = types.Value(new(types.List))}
	case 16:
		//line parser.y:59
		{
			yyS[yypt-3].v.(*types.List).Insert(yyS[yypt-1].v)
		  }
	case 17:
		//line parser.y:65
		{ yyVAL.v = types.Value(nil) }
	case 21:
		//line parser.y:82
		{ 
		  	set := types.Sequence(types.Set{})
		  	values := types.Sequence(yyS[yypt-1].v.(*types.List))
		  	yyVAL.v = set.Into(values)
		  }
	case 22:
		//line parser.y:90
		{ yyVAL.k = yyS[yypt-2].v; yyVAL.v = yyS[yypt-0].v }
	case 23:
		//line parser.y:91
		{ yylex.Error("Map literal must contain an even number of forms") }
	case 24:
		//line parser.y:95
		{ yyVAL.v = types.Map{} }
	case 25:
		//line parser.y:97
		{
		  	yyS[yypt-2].v.(types.Map)[yyS[yypt-1].k] = yyS[yypt-1].v
		  }
	case 26:
		//line parser.y:103
		{ yyVAL = yyS[yypt-1] }
	case 27:
		//line parser.y:107
		{ yyVAL = yyS[yypt-1] }
	case 28:
		//line parser.y:112
		{
		  	vec := types.Sequence(types.Vector{})
		  	values := types.Sequence(yyS[yypt-1].v.(*types.List))
			yyVAL.v = vec.Into(values)
		  }
	}
	goto yystack /* stack new state and value */
}
