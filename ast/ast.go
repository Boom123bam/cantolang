package ast

import "cantolang/token"

type Program struct {
	Statements []Statement
}

type Statement interface {
}

type InitializeStatement struct {
	Token      token.Token
	Identifier string
}
