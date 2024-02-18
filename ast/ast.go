package ast

import "cantolang/token"

type Program struct {
	Statements []Statement
}

type Statement interface {
}

type Expression interface {
	// IntegerLiteral, InfixExpression, PrefixExpression
	token() *token.Token
}

type InitializeStatement struct {
	Token      token.Token // token.initialize
	Identifier string
}

type ExpressionStatement struct {
	Token      token.Token // first token of ExpressionStatement
	Expression Expression
}

type IntegerLiteral struct {
	Token token.Token
}

type PrefixExpression struct {
	PrefixToken token.Token
	Expression  Expression
}

func (il *IntegerLiteral) token() *token.Token {
	return &il.Token
}
