package object

import (
	"bytes"
	"cantolang/ast"
	"fmt"
)

var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
	ERROR = &Error{}

	//types
	INT_OBJ      = "INT_OBJ"
	STRING_OBJ   = "STRING_OBJ"
	NULL_OBJ     = "NULL_OBJ"
	BOOL_OBJ     = "BOOL_OBJ"
	ERROR_OBJ    = "ERROR_OBJ"
	FUNCTION_OBJ = "FUNCTION_OBJ"
	RETURN_OBJ   = "RETURN_OBJ"
	BUILTIN_OBJ  = "BUILTIN_OBJ"
)

type Object interface {
	Inspect() string
	Type() string
}

type BuiltInFunction func(args ...Object) Object

type Integer struct {
	Value int
}

type String struct {
	Value string
}

type Boolean struct {
	Value bool
}

type Null struct {
}

type Error struct {
	Message     string
	Description string
}

type Function struct {
	Parameters []ast.Identifier
	Body       *ast.BlockStatement
}

type BuiltIn struct {
	Fn BuiltInFunction
}

type ReturnValue struct {
	Value Object
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() string {
	return INT_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}
func (s *String) Type() string {
	return STRING_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
func (b *Boolean) Type() string {
	return BOOL_OBJ
}

func (n *Null) Inspect() string {
	return "NULL"
}
func (n *Null) Type() string {
	return NULL_OBJ
}

func (e *Error) Inspect() string {
	return e.Description + ": " + e.Message
}
func (e *Error) Type() string {
	return ERROR_OBJ
}

func (f *Function) Inspect() string {
	b := bytes.Buffer{}
	b.WriteString("(")
	for i, p := range f.Parameters {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(p.String())
	}
	b.WriteString(") { ")
	b.WriteString(f.Body.String())
	b.WriteString(" }")
	return b.String()
}
func (f *Function) Type() string {
	return FUNCTION_OBJ
}

func (f *BuiltIn) Inspect() string {
	return "builtin function"
}
func (f *BuiltIn) Type() string {
	return BUILTIN_OBJ
}

func (r *ReturnValue) Inspect() string {
	return "return " + r.Value.Inspect()
}
func (r *ReturnValue) Type() string {
	return RETURN_OBJ
}
