package ast

import (
	"bytes"
	"cantolang/token"
)

type Program struct {
	Statements []Statement
}

type Statement interface {
	String() string
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

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Statements []Statement
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

func (ie *IfExpression) token() *token.Token {
	return &ie.Token
}
func (ie *IfExpression) String() string {
	buff := bytes.Buffer{}
	buff.WriteString("if")
	buff.WriteString(ie.Condition.String())
	buff.WriteString(ie.Consequence.String())
	if ie.Alternative == nil {
		return buff.String()
	}
	buff.WriteString("else")
	buff.WriteString(ie.Alternative.String())
	return buff.String()
}

func (bs *BlockStatement) String() string {
	buff := bytes.Buffer{}
	for _, s := range bs.Statements {
		buff.WriteString(s.String())
	}
	return buff.String()
}

func (is *InitializeStatement) String() string {
	return is.Token.TokenLiteral + is.Identifier
}

func (es *ExpressionStatement) String() string {
	return es.Expression.String()
}
