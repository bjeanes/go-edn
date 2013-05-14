package edn

import "github.com/bjeanes/go-edn/types"
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

var a0 [17]dfa
var a []family

func init() {
	a = make([]family, 1)
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 13:
				return 1
			case 10:
				return 1
			case 9:
				return 1
			case 44:
				return 1
			case 32:
				return 1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			case 13:
				return 1
			case 10:
				return 1
			case 9:
				return 1
			case 44:
				return 1
			case 32:
				return 1
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
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 91:
				return 1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			case 91:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[1].acc = acc[:]
		a0[1].f = fun[:]
		a0[1].id = 1
	}
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 93:
				return 1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			case 93:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[2].acc = acc[:]
		a0[2].f = fun[:]
		a0[2].id = 2
	}
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 123:
				return 1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			case 123:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[3].acc = acc[:]
		a0[3].f = fun[:]
		a0[3].id = 3
	}
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 125:
				return 1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			case 125:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[4].acc = acc[:]
		a0[4].f = fun[:]
		a0[4].id = 4
	}
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 40:
				return 1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			case 40:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[5].acc = acc[:]
		a0[5].f = fun[:]
		a0[5].id = 5
	}
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 41:
				return 1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			case 41:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[6].acc = acc[:]
		a0[6].f = fun[:]
		a0[6].id = 6
	}
	{
		var acc [2]bool
		var fun [2]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 35:
				return 1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[1] = true
		fun[1] = func(r rune) int {
			switch r {
			case 35:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[7].acc = acc[:]
		a0[7].f = fun[:]
		a0[7].id = 7
	}
	{
		var acc [9]bool
		var fun [9]func(rune) int
		fun[1] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 110:
				return 2
			case 101:
				return -1
			case 119:
				return -1
			case 108:
				return -1
			case 105:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[2] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 110:
				return -1
			case 101:
				return 3
			case 119:
				return -1
			case 108:
				return -1
			case 105:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[3] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 110:
				return -1
			case 101:
				return -1
			case 119:
				return 4
			case 108:
				return -1
			case 105:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[7] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 110:
				return -1
			case 101:
				return 8
			case 119:
				return -1
			case 108:
				return -1
			case 105:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[0] = func(r rune) int {
			switch r {
			case 92:
				return 1
			case 110:
				return -1
			case 101:
				return -1
			case 119:
				return -1
			case 108:
				return -1
			case 105:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[4] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 110:
				return -1
			case 101:
				return -1
			case 119:
				return -1
			case 108:
				return 5
			case 105:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[5] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 110:
				return -1
			case 101:
				return -1
			case 119:
				return -1
			case 108:
				return -1
			case 105:
				return 6
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[6] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 110:
				return 7
			case 101:
				return -1
			case 119:
				return -1
			case 108:
				return -1
			case 105:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[8] = true
		fun[8] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 110:
				return -1
			case 101:
				return -1
			case 119:
				return -1
			case 108:
				return -1
			case 105:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[8].acc = acc[:]
		a0[8].f = fun[:]
		a0[8].id = 8
	}
	{
		var acc [8]bool
		var fun [8]func(rune) int
		fun[1] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 114:
				return 2
			case 101:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			case 110:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[4] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 114:
				return -1
			case 101:
				return -1
			case 116:
				return -1
			case 117:
				return 5
			case 110:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[5] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 114:
				return 6
			case 101:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			case 110:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[7] = true
		fun[7] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 114:
				return -1
			case 101:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			case 110:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[0] = func(r rune) int {
			switch r {
			case 92:
				return 1
			case 114:
				return -1
			case 101:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			case 110:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[2] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 114:
				return -1
			case 101:
				return 3
			case 116:
				return -1
			case 117:
				return -1
			case 110:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[3] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 114:
				return -1
			case 101:
				return -1
			case 116:
				return 4
			case 117:
				return -1
			case 110:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[6] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 114:
				return -1
			case 101:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			case 110:
				return 7
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[9].acc = acc[:]
		a0[9].f = fun[:]
		a0[9].id = 9
	}
	{
		var acc [5]bool
		var fun [5]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 92:
				return 1
			case 116:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[1] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 116:
				return 2
			case 97:
				return -1
			case 98:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[2] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 116:
				return -1
			case 97:
				return 3
			case 98:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[3] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 116:
				return -1
			case 97:
				return -1
			case 98:
				return 4
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[4] = true
		fun[4] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 116:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[10].acc = acc[:]
		a0[10].f = fun[:]
		a0[10].id = 10
	}
	{
		var acc [7]bool
		var fun [7]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 92:
				return 1
			case 115:
				return -1
			case 112:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[1] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 115:
				return 2
			case 112:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[2] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 115:
				return -1
			case 112:
				return 3
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[3] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 115:
				return -1
			case 112:
				return -1
			case 97:
				return 4
			case 99:
				return -1
			case 101:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[4] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 115:
				return -1
			case 112:
				return -1
			case 97:
				return -1
			case 99:
				return 5
			case 101:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[5] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 115:
				return -1
			case 112:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return 6
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[6] = true
		fun[6] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 115:
				return -1
			case 112:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[11].acc = acc[:]
		a0[11].f = fun[:]
		a0[11].id = 11
	}
	{
		var acc [7]bool
		var fun [7]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 92:
				return 1
			case 117:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 70:
					return -1
				case 97 <= r && r <= 102:
					return -1
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[1] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 117:
				return 2
			default:
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 70:
					return -1
				case 97 <= r && r <= 102:
					return -1
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[2] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 117:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return 3
				case 65 <= r && r <= 70:
					return 3
				case 97 <= r && r <= 102:
					return 3
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[3] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 117:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return 4
				case 65 <= r && r <= 70:
					return 4
				case 97 <= r && r <= 102:
					return 4
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[4] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 117:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return 5
				case 65 <= r && r <= 70:
					return 5
				case 97 <= r && r <= 102:
					return 5
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[5] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 117:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return 6
				case 65 <= r && r <= 70:
					return 6
				case 97 <= r && r <= 102:
					return 6
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[6] = true
		fun[6] = func(r rune) int {
			switch r {
			case 92:
				return -1
			case 117:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 70:
					return -1
				case 97 <= r && r <= 102:
					return -1
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[12].acc = acc[:]
		a0[12].f = fun[:]
		a0[12].id = 12
	}
	{
		var acc [3]bool
		var fun [3]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 92:
				return 1
			default:
				switch {
				case 33 <= r && r <= 126:
					return -1
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[1] = func(r rune) int {
			switch r {
			case 92:
				return 2
			default:
				switch {
				case 33 <= r && r <= 126:
					return 2
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[2] = true
		fun[2] = func(r rune) int {
			switch r {
			case 92:
				return -1
			default:
				switch {
				case 33 <= r && r <= 126:
					return -1
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[13].acc = acc[:]
		a0[13].f = fun[:]
		a0[13].id = 13
	}
	{
		var acc [4]bool
		var fun [4]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 58:
				return 1
			default:
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[1] = func(r rune) int {
			switch r {
			case 58:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return 2
				case 97 <= r && r <= 122:
					return 2
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[2] = true
		fun[2] = func(r rune) int {
			switch r {
			case 58:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return 3
				case 65 <= r && r <= 90:
					return 3
				case 97 <= r && r <= 122:
					return 3
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		acc[3] = true
		fun[3] = func(r rune) int {
			switch r {
			case 58:
				return -1
			default:
				switch {
				case 48 <= r && r <= 57:
					return 3
				case 65 <= r && r <= 90:
					return 3
				case 97 <= r && r <= 122:
					return 3
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		a0[14].acc = acc[:]
		a0[14].f = fun[:]
		a0[14].id = 14
	}
	{
		var acc [6]bool
		var fun [6]func(rune) int
		fun[0] = func(r rune) int {
			switch r {
			case 34:
				return 1
			case 92:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[1] = func(r rune) int {
			switch r {
			case 34:
				return 2
			case 92:
				return 3
			default:
				switch {
				default:
					return 4
				}
			}
			panic("unreachable")
		}
		acc[2] = true
		fun[2] = func(r rune) int {
			switch r {
			case 34:
				return -1
			case 92:
				return -1
			default:
				switch {
				default:
					return -1
				}
			}
			panic("unreachable")
		}
		fun[3] = func(r rune) int {
			switch r {
			case 34:
				return 5
			case 92:
				return 5
			default:
				switch {
				default:
					return 5
				}
			}
			panic("unreachable")
		}
		fun[4] = func(r rune) int {
			switch r {
			case 34:
				return 2
			case 92:
				return 3
			default:
				switch {
				default:
					return 4
				}
			}
			panic("unreachable")
		}
		fun[5] = func(r rune) int {
			switch r {
			case 34:
				return 2
			case 92:
				return 3
			default:
				switch {
				default:
					return 4
				}
			}
			panic("unreachable")
		}
		a0[15].acc = acc[:]
		a0[15].f = fun[:]
		a0[15].id = 15
	}
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
		a0[16].acc = acc[:]
		a0[16].f = fun[:]
		a0[16].id = 16
	}
	a[0].endcase = 17
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

type lexer []*frame

func newLexer(in io.Reader) lexer {
	stack := make([]*frame, 0, 4)
	stack = append(stack, newFrame(bufio.NewReader(in), 0))
	return stack
}
func (stack lexer) isDone() bool {
	return 1 == len(stack) && stack[0].atEOF
}
func (stack lexer) nextAction() int {
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
func (stack lexer) push(index int) lexer {
	c := stack[len(stack)-1]
	return append(stack,
		newFrame(bufio.NewReader(strings.NewReader(c.text)), index))
}
func (stack lexer) pop() lexer {
	return stack[:len(stack)-1]
}
func (stack lexer) Text() string {
	c := stack[len(stack)-1]
	return c.text
}
func (yylex lexer) Error(e string) {
	panic(e)
}
func (yylex lexer) Lex(lval *yySymType) int {
	for !yylex.isDone() {
		switch yylex.nextAction() {
		case -1:
		case 0: //[\r\n\t, ]+/
			{
				return tWhitespace
			}
		case 1: //\[/
			{
				return tOpenBracket
			}
		case 2: //\]/
			{
				return tCloseBracket
			}
		case 3: //{/
			{
				return tOpenBrace
			}
		case 4: //}/
			{
				return tCloseBrace
			}
		case 5: //\(/
			{
				return tOpenParen
			}
		case 6: //\)/
			{
				return tCloseParen
			}
		case 7: //#/
			{
				return tOctothorpe
			}
		case 8: //\\newline/
			{
				lval.v = types.Character('\n')
				return tCharacter
			}
		case 9: //\\return/
			{
				lval.v = types.Character('\r')
				return tCharacter
			}
		case 10: //\\tab/
			{
				lval.v = types.Character('\t')
				return tCharacter
			}
		case 11: //\\space/
			{
				lval.v = types.Character(' ')
				return tCharacter
			}
		case 12: //\\u[0-9A-Fa-f][0-9A-Fa-f][0-9A-Fa-f][0-9A-Fa-f]/
			{
				panic("Unicode characters are currently unimplemented")
			}
		case 13: //\\[!-~]/
			{
				s := yylex.Text()[1:]
				ch, _, _ := strings.NewReader(s).ReadRune()
				lval.v = types.Character(ch)
				return tCharacter
			}
		case 14: //:[a-zA-Z][a-zA-Z0-9]*/
			{
				lval.v = types.Keyword(yylex.Text()[1:])
				return tKeyword
			}
		case 15: //"(\\.|[^"\\])*"/
			{
				s := yylex.Text()
				lval.v = types.String(s[1 : len(s)-1])
				return tString
			}
		case 16: //./
			{ // (This rule must be last)
				// Unmatched token...
				return -1
			}
		case 17: ///
			// [END]
		}
	}
	return 0
}
func init() {
	/* (this has to be in a func or lex doesn't write it out to final file)
	 *
	 * If this file is not lexer.nn, it was generated from lexer.nn and
	 * should not be edited directly.
	 *
	 */
}
