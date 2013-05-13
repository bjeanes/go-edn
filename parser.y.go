
//line parser.y:2
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

//line parser.y:17
type yySymType struct {
	yys int 
	v Value
}

const tOpenBracket = 57346
const tCloseBracket = 57347
const tOpenParen = 57348
const tCloseParen = 57349
const tOpenBrace = 57350
const tCloseBrace = 57351
const tOctothorpe = 57352
const tString = 57353

var yyToknames = []string{
	"tOpenBracket",
	"tCloseBracket",
	"tOpenParen",
	"tCloseParen",
	"tOpenBrace",
	"tCloseBrace",
	"tOctothorpe",
	"tString",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:84

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 23
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 51

var yyAct = []int{

	14, 2, 13, 31, 24, 23, 17, 18, 19, 20,
	21, 28, 22, 16, 25, 15, 7, 6, 17, 18,
	19, 20, 21, 25, 5, 4, 29, 30, 26, 3,
	25, 1, 27, 17, 18, 19, 20, 21, 17, 18,
	19, 20, 21, 9, 0, 8, 0, 12, 0, 11,
	10,
}
var yyPact = []int{

	39, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 26, 26,
	-1000, -3, -5, 21, -1000, -1000, 26, -1000, -1000, -1000,
	-1000, -1000, 6, 26, -1000, 39, -1000, -1000, -1000, -6,
	-1000, -1000,
}
var yyPgo = []int{

	0, 31, 1, 29, 25, 24, 17, 16, 13, 15,
	0, 2,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 8, 8, 8,
	8, 8, 9, 9, 10, 10, 11, 11, 5, 6,
	7, 3, 4,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 0, 1, 1, 3, 1, 4,
	2, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 6, 4,
	11, 10, 8, -11, -10, -9, -8, 12, 13, 14,
	15, 16, -11, 8, 9, -10, 7, -9, 5, -11,
	-2, 9,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 14, 14,
	18, 0, 0, 14, 16, 15, 12, 7, 8, 9,
	10, 11, 14, 14, 20, 0, 21, 13, 22, 14,
	17, 19,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 13,
	16, 3, 3, 15, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 12, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 14,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
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
	if c > 0 && c <= len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return fmt.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return fmt.Sprintf("state-%v", s)
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
		fmt.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
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
		fmt.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
				fmt.Printf("%s", yyStatname(yystate))
				fmt.Printf("saw %s\n", yyTokname(yychar))
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
					fmt.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				fmt.Printf("error recovery discards %s\n", yyTokname(yychar))
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
		fmt.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
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
		//line parser.y:29
		{ lastResult = Value(yyS[yypt-0].v) }
	case 16:
		//line parser.y:57
		{ yyVAL.v = Value(new(List))}
	case 17:
		//line parser.y:58
		{
			yyVAL.v.(*List).raw().PushBack(yyS[yypt-0].v)
		  }
	case 19:
		//line parser.y:69
		{ yyVAL.v = Sequence(Set{}).Into(Sequence(yyS[yypt-1].v.(*List))) }
	case 20:
		//line parser.y:74
		{ yyVAL.v = Map{} }
	case 21:
		//line parser.y:78
		{ yyVAL.v = yyS[yypt-1].v }
	case 22:
		//line parser.y:82
		{ yyVAL.v = Sequence(Vector{}).Into(Sequence(yyS[yypt-1].v.(*List))) }
	}
	goto yystack /* stack new state and value */
}
