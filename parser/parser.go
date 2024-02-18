package parser

import (
	"cantolang/ast"
	"cantolang/lexer"
	"cantolang/token"
	"fmt"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.advance()
	p.advance()
	return p
}

func (p *Parser) advance() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.ReadToken()
}

func (p *Parser) expectPeek(expectedTokenType string) bool {
	p.advance()
	if p.currentToken.TokenType != expectedTokenType {
		p.errors = append(p.errors, fmt.Sprintf("expected %s got %s", expectedTokenType, p.currentToken.TokenType))
		return false
	}
	return true
}

func (p *Parser) ParseStatement() ast.Statement {
	for p.currentToken.TokenType == token.COMMENT {
		p.advance()
	}
	fmt.Println("T:", p.currentToken.TokenLiteral, p.currentToken.TokenType)
	switch p.currentToken.TokenType {
	default:
		// case token.INITIALIZE:
		s := p.parseInitializeStatement()
		return s
	}
}

func (p *Parser) parseInitializeStatement() *ast.InitializeStatement {
	statement := &ast.InitializeStatement{Token: p.currentToken}
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}
	statement.Identifier = p.currentToken.TokenLiteral
	if !p.expectPeek(token.FULLSTOP) {
		return nil
	}
	p.advance()
	return statement
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.currentToken.TokenType != token.EOF {
		program.Statements = append(program.Statements, p.ParseStatement())
	}
	return program
}
