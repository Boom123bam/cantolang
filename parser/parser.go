package parser

import (
	"cantolang/ast"
	"cantolang/lexer"
	"cantolang/token"
	"fmt"
	"strconv"
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
	INDEX      // myArr[X]
)

var precedences = map[string]int{
	token.EQUAL_TO:     EQUALS,
	token.LESS_THAN:    LESSGRATER,
	token.GREATER_THAN: LESSGRATER,
	token.ADD:          SUM,
	token.MINUS:        SUM,
	token.MULTIPLY:     PRODUCT,
	token.DIVIDE:       PRODUCT,
	token.OPEN_PAREN:   CALL,
	token.OPEN_BRACKET: INDEX,
}

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	Errors       []string
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
		p.Errors = append(p.Errors, fmt.Sprintf("expected %s got %s", expectedTokenType, p.currentToken.TokenType))
		return false
	}
	return true
}

func (p *Parser) ParseStatement() ast.Statement {
	for p.currentToken.TokenType == token.COMMENT {
		p.advance()
	}
	var s ast.Statement
	switch p.currentToken.TokenType {
	case token.RETURN:
		s = p.parseReturnStatement()
	case token.ASSIGN:
		s = p.parseAssignStatement()
	case token.FUNCTION:
		s = p.parseFunctionDefStatement()
	default:
		s = p.parseExpressionStatement()
	}
	p.advance()
	return s
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.currentToken}
	p.advance()
	statement.Expression = p.parseExpression(LOWEST)
	if p.peekToken.TokenType == token.EOL {
		p.advance()
	}
	return statement
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	statement := &ast.AssignStatement{Token: p.currentToken}
	p.advance()
	statement.Expression = p.parseExpression(LOWEST)
	p.expectPeek(token.TO)
	p.expectPeek(token.IDENTIFIER)
	statement.Identifier = p.currentToken.TokenLiteral
	if p.peekToken.TokenType == token.EOL {
		p.advance()
	}
	return statement
}

func (p *Parser) parseFunctionDefStatement() *ast.FunctionDefStatment {
	statement := &ast.FunctionDefStatment{Token: p.currentToken}
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}
	statement.Identifier = p.currentToken.TokenLiteral
	if !p.expectPeek(token.OPEN_PAREN) {
		return nil
	}
	statement.Parameters = p.parseParams()
	if !p.expectPeek(token.GEWA) {
		return nil
	}
	if !p.expectPeek(token.COMMA) {
		return nil
	}
	if !p.expectPeek(token.THEN) {
		return nil
	}
	if !p.expectPeek(token.OPEN_BRACE) {
		return nil
	}
	p.advance()
	statement.Body = p.parseBlockStatement()
	if p.peekToken.TokenType == token.EOL {
		p.advance()
	}
	return statement

}

func (p *Parser) parseParams() []ast.Identifier {
	params := []ast.Identifier{}
	for p.peekToken.TokenType == token.IDENTIFIER {
		p.advance()
		i := ast.Identifier{Token: p.currentToken}
		params = append(params, i)
		p.advance()
	}
	if p.currentToken.TokenType != token.CLOSE_PAREN {
		return nil
	}
	return params
}

