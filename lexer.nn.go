package edn

import (
	"bufio"
	"io"
	"strings"
)

type dfa struct {
	acc []bool
	f   []func(rune) int
	id  int
}
type family struct {
	a       []dfa
	endcase int
}

var a0 [1]dfa
var a []family

func init() {
	a = make([]family, 1)
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			default:
				switch {
				default:
					return 1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[0].acc = acc[:]
		a0[0].f = fun[:]
		a0[0].id = 0
	}
	a[0].endcase = 1
	a[0].a = a0[:]
}
func getAction(c *frame) int {
	if -1 == c.match {
		return -1
	}
	c.action = c.fam.a[c.match].id
	c.match = -1
	return c.action
}

type frame struct {
	atEOF                    bool
	action, match, matchn, n int
	buf                      []rune
	text                     string
	in                       *bufio.Reader
	state                    []int
	fam                      family
}

func newFrame(in *bufio.Reader, index int) *frame {
	f := new(frame)
	f.buf = make([]rune, 0, 128)
	f.in = in
	f.match = -1
	f.fam = a[index]
	f.state = make([]int, len(f.fam.a))
	return f
}

type Lexer []*frame

func NewLexer(in io.Reader) Lexer {
	stack := make([]*frame, 0, 4)
	stack = append(stack, newFrame(bufio.NewReader(in), 0))
	return stack
}
func (stack Lexer) isDone() bool {
	return 1 == len(stack) && stack[0].atEOF
}
func (stack Lexer) nextAction() int {
	c := stack[len(stack)-1]
	for {
		if c.atEOF {
			return c.fam.endcase
		}
		if c.n == len(c.buf) {
			r, _, er := c.in.ReadRune()
			switch er {
			case nil:
				c.buf = append(c.buf, r)
			case io.EOF:
				c.atEOF = true
				if c.n > 0 {
					c.text = string(c.buf)
					return getAction(c)
				}
				return c.fam.endcase
			default:
				panic(er.Error())
			}
		}
		jammed := true
		r := c.buf[c.n]
		for i, x := range c.fam.a {
			if -1 == c.state[i] {
				continue
			}
			c.state[i] = x.f[c.state[i]](r)
			if -1 == c.state[i] {
				continue
			}
			jammed = false
			if x.acc[c.state[i]] {
				if -1 == c.match || c.matchn < c.n+1 || c.match > i {
					c.match = i
					c.matchn = c.n + 1
				}
			}
		}
		if jammed {
			a := getAction(c)
			if -1 == a {
				c.matchn = c.n + 1
			}
			c.n = 0
			for i, _ := range c.state {
				c.state[i] = 0
			}
			c.text = string(c.buf[:c.matchn])
			copy(c.buf, c.buf[c.matchn:])
			c.buf = c.buf[:len(c.buf)-c.matchn]
			return a
		}
		c.n++
	}
	panic("unreachable")
}
func (stack Lexer) push(index int) Lexer {
	c := stack[len(stack)-1]
	return append(stack,
		newFrame(bufio.NewReader(strings.NewReader(c.text)), index))
}
func (stack Lexer) pop() Lexer {
	return stack[:len(stack)-1]
}
func (stack Lexer) Text() string {
	c := stack[len(stack)-1]
	return c.text
}
func (yylex Lexer) Error(e string) {
	panic(e)
}
func (yylex Lexer) Lex(lval *yySymType) int {
	for !yylex.isDone() {
		switch yylex.nextAction() {
		case -1:
		case 0: //./
			{
				return 0
			}
		case 1: ///
			// [END]
		}
	}
	return 0
}

//
