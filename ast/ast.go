package ast

import (
	"bytes"
	"cantolang/token"
)

type Node interface {
	String() string
}

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

type AssignStatement struct {
	Token      token.Token // token.assign
	Identifier string
	Expression Expression
}

type ReturnStatement struct {
	Token      token.Token // token.return
	Expression Expression
}

type ExpressionStatement struct {
	Token      token.Token // first token of ExpressionStatement
	Expression Expression
}

type IntegerLiteral struct {
	Token token.Token
	Value int
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type ArrayLiteral struct {
	Token token.Token
	Items []Expression
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

type Boolean struct {
	Token token.Token
	Value bool
}

type FunctionDefStatment struct {
	Token      token.Token // token.function
	Identifier string
	Parameters []Identifier
	Body       *BlockStatement
}

type Identifier struct {
	Token token.Token
}

type FunctionCallExpression struct {
	Identifier *Identifier
	Parameters []Expression
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

type WhileLoop struct {
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
}

type BlockStatement struct {
	Statements []Statement
}

func (p *Program) String() string {
	buff := bytes.Buffer{}
	for _, s := range p.Statements {
		buff.WriteString(s.String())
	}
	return buff.String()
}

func (il *IntegerLiteral) token() *token.Token {
	return &il.Token
}
func (il *IntegerLiteral) String() string {
	return il.Token.TokenLiteral
}

func (sl *StringLiteral) token() *token.Token {
	return &sl.Token
}
func (sl *StringLiteral) String() string {
	return `"` + sl.Token.TokenLiteral + `"`
}

func (al *ArrayLiteral) token() *token.Token {
	return &al.Token
}
func (al *ArrayLiteral) String() string {
	buff := bytes.Buffer{}
	buff.WriteString("[")
	for i, item := range al.Items {
		if i != 0 {
			buff.WriteString(", ")
		}
		buff.WriteString(item.String())
	}
	buff.WriteString("]")
	return buff.String()
}

func (ie *IndexExpression) token() *token.Token {
	return &ie.Token
}
func (ie *IndexExpression) String() string {
	return ie.Left.String() + "[" + ie.Index.String() + "]"
}

func (b *Boolean) token() *token.Token {
	return &b.Token
}
func (b *Boolean) String() string {
	return b.Token.TokenLiteral
}

func (fd *FunctionDefStatment) String() string {
	buff := bytes.Buffer{}
	buff.WriteString(fd.Token.TokenLiteral + " ")
	buff.WriteString(fd.Identifier + "(")
	for i, param := range fd.Parameters {
		if i != 0 {
			buff.WriteString(",")
		}
		buff.WriteString(param.String())
	}
	buff.WriteString(") {")
	buff.WriteString(fd.Body.String())
	buff.WriteString("}")

	return buff.String()
}

func (i *Identifier) token() *token.Token {
	return &i.Token
}
func (i *Identifier) String() string {
	return i.Token.TokenLiteral
}

func (fe *FunctionCallExpression) token() *token.Token {
	return &fe.Identifier.Token
}
func (fe *FunctionCallExpression) String() string {
	buff := bytes.Buffer{}
	buff.WriteString(fe.token().TokenLiteral + "(")
	for i, param := range fe.Parameters {
		if i != 0 {
			buff.WriteString(", ")
		}
		buff.WriteString(param.String())
	}
	return buff.String()
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

func (as *AssignStatement) String() string {
	return as.Token.TokenLiteral + as.Expression.String() + "-> " + as.Identifier
}

func (rs *ReturnStatement) String() string {
	return rs.Token.TokenLiteral + rs.Expression.String()
}

func (es *ExpressionStatement) String() string {
	return es.Expression.String()
}

func (wl *WhileLoop) token() *token.Token {
	return &wl.Token
}
func (wl *WhileLoop) String() string {
	buff := bytes.Buffer{}
	buff.WriteString("while")
	buff.WriteString(wl.Condition.String())
	buff.WriteString(wl.Body.String())
	return buff.String()
}