func (p *Parser) parseCallParams() []ast.Expression {
	params := []ast.Expression{}
	if p.peekToken.TokenType == token.CLOSE_PAREN {
		return params
	}
	p.advance()
	params = append(params, p.parseExpression(LOWEST))
	for p.peekToken.TokenType == token.COMMA {
		p.advance()
		p.advance()
		params = append(params, p.parseExpression(LOWEST))
	}
	p.advance()
	if p.currentToken.TokenType != token.CLOSE_PAREN {
		return nil
	}
	return params
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: p.currentToken}
	statement.Expression = p.parseExpression(LOWEST)
	if p.peekToken.TokenType == token.EOL {
		p.advance()
	}
	return statement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	// check for prefix
	var left ast.Expression
	if p.isPrefix(p.currentToken.TokenType) {
		left = p.parsePrefixExpression()
	} else {
		switch p.currentToken.TokenType {
		case token.IDENTIFIER:
			left = &ast.Identifier{Token: p.currentToken}
		case token.OPEN_PAREN:
			p.advance()
			left = p.parseGroupedExpression()
		case token.OPEN_BRACKET:
			left = p.parseArray()
		case token.IF:
			left = p.parseIfExpression()
		case token.TRUE:
			left = &ast.Boolean{Token: p.currentToken, Value: true}
		case token.FALSE:
			left = &ast.Boolean{Token: p.currentToken, Value: false}
		case token.NUMBER:
			val, err := strconv.Atoi(p.currentToken.TokenLiteral)
			if err != nil {
				p.Errors = append(p.Errors, fmt.Sprintf("cannot convert %s(%s) to number", p.currentToken.TokenLiteral, p.currentToken.TokenType))
			}
			left = &ast.IntegerLiteral{Token: p.currentToken, Value: val}
		case token.STRING:
			left = &ast.StringLiteral{Token: p.currentToken, Value: p.currentToken.TokenLiteral}

		default:
			p.Errors = append(p.Errors, fmt.Sprintf("invalid token %s(%s)", p.currentToken.TokenLiteral, p.currentToken.TokenType))
		}
	}

	for p.peekToken.TokenType != token.EOL && precedence < precedences[p.peekToken.TokenType] {
		p.advance()
		// check for infix
		if p.isInfix(p.currentToken.TokenType) {
			left = p.parseInfixExpression(left)
			continue
		} else if p.currentToken.TokenType == token.OPEN_PAREN {
			left = p.parseFunctionCall(left)
			continue
		} else if p.currentToken.TokenType == token.OPEN_BRACKET {
			left = p.parseIndexExpression(left)
			continue
		}
		p.Errors = append(p.Errors, fmt.Sprintf("infix token expected, got %s", p.currentToken.TokenType))
	}
	return left
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	ex := p.parseExpression(LOWEST)
	p.advance()
	if p.currentToken.TokenType != token.CLOSE_PAREN {
		p.Errors = append(p.Errors, fmt.Sprintf("expected ) got %s", p.currentToken.TokenLiteral))
		return nil
	}
	return ex
}

func (p *Parser) parseArray() ast.Expression {
	arr := &ast.ArrayLiteral{Token: p.currentToken}
	p.advance()
	for p.currentToken.TokenType != token.CLOSE_BRACKET && p.currentToken.TokenType != token.EOF {
		arr.Items = append(arr.Items, p.parseExpression(LOWEST))
		p.advance()
		if p.currentToken.TokenType == token.COMMA {
			p.advance()
		}
	}
	if p.currentToken.TokenType == token.CLOSE_BRACKET {
		return arr
	}
	return nil
}

func (p *Parser) parseIfExpression() ast.Expression {
	ex := &ast.IfExpression{Token: p.currentToken}
	if !p.expectPeek(token.OPEN_PAREN) {
		return nil
	}
	p.advance()
	ex.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.CLOSE_PAREN) {
		return nil
	}
	if !p.expectPeek(token.GEWA) {
		return nil
	}
	if !p.expectPeek(token.COMMA) {
		return nil
	}
	if !p.expectPeek(token.THEN) {
		return nil
	}
	if !p.expectPeek(token.OPEN_BRACE) {
		return nil
	}
	p.advance()
	ex.Consequence = p.parseBlockStatement()

	if p.peekToken.TokenType == token.ELSE {
		p.advance()
		if !p.expectPeek(token.OPEN_BRACE) {
			return nil
		}
		p.advance()
		ex.Alternative = p.parseBlockStatement()
	}
	return ex
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Left: left, Token: p.currentToken}
	p.advance()
	index := p.parseExpression(LOWEST)
	p.expectPeek(token.CLOSE_BRACKET)
	exp.Index = index
	return exp
}

func (p *Parser) parseFunctionCall(left ast.Expression) ast.Expression {
	id, ok := left.(*ast.Identifier)
	if !ok {
		p.Errors = append(p.Errors, fmt.Sprintf("expected identifier got %T", left))
	}
	fce := &ast.FunctionCallExpression{Identifier: id}
	fce.Parameters = p.parseCallParams()
	return fce
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	bs := &ast.BlockStatement{}
	for p.currentToken.TokenType != token.CLOSE_BRACE {
		s := p.ParseStatement()
		if s != nil {
			bs.Statements = append(bs.Statements, s)
		}
	}
	return bs
}

func (p *Parser) isPrefix(tokenType string) bool {
	for _, prefix := range p.prefixes {
		if tokenType == prefix {
			return true
		}
	}
	return false
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
	expression := &ast.PrefixExpression{PrefixToken: p.currentToken}
	p.advance()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{Left: left, Infix: p.currentToken}
	precedence, ok := precedences[expression.Infix.TokenType]
	if !ok {
		p.Errors = append(p.Errors, fmt.Sprintf("Infix not found: %s", expression.Infix.TokenType))
		p.advance()
		return nil
	}
	p.advance()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.currentToken.TokenType != token.EOF {
		program.Statements = append(program.Statements, p.ParseStatement())
	}
	return program
}
