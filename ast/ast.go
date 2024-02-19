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
	String() string
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

type Identifier struct {
	Token token.Token
}

type PrefixExpression struct {
	PrefixToken token.Token
	Right       Expression
}

type InfixExpression struct {
	Left  Expression
	Infix token.Token
	Right Expression
}

func (il *IntegerLiteral) token() *token.Token {
	return &il.Token
}
func (il *IntegerLiteral) String() string {
	return il.Token.TokenLiteral
}

func (i *Identifier) token() *token.Token {
	return &i.Token
}
func (i *Identifier) String() string {
	return i.Token.TokenLiteral
}

func (pe *PrefixExpression) token() *token.Token {
	return &pe.PrefixToken
}
func (pe *PrefixExpression) String() string {
	return pe.PrefixToken.TokenLiteral + pe.Right.String()
}

func (ie *InfixExpression) token() *token.Token {
	return &ie.Infix
}
func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + ie.Infix.TokenLiteral + " " + ie.Right.String() + ")"
}
