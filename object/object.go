package object

var (
	NULL  = Null{}
	TRUE  = Boolean{Value: true}
	FALSE = Boolean{Value: false}
)

type Object interface {
}

type Integer struct {
	Value int
}

type Boolean struct {
	Value bool
}

type Null struct {
}
