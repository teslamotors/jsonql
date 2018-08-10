// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/teslamotors/jsonql/token"
)

const (
	NoState    = -1
	NumStates  = 80
	NumSymbols = 68
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '_'
1: '.'
2: '.'
3: '"'
4: '\'
5: '"'
6: '"'
7: '''
8: '\'
9: '''
10: '''
11: '-'
12: '!'
13: 'n'
14: 'u'
15: 'l'
16: 'l'
17: 't'
18: 'r'
19: 'u'
20: 'e'
21: 'f'
22: 'a'
23: 'l'
24: 's'
25: 'e'
26: '.'
27: '['
28: ']'
29: '_'
30: 'e'
31: 'E'
32: '+'
33: '-'
34: '0'
35: '0'
36: 'x'
37: 'X'
38: '\'
39: 'x'
40: '\'
41: 'u'
42: 'b'
43: 'f'
44: 'n'
45: 'r'
46: 't'
47: 'v'
48: '\'
49: '\'
50: ' '
51: '\t'
52: '\f'
53: '\v'
54: \u00a0
55: \u202f
56: \u205f
57: \u3000
58: \ufeff
59: 'A'-'Z'
60: 'a'-'z'
61: '0'-'9'
62: '1'-'9'
63: '0'-'7'
64: 'a'-'f'
65: 'A'-'F'
66: \u2000-\u200a
67: .
*/
