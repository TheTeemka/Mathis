package logic

import (
	"fmt"
)

type Parser struct {
	*lexer
}

func NewParser(text string) *Parser {
	return &Parser{
		lexer: newLexer(text),
	}
}

func (p *Parser) Parse() *Node {
	return p.expr()
}

func (p *Parser) expr() *Node {
	node := p.term()
	for kind, x := p.nextToken(); kind == C_Plus || kind == C_Minus; kind, x = p.nextToken() {
		tmp := NewNode(kind, x)
		p.consumeToken()
		tmp.Left = node
		node = tmp
		node.Right = p.term()
	}
	return node
}
func (p *Parser) term() *Node {
	node := p.factor()
	for kind, x := p.nextToken(); kind == C_Divide || kind == C_Multiply; kind, x = p.nextToken() {
		tmp := NewNode(kind, x)
		p.consumeToken()
		tmp.Left = node
		node = tmp
		node.Right = p.factor()
	}
	return node
}
func (p *Parser) factor() *Node {
	kind, x := p.nextToken()

	switch kind {
	case C_Minus, C_Plus:
		p.consumeToken()
		factor := p.factor()
		return UnaryNode(kind, factor)
	case C_Num:
		p.consumeToken()
		return NewNode(kind, x)
	case C_LeftBracket:
		p.consumeToken()
		expr := p.expr()
		kind, _ = p.nextToken()
		if kind != C_RightBracket {
			fmt.Printf("Expected Right Bracket")
		}
		p.consumeToken()
		return expr
	default:
		fmt.Printf("Wanted a Number or LeftBracket")
	}
	return nil
}