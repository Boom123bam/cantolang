package parser

import (
	"cantolang/ast"
	"cantolang/lexer"
	"cantolang/token"
	"fmt"
)

const (
	_ int = iota
	LOWEST
	EQUALS     // ==
	LESSGRATER // > or <
	SUM        // +
	PRODUCT    // *
	PREFIX     // -X or !X
	CALL       // myFunction(X)
)

var precedences = map[string]int{
	token.EQUAL_TO:     EQUALS,
	token.LESS_THAN:    LESSGRATER,
	token.GREATER_THAN: LESSGRATER,
	token.ADD:          SUM,
	token.MINUS:        SUM,
	token.MULTIPLY:     PRODUCT,
	token.DIVIDE:       PRODUCT,
}

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
	prefixes     []string
	infixes      []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.advance()
	p.advance()
	p.prefixes = []string{token.MINUS, token.NOT}
	p.infixes = []string{token.ADD, token.MINUS, token.MULTIPLY, token.DIVIDE, token.EQUAL_TO, token.GREATER_THAN, token.LESS_THAN}

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
	switch p.currentToken.TokenType {
	case token.INITIALIZE:
		s := p.parseInitializeStatement()
		return s
	default:
		s := p.parseExpressionStatement()
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

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: p.currentToken}
	statement.Expression = p.parseExpression(LOWEST)
	return statement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	// check for prefix
	var left ast.Expression
	for _, prefix := range p.prefixes {
		if p.currentToken.TokenType == prefix {
			left = p.parsePrefixExpression()
		}
	}
	// if no prefix set left
	if left == nil {
		left = &ast.IntegerLiteral{Token: p.currentToken}
	}
	p.advance()

	for p.currentToken.TokenType != token.FULLSTOP && precedence < precedences[p.currentToken.TokenType] {
		// check for infix
		if p.isInfix(p.peekToken.TokenType) {
			left = p.parseInfixExpression(left)
			continue
		}
		p.errors = append(p.errors, fmt.Sprintf("infix token expected, got %s", p.currentToken.TokenType))
		p.advance()
	}

	if p.currentToken.TokenType == token.FULLSTOP {
		p.advance()
	}
	return left
}

func (p *Parser) isInfix(tokenType string) bool {
	for _, infix := range p.infixes {
		if tokenType == infix {
			return true
		}
	}
	return false
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	// TODO
	return nil
}
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	// TODO
	return nil
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.currentToken.TokenType != token.EOF {
		program.Statements = append(program.Statements, p.ParseStatement())
	}
	return program
}
