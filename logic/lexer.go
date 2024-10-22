package logic

import "strings"

type TokenType int16

const (
	C_Start        TokenType = iota
	C_End          TokenType = iota
	C_Num          TokenType = iota
	C_LeftBracket  TokenType = iota
	C_RightBracket TokenType = iota
	C_Plus         TokenType = iota
	C_Minus        TokenType = iota
	C_Multiply     TokenType = iota
	C_Divide       TokenType = iota
)

type lexer struct {
	text   []byte
	cursor int
}

func newLexer(text string) *lexer {
	text = strings.ReplaceAll(text, " ", "")
	return &lexer{
		text: []byte(text),
	}
}
func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func (r *lexer) nextToken() (TokenType, ExpType) {
	cur := r.cursor
	if len(r.text) == cur {
		return C_End, -1
	}
	b := r.text[cur]
	cur++
	if isDigit(b) {
		x := ExpType(b - '0')
		for len(r.text) != cur && isDigit(r.text[cur]) {
			x = x*10 + ExpType(r.text[cur]-'0')
			cur++
		}
		return C_Num, x
	}
	switch b {
	case '(':
		return C_LeftBracket, 0
	case ')':
		return C_RightBracket, 0
	case '+':
		return C_Plus, 0
	case '-':
		return C_Minus, 0
	case '*':
		return C_Multiply, 0
	case '/':
		return C_Divide, 0
	}

	return C_End, -1
}
func (r *lexer) consumeToken() {
	cur := r.cursor
	if len(r.text) == cur {
		return
	}
	b := r.text[cur]
	cur++
	if isDigit(b) {
		x := ExpType(b - '0')
		for len(r.text) != cur && isDigit(r.text[cur]) {
			x = x*10 + ExpType(r.text[cur]-'0')
			cur++
		}
	}
	r.cursor = cur
}
