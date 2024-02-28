package object

import "fmt"

var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
	ERROR = &Error{}

	//types
	INT_OBJ   = "INT_OBJ"
	NULL_OBJ  = "NULL_OBJ"
	BOOL_OBJ  = "BOOL_OBJ"
	ERROR_OBJ = "ERROR_OBJ"
)

type Object interface {
	Inspect() string
	Type() string
}

type Integer struct {
	Value int
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

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() string {
	return INT_OBJ
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
