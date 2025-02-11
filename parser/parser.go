package parser

import (
	"github.com/syedazeez337/monkey-interpreter/ast"
	"github.com/syedazeez337/monkey-interpreter/lexer"
	"github.com/syedazeez337/monkey-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}