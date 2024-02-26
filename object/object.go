package object

import "fmt"

var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

type Object interface {
	Inspect() string
}

type Integer struct {
	Value int
}

type Boolean struct {
	Value bool
}

type Null struct {
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (n *Null) Inspect() string {
	return "NULL"
}
